package database

import (
	"context"
	"fmt"
	"nyauth_backed/source/models"
	"nyauth_backed/source/untils"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// CreateUserIdentity 创建用户多身份
func CreateUserIdentity(userID, email, displayName, description, avatar string) (string, error) {
	collection := client.Database(DatabaseName).Collection(MultiUserCollection)

	objectId := bson.NewObjectID()
	userUUID := untils.ToUUIDv5(objectId.Hex())
	now := bson.DateTime(time.Now().UnixNano() / int64(time.Millisecond))

	// 创建新的多身份对象
	identity := &models.DatabaseUserIdentity{
		UserID:      objectId,
		UserUUID:    userUUID,
		UserEmail:   email,
		Attributed:  userID,
		DisplayName: displayName,
		Avatar:      avatar,
		Description: description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// 插入多身份到数据库
	_, err := collection.InsertOne(context.TODO(), identity)
	if err != nil {
		return "", err
	}

	return objectId.Hex(), nil
}

// GetUserIdentities 获取用户的所有多身份
func GetUserIdentities(userID string) ([]map[string]interface{}, error) {
	collection := client.Database(DatabaseName).Collection(MultiUserCollection)

	// 初始化一个空数组，确保即使没有记录也会返回空数组而不是 null
	identities := []map[string]interface{}{}

	cursor, err := collection.Find(context.TODO(), bson.M{"attributed": userID})
	if err != nil {
		return identities, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var identity models.DatabaseUserIdentity
		if err := cursor.Decode(&identity); err != nil {
			return identities, err
		}

		// 转换为前端友好的格式
		identityMap := map[string]interface{}{
			"identity_id":  identity.UserID.Hex(),
			"uuid":         identity.UserUUID,
			"email":        identity.UserEmail,
			"display_name": identity.DisplayName,
			"avatar":       identity.Avatar,
			"description":  identity.Description,
			"created_at":   identity.CreatedAt,
			"is_primary":   false, // 默认为非主账号
		}

		identities = append(identities, identityMap)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// 使用 GetUserByID 函数查找主账号信息
	primaryUser, err := GetUserByID(userID)
	if err == nil && primaryUser != nil {
		// 主账号信息
		primaryIdentity := map[string]interface{}{
			"identity_id":  primaryUser.UserID.Hex(),
			"uuid":         primaryUser.UserUUID,
			"email":        primaryUser.UserEmail,
			"display_name": primaryUser.Username,
			"avatar":       primaryUser.Avatar,
			"description":  "主账号", // 标记为主账号
			"created_at":   primaryUser.RegisterAt,
			"is_primary":   true,
		}
		// 将主账号放在列表最前面
		identities = append([]map[string]interface{}{primaryIdentity}, identities...)
	}

	return identities, nil
}

// GetIdentityByID 通过身份ID获取身份信息
func GetIdentityByID(identityID string) (*models.DatabaseUserIdentity, error) {
	collection := client.Database(DatabaseName).Collection(MultiUserCollection)

	// 将字符串类型的 identityID 转换为 ObjectID
	objID, err := bson.ObjectIDFromHex(identityID)
	if err != nil {
		return nil, err
	}

	// 查找一个匹配的文档
	var identity models.DatabaseUserIdentity
	err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&identity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &identity, nil
}

// UpdateIdentity 更新用户多身份信息
func UpdateIdentity(identityID string, updates map[string]interface{}) error {
	collection := client.Database(DatabaseName).Collection(MultiUserCollection)

	// 将字符串类型的 identityID 转换为 ObjectID
	objID, err := bson.ObjectIDFromHex(identityID)
	if err != nil {
		return fmt.Errorf("无效的身份ID: %w", err)
	}

	// 添加更新时间
	if updates == nil {
		updates = make(map[string]interface{})
	}
	updates["updated_at"] = bson.DateTime(time.Now().UnixNano() / int64(time.Millisecond))

	// 创建更新文档
	update := bson.M{"$set": updates}

	// 执行更新操作
	_, err = collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		update,
	)

	if err != nil {
		return fmt.Errorf("更新身份信息失败: %w", err)
	}

	return nil
}
