package helper

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"strconv"
	"sync"
	"time"

	"nyauth_backed/source"
)

// 验证码结构体，包含验证码和过期时间
type VerificationCode struct {
	Code      string
	ExpiresAt time.Time
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
	from := source.AppConfig.SMTP.Username
	password := source.AppConfig.SMTP.Password

	// 设置SMTP服务器信息
	smtpHost := source.AppConfig.SMTP.Host
	smtpPort := source.AppConfig.SMTP.Port

	// 设置邮件内容
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	// 发送邮件
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+strconv.Itoa(smtpPort), auth, from, []string{to}, []byte(msg))
	return err
}

// SendVerificationCodeByEmail 发送验证码到用户的电子邮件
func SendVerificationCodeByEmail(to string) error {
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
	}
	codeCache.Unlock()

	// 发送验证码到用户的电子邮件
	subject := "[Nyauth] 你的验证码来啦~ 请查收!"
	body := fmt.Sprintf("您的验证码为: %s, 有效期为：%d 分钟, 千万不要泄露给他人哦!", code, expirationMinutes)
	if err := SendEmail(to, subject, body); err != nil {
		return fmt.Errorf("failed to send verification code: %w", err)
	}

	return nil
}

// VerifyCode 验证用户输入的验证码是否正确
func VerifyCode(email, code string) bool {
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

	return storedCode.Code == code
}
