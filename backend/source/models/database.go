package models

import "go.mongodb.org/mongo-driver/v2/bson"

// user 集合中的文档结构
type DatabaseUser struct {
	UserID       bson.ObjectID `bson:"_id"`
	UserUUID     string        `bson:"user_uuid"`
	UserPassword string        `bson:"user_pass"`
	Username     string        `bson:"user_name"`
	UserEmail    string        `bson:"user_email"`
	Avatar       string        `bson:"avatar"`
	RegisterAt   bson.DateTime `bson:"register_at"`
	UpdatedAt    bson.DateTime `bson:"updated_at"`
	IsBanned     bool          `bson:"is_banned"`
	Role         string        `bson:"role"`
}

// client 集合中的文档结构
type DatabaseClient struct {
	ID           bson.ObjectID `bson:"_id"`
	ClientName   string        `bson:"client_name"`
	Avatar       string        `bson:"avatar"`
	ClientSecret string        `bson:"client_secret"`
	RedirectURI  string        `bson:"redirect_uri"`
	Permissions  []string      `bson:"permissions"`
	Status       int           `bson:"status"`
	CreatedBy    string        `bson:"createdBy"`
	CreatedAt    bson.DateTime `bson:"created_at"`
	UpdatedAt    bson.DateTime `bson:"updated_at"`
}

// authorization 集合中的文档结构
type DatabaseUserAuthorization struct {
	ID          bson.ObjectID `bson:"_id"`
	UserID      string        `bson:"user_id"`      // 授权用户的UUID
	ClientID    string        `bson:"client_id"`    // 被授权的客户端ID
	Scope       []string      `bson:"scope"`        // 授权范围
	AccessToken string        `bson:"access_token"` // 访问令牌
	TokenType   string        `bson:"token_type"`   // 令牌类型
	ExpiresAt   bson.DateTime `bson:"expires_at"`   // 令牌过期时间
	CreatedAt   bson.DateTime `bson:"created_at"`   // 授权创建时间
	UpdatedAt   bson.DateTime `bson:"updated_at"`   // 授权更新时间
}
