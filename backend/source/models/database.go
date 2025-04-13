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
	Description  string        `bson:"description"`
	Avatar       string        `bson:"avatar"`
	ClientSecret string        `bson:"client_secret"`
	RedirectURI  string        `bson:"redirect_uri"`
	Permissions  []string      `bson:"permissions"`
	Status       int           `bson:"status"`
	CreatedBy    string        `bson:"createdBy"`
	CreatedAt    bson.DateTime `bson:"created_at"`
	UpdatedAt    bson.DateTime `bson:"updated_at"`
}

// identity 集合中的文档结构 (用户的多身份)
type DatabaseUserIdentity struct {
	UserID      bson.ObjectID `bson:"_id"`
	UserUUID    string        `bson:"user_uuid"`
	UserEmail   string        `bson:"user_email"`
	Attributed  string        `bson:"attributed"`   // 归属的用户ID
	DisplayName string        `bson:"display_name"` // 显示名称
	Avatar      string        `bson:"avatar"`       // 身份头像
	Description string        `bson:"description"`  // 身份描述
	CreatedAt   bson.DateTime `bson:"created_at"`
	UpdatedAt   bson.DateTime `bson:"updated_at"`
}
