package helper

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"nyauth_backed/source"
	"nyauth_backed/source/untils"

	"gopkg.in/gomail.v2"
)

var ErrVerificationCodeExists = errors.New("verification code already exists and has not expired")

// 验证码结构体
type VerificationCode struct {
	Code      string
	ExpiresAt time.Time
	Email     string
	UseFor    string
}

// 临时注册码结构体
type TemporaryCode struct {
	Code      string
	ExpiresAt time.Time
	Email     string
	UseFor    string
}

// 内存缓存，用于存储验证码
var codeCache = struct {
	sync.RWMutex
	m map[string]VerificationCode
}{m: make(map[string]VerificationCode)}

// 临时注册码缓存
var tempCodeCache = struct {
	sync.RWMutex
	m map[string]TemporaryCode
}{m: make(map[string]TemporaryCode)}

type VerificationRequest struct {
	Email string `json:"email"`
}

func init() {
	// 启动定期清理过期验证码
	go cleanupExpiredCodes()
}

// cleanupExpiredCodes 定期清理过期的验证码和临时码
func cleanupExpiredCodes() {
	ticker := time.NewTicker(5 * time.Minute) // 每5分钟清理一次
	defer ticker.Stop()

	for range ticker.C {
		// 清理过期的验证码
		removeExpiredVerificationCodes()

		// 清理过期的临时注册码
		removeExpiredTemporaryCodes()
	}
}

// removeExpiredVerificationCodes 清理过期的验证码
func removeExpiredVerificationCodes() {
	now := time.Now()

	codeCache.Lock()
	defer codeCache.Unlock()

	// 遍历所有验证码，删除已过期的
	for email, code := range codeCache.m {
		if now.After(code.ExpiresAt) {
			delete(codeCache.m, email)
		}
	}
}

// removeExpiredTemporaryCodes 清理过期的临时注册码
func removeExpiredTemporaryCodes() {
	now := time.Now()

	tempCodeCache.Lock()
	defer tempCodeCache.Unlock()

	// 遍历所有临时注册码，删除已过期的
	for key, code := range tempCodeCache.m {
		if now.After(code.ExpiresAt) {
			delete(tempCodeCache.m, key)
		}
	}
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
	code, err := untils.GenerateRandomCode(6, true)
	if err != nil {
		return fmt.Errorf("failed to generate verification code: %w", err)
	}

	// 验证码有效期为10分钟
	expirationMinutes := 10

	// 将验证码存储在缓存中
	expiration := time.Now().Add(time.Duration(expirationMinutes) * time.Minute)
	codeCache.Lock()
	codeCache.m[to] = VerificationCode{
		Code:      code,
		ExpiresAt: expiration,
		UseFor:    usefor,
		Email:     to,
	}
	codeCache.Unlock()

	var subject string
	var useType string

	switch usefor {
	case "register":
		useType = "注册"
	case "reset_password":
		useType = "重置密码"
	case "multi_identity":
		useType = "绑定多身份"
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

	// 检查验证码是否是对应邮箱
	if storedCode.Email != email {
		return false
	}

	// 检查验证码是否已过期
	if time.Now().After(storedCode.ExpiresAt) {
		// 删除过期的临时码
		tempCodeCache.Lock()
		delete(tempCodeCache.m, email)
		tempCodeCache.Unlock()
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

// GenerateTempCode 生成临时注册码并存储在缓存中
func GenerateTempCode(email, usefor string, expirationMinutes int) (string, error) {
	// 生成验证码
	code, err := untils.GenerateRandomCode(6, true)
	if err != nil {
		return "", fmt.Errorf("failed to generate verification code: %w", err)
	}

	// 创建缓存键
	key := fmt.Sprintf("%s:%s", email, usefor)

	// 存储临时码
	expiration := time.Now().Add(time.Duration(expirationMinutes) * time.Minute)
	tempCodeCache.Lock()
	tempCodeCache.m[key] = TemporaryCode{
		Code:      code,
		ExpiresAt: expiration,
		Email:     email,
		UseFor:    usefor,
	}
	tempCodeCache.Unlock()

	return code, nil
}

// VerifyTempCode 验证临时注册码
func VerifyTempCode(email, code, usefor string) bool {
	key := fmt.Sprintf("%s:%s", email, usefor)

	tempCodeCache.RLock()
	storedCode, exists := tempCodeCache.m[key]
	tempCodeCache.RUnlock()

	if !exists {
		return false
	}

	// 检查是否过期
	if time.Now().After(storedCode.ExpiresAt) {
		// 删除过期的临时码
		tempCodeCache.Lock()
		delete(tempCodeCache.m, key)
		tempCodeCache.Unlock()
		return false
	}

	// 验证码是否匹配
	if storedCode.Code == code {
		// 验证成功后删除
		tempCodeCache.Lock()
		delete(tempCodeCache.m, key)
		tempCodeCache.Unlock()
		return true
	}

	return false
}
