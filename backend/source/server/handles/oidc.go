package handles

import (
	"crypto/rsa"
	"math/big"
	"net/http"
	"nyauth_backed/source"
	"nyauth_backed/source/helper"
	"nyauth_backed/source/models"
	"nyauth_backed/source/untils"

	"github.com/gin-gonic/gin"
)

// JWKS 表示 JSON Web Key Set
type JWKS struct {
	Keys []JWK `json:"keys"`
}

// JWK 表示JSON Web Key
type JWK struct {
	Kid string `json:"kid"`
	Kty string `json:"kty"`
	Use string `json:"use"`
	Alg string `json:"alg"`
	N   string `json:"n"`
	E   string `json:"e"`
}

// GetOpenIDConfiguration 处理 /.well-known/openid-configuration 请求
// 返回OpenID Connect提供者的配置信息
func GetOpenIDConfiguration(c *gin.Context) {
	baseURL := source.AppConfig.Server.BaseURL

	config := models.OIDCConfiguration{
		Issuer:                           baseURL,
		AuthorizationEndpoint:            baseURL + "/oauth/authorize",
		TokenEndpoint:                    baseURL + "/api/v0/oauth/token",
		UserinfoEndpoint:                 baseURL + "/api/v0/account/info",
		JwksURI:                          baseURL + "/.well-known/jwks.json",
		ScopesSupported:                  []string{"openid", "profile", "email"},
		ResponseTypesSupported:           []string{"code"},
		SubjectTypesSupported:            []string{"public"},
		IDTokenSigningAlgValuesSupported: []string{"RS256"},
	}

	c.JSON(http.StatusOK, config)
}

// 返回用于验证签名的公钥信息
func GetJWKS(c *gin.Context) {
	// 获取公钥
	pubKey := helper.JwtHelper.GetPublicKey()
	if pubKey == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "公钥不可用"})
		return
	}

	// 转换RSA公钥到JWK格式
	jwk := rsaPublicKeyToJWK(pubKey)

	// 构建JWKS响应
	jwks := JWKS{
		Keys: []JWK{jwk},
	}

	c.JSON(http.StatusOK, jwks)
}

// rsaPublicKeyToJWK 将RSA公钥转换为JWK格式
func rsaPublicKeyToJWK(pubKey *rsa.PublicKey) JWK {
	// 将模数转换为Base64URL编码
	nBytes := pubKey.N.Bytes()
	e := big.NewInt(int64(pubKey.E)).Bytes()

	// 确保指数至少是3字节长
	if len(e) < 3 {
		padding := make([]byte, 3-len(e))
		e = append(padding, e...)
	}

	return JWK{
		Kid: "Nyauth-Server", // 密钥ID
		Kty: "RSA",           // 密钥类型
		Use: "sig",           // 用途：签名
		Alg: "RS256",         // 算法
		N:   untils.Base64URLEncode(nBytes),
		E:   untils.Base64URLEncode(e),
	}
}
