package database

import (
	"fmt"

	"nyauth_backed/source"
	"nyauth_backed/source/logger"

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
