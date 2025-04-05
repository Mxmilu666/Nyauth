package handles

import (
	"crypto/rsa"
	"math/big"
	"net/http"
	"nyauth_backed/source"
	"nyauth_backed/source/helper"
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

// OIDCConfiguration 表示 OpenID Connect 配置
type OIDCConfiguration struct {
	Issuer                            string   `json:"issuer"`
	AuthorizationEndpoint             string   `json:"authorization_endpoint"`
	TokenEndpoint                     string   `json:"token_endpoint"`
	UserinfoEndpoint                  string   `json:"userinfo_endpoint"`
	JwksURI                           string   `json:"jwks_uri"`
	RegistrationEndpoint              string   `json:"registration_endpoint,omitempty"`
	ScopesSupported                   []string `json:"scopes_supported"`
	ResponseTypesSupported            []string `json:"response_types_supported"`
	ResponseModesSupported            []string `json:"response_modes_supported,omitempty"`
	GrantTypesSupported               []string `json:"grant_types_supported,omitempty"`
	SubjectTypesSupported             []string `json:"subject_types_supported"`
	IDTokenSigningAlgValuesSupported  []string `json:"id_token_signing_alg_values_supported"`
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported,omitempty"`
	ClaimsSupported                   []string `json:"claims_supported"`
}

// GetOpenIDConfiguration 处理 /.well-known/openid-configuration 请求
// 返回OpenID Connect提供者的配置信息
func GetOpenIDConfiguration(c *gin.Context) {
	baseURL := source.AppConfig.Server.BaseURL

	config := OIDCConfiguration{
		Issuer:                            baseURL,
		AuthorizationEndpoint:             baseURL + "/oauth/authorize",
		TokenEndpoint:                     baseURL + "/api/v0/oauth/token",
		UserinfoEndpoint:                  baseURL + "/api/v0/account/info",
		JwksURI:                           baseURL + "/.well-known/jwks.json",
		ScopesSupported:                   []string{"openid", "profile", "email"},
		ResponseTypesSupported:            []string{"code"},
		SubjectTypesSupported:             []string{"public"},
		IDTokenSigningAlgValuesSupported:  []string{"RS256"},
		TokenEndpointAuthMethodsSupported: []string{"client_secret_post", "client_secret_basic"},
		ClaimsSupported:                   []string{"sub", "iss", "name", "email", "picture"},
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
		Kid: "default-key", // 密钥ID，在生产环境中应该使用唯一的ID
		Kty: "RSA",         // 密钥类型
		Use: "sig",         // 用途：签名
		Alg: "RS256",       // 算法
		N:   untils.Base64URLEncode(nBytes),
		E:   untils.Base64URLEncode(e),
	}
}
