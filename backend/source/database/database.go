package database

import (
	"context"
	"nyauth_backed/source/logger"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

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
