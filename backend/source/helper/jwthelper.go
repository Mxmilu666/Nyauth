package helper

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtHelperCert struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

var JwtHelper *JwtHelperCert

const (
	privateKeyPath = "./data/private.key"
	publicKeyPath  = "./data/public.key"
)

func GetInstance() (*JwtHelperCert, error) {
	if JwtHelper == nil {
		helper := &JwtHelperCert{}
		err := helper.loadKeys()
		if err != nil {
			return nil, err
		}
		JwtHelper = helper
	}
	return JwtHelper, nil
}

func ensureDataDir() error {
	dir := "./data"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return nil
}

// 加载或生成密钥对
func (j *JwtHelperCert) loadKeys() error {
	if fileExists(privateKeyPath) && fileExists(publicKeyPath) {
		// 如果密钥文件存在
		privKeyData, err := ioutil.ReadFile(privateKeyPath)
		if err != nil {
			return err
		}
		privKey, err := parsePrivateKey(privKeyData)
		if err != nil {
			return err
		}
		j.privateKey = privKey

		pubKeyData, err := ioutil.ReadFile(publicKeyPath)
		if err != nil {
			return err
		}
		pubKey, err := parsePublicKey(pubKeyData)
		if err != nil {
			return err
		}
		j.publicKey = pubKey
	} else {
		// 如果密钥不存在，生成新的并保存
		if err := ensureDataDir(); err != nil {
			return err
		}
		privKey, pubKey, err := generateKeys()
		if err != nil {
			return err
		}
		j.privateKey = privKey
		j.publicKey = pubKey

		// 保存密钥到本地文件
		err = savePrivateKey(j.privateKey, privateKeyPath)
		if err != nil {
			return err
		}
		err = savePublicKey(j.publicKey, publicKeyPath)
		if err != nil {
			return err
		}
	}
	return nil
}

// 签发 JWT
func (j *JwtHelperCert) IssueToken(payload map[string]interface{}, audience string, expiresInSeconds int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"aud":  audience,
		"iss":  "Nyauth-Server",
		"exp":  time.Now().Add(time.Duration(expiresInSeconds) * time.Second).Unix(),
		"data": payload,
	})
	return token.SignedString(j.privateKey)
}

// 验证 JWT
func (j *JwtHelperCert) VerifyToken(tokenString string, audience string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return j.publicKey, nil
	}, jwt.WithAudience(audience))

	if err != nil {
		return nil, err
	}
	return token, nil
}

// 生成 RSA 公钥和私钥对
func generateKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

// 保存私钥
func savePrivateKey(key *rsa.PrivateKey, filePath string) error {
	keyBytes := x509.MarshalPKCS1PrivateKey(key)
	pemBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: keyBytes,
	}
	return ioutil.WriteFile(filePath, pem.EncodeToMemory(pemBlock), 0600)
}

// 保存公钥
func savePublicKey(key *rsa.PublicKey, filePath string) error {
	keyBytes, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		return err
	}
	pemBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: keyBytes,
	}
	return ioutil.WriteFile(filePath, pem.EncodeToMemory(pemBlock), 0644)
}

// 解析私钥
func parsePrivateKey(data []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(data)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("failed to decode PEM block containing private key")
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

// 解析公钥
func parsePublicKey(data []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(data)
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		return nil, errors.New("failed to decode PEM block containing public key")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	if pubKey, ok := pub.(*rsa.PublicKey); ok {
		return pubKey, nil
	}
	return nil, errors.New("not RSA public key")
}

// 判断文件是否存在
func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}
