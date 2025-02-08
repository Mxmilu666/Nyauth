package database

import (
	"context"
	"nyauth_backed/source/logger"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// user 集合中的文档结构
type User struct {
	UserID       bson.ObjectID `bson:"_id"`
	UserPassword string        `bson:"user_pass"`
	Username     string        `bson:"user_name"`
	UserEmail    string        `bson:"user_email"`
	RegisterAt   bson.DateTime `bson:"register_at"`
	IsBanned     bool          `bson:"is_banned"`
	Role         string        `bson:"role"`
}

// SetupDatabase 连接到 MongoDB
func SetupDatabase(uri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, err
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	logger.Info("Connected to MongoDB")
	return client, nil
}

// EnsureCollection 确保指定的集合存在
func EnsureCollection(client *mongo.Client, dbName, collectionName string) error {
	collectionNames, err := client.Database(dbName).ListCollectionNames(context.TODO(), bson.M{})
	if err != nil {
		return err
	}

	// 检查集合是否存在
	collectionExists := false
	for _, name := range collectionNames {
		if name == collectionName {
			collectionExists = true
			break
		}
	}

	// 如果集合不存在，则创建集合
	if !collectionExists {
		err := client.Database(dbName).CreateCollection(context.TODO(), collectionName)
		if err != nil {
			return err
		}
		logger.Debug("Collection %s created successfully", collectionName)
	} else {
		logger.Debug("Collection %s already exists. Skip", collectionName)
	}

	return nil
}

// GetUserByUsername 通过用户名获取用户信息
func GetUserByUsername(username string) (bool, *User, error) {
	collection := client.Database(DatabaseName).Collection(UserCollection)

	filter := bson.M{
		"$or": []bson.M{
			{"user_name": username},
			{"user_email": username},
		},
	}
	// 查找一个匹配的文档
	var user User
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil, nil
		}
		return false, nil, err
	}

	return true, &user, nil
}
