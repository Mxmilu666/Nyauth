package helper

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"nyauth_backed/source"

	"gopkg.in/gomail.v2"
)

var ErrVerificationCodeExists = errors.New("verification code already exists and has not expired")

// 验证码结构体，包含验证码和过期时间
type VerificationCode struct {
	Code      string
	ExpiresAt time.Time
	UseFor    string
}

// 内存缓存，用于存储验证码
var codeCache = struct {
	sync.RWMutex
	m map[string]VerificationCode
}{m: make(map[string]VerificationCode)}

type VerificationRequest struct {
	Email string `json:"email"`
}

func generateVerificationCode() string {
	rand.Seed(time.Now().UnixNano())
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func SendEmail(to, subject, body string) error {
	// 从配置中读取 SMTP 相关信息
	username := source.AppConfig.SMTP.Username
	from := source.AppConfig.SMTP.From
	password := source.AppConfig.SMTP.Password
	smtpHost := source.AppConfig.SMTP.Host
	smtpPort := source.AppConfig.SMTP.Port

	// 创建一封邮件
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(smtpHost, smtpPort, username, password)
	// 发送邮件
	return d.DialAndSend(m)
}

// SendVerificationCodeByEmail 发送验证码到用户的电子邮件
func SendVerificationCodeByEmail(to, usefor string) error {
	// 检查是否存在未过期的验证码
	codeCache.RLock()
	existingCode, exists := codeCache.m[to]
	codeCache.RUnlock()

	// 如果已存在未过期的验证码且用途相同，则不发送验证码防止被刷爆接口
	if exists && time.Now().Before(existingCode.ExpiresAt) && existingCode.UseFor == usefor {
		return fmt.Errorf("%w: please check your email or wait for expiration", ErrVerificationCodeExists)
	}

	// 生成验证码
	code := generateVerificationCode()

	// 验证码有效期为10分钟
	expirationMinutes := 10

	// 将验证码存储在缓存中
	expiration := time.Now().Add(time.Duration(expirationMinutes) * time.Minute)
	codeCache.Lock()
	codeCache.m[to] = VerificationCode{
		Code:      code,
		ExpiresAt: expiration,
		UseFor:    usefor,
	}
	codeCache.Unlock()

	var subject string
	var useType string

	switch usefor {
	case "register":
		useType = "注册"
	case "reset_password":
		useType = "重置密码"
	default:
		return fmt.Errorf("invalid usefor: %s", usefor)
	}

	subject = fmt.Sprintf("[Nyauth] 你的%s验证码来啦~ 请查收!", useType)
	body := fmt.Sprintf("您的%s验证码为: %s, 有效期为：%d 分钟, 千万不要泄露给他人哦!",
		useType,
		code,
		expirationMinutes)

	if err := SendEmail(to, subject, body); err != nil {
		return fmt.Errorf("failed to send verification code: %w", err)
	}

	return nil
}

// VerifyCode 验证用户输入的验证码是否正确
func VerifyCode(email, code, usefor string) bool {
	codeCache.RLock()
	storedCode, exists := codeCache.m[email]
	codeCache.RUnlock()

	if !exists {
		return false
	}

	// 检查验证码是否已过期
	if time.Now().After(storedCode.ExpiresAt) {
		return false
	}

	// 检查验证码用途是否正确
	if usefor != storedCode.UseFor {
		return false
	}

	// 验证码正确，删除缓存中的验证码
	if storedCode.Code == code {
		codeCache.Lock()
		delete(codeCache.m, email)
		codeCache.Unlock()
		return true
	}

	return false
}
