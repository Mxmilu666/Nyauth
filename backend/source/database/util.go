package database

import (
	"context"
	"fmt"
	"time"

	"nyauth_backed/source"
	"nyauth_backed/source/logger"
	"nyauth_backed/source/models"
	"nyauth_backed/source/untils"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var DatabaseName = "nyauth"
var UserCollection = "users"
var ClientCollection = "clients"

// 保存数据库连接
var client *mongo.Client

// InitDatabase 初始化数据库
func InitDatabase() error {
	// 初始化数据库
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d",
		source.AppConfig.Database.Username,
		source.AppConfig.Database.Password,
		source.AppConfig.Database.Host,
		source.AppConfig.Database.Port,
	)

	logger.Info("Try to use %v to connect to the MongoDB on %v:%v", source.AppConfig.Database.Username, source.AppConfig.Database.Host, source.AppConfig.Database.Port)

	var err error
	client, err = SetupDatabase(uri)
	if err != nil {
		return err
	}

	// 初始化用户集合
	err = EnsureCollection(client, DatabaseName, UserCollection)
	if err != nil {
		return err
	}

	// 初始化 Client 集合
	err = EnsureCollection(client, DatabaseName, ClientCollection)
	if err != nil {
		return err
	}

	return nil
}

// GetUserByUsername 通过用户名获取用户信息
func GetUserByUsername(username string) (bool, *models.DatabaseUser, error) {
	collection := client.Database(DatabaseName).Collection(UserCollection)

	filter := bson.M{
		"$or": []bson.M{
			{"user_name": username},
			{"user_email": username},
		},
	}
	// 查找一个匹配的文档
	var user models.DatabaseUser
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil, nil
		}
		return false, nil, err
	}

	return true, &user, nil
}

// CreateUser 注册新用户
func CreateUser(username, email, password, avatar string) error {
	collection := client.Database(DatabaseName).Collection(UserCollection)

	objectId := bson.NewObjectID()

	// 创建新用户对象
	user := &models.DatabaseUser{
		UserID:       objectId,
		UserUUID:     untils.ToUUIDv5(objectId.Hex()),
		Username:     username,
		UserEmail:    email,
		UserPassword: password,
		Avatar:       avatar,

		// 注册时间
		RegisterAt: bson.DateTime(time.Now().UnixNano() / int64(time.Millisecond)),
		IsBanned:   false,
		Role:       "0",
	}

	// 插入用户到数据库
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return nil
}

// GetUserByID 通过用户ID获取用户信息
func GetUserByID(userID string) (*models.DatabaseUser, error) {
	collection := client.Database(DatabaseName).Collection(UserCollection)

	// 将字符串类型的 userID 转换为 ObjectID
	objID, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	// 查找一个匹配的文档
	var user models.DatabaseUser
	err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// UpdateUser 通过用户ID更新用户信息
func UpdateUser(userID string, updates map[string]interface{}) error {
	collection := client.Database(DatabaseName).Collection(UserCollection)

	// 将字符串类型的 userID 转换为 ObjectID
	objID, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("invalid user ID: %w", err)
	}

	// 创建更新文档
	update := bson.M{"$set": updates}

	// 执行更新操作
	_, err = collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		update,
	)

	if err != nil {
		return fmt.Errorf("failed to update user information: %w", err)
	}

	return nil
}
