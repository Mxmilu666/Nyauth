package database

import (
	"context"
	"fmt"
	"nyauth_backed/source/models"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// UserHasTOTP 检查用户是否启用了TOTP
func UserHasTOTP(userID string) (bool, error) {
	// 将字符串类型的 userID 转换为 ObjectID
	objID, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return false, err
	}

	collection := client.Database(DatabaseName).Collection(UserCollection)

	var result models.DatabaseUser
	err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&result)
	if err != nil {
		return false, err
	}

	return result.TOTPEnabled && result.TOTPSecret != "", nil
}

// EnableTOTP 启用TOTP并保存密钥
func EnableTOTP(userID, secret string) error {
	// 将字符串类型的 userID 转换为 ObjectID
	objID, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	collection := client.Database(DatabaseName).Collection(UserCollection)
	update := bson.M{
		"$set": bson.M{
			"totp_enabled": true,
			"totp_secret":  secret,
			"updated_at":   bson.DateTime(time.Now().UnixNano() / int64(time.Millisecond)),
		},
	}

	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	return err
}

// DisableTOTP 禁用TOTP
func DisableTOTP(userID string) error {
	// 将字符串类型的 userID 转换为 ObjectID
	objID, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	collection := client.Database(DatabaseName).Collection(UserCollection)
	update := bson.M{
		"$set": bson.M{
			"totp_enabled": false,
			"totp_secret":  "",
			"updated_at":   bson.DateTime(time.Now().UnixNano() / int64(time.Millisecond)),
		},
		"$unset": bson.M{
			"recovery_codes": "",
		},
	}

	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	return err
}

// GetUserTOTPSecret 获取用户的TOTP密钥
func GetUserTOTPSecret(userID string) (bool, string, error) {
	// 将字符串类型的 userID 转换为 ObjectID
	objID, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return false, "", err
	}

	collection := client.Database(DatabaseName).Collection(UserCollection)

	var result models.DatabaseUser
	err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&result)
	if err != nil {
		return false, "", err
	}

	return result.TOTPEnabled, result.TOTPSecret, nil
}

// SaveRecoveryCodes 保存恢复码
func SaveRecoveryCodes(userID string, codes []string) error {
	// 将字符串类型的 userID 转换为 ObjectID
	objID, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	collection := client.Database(DatabaseName).Collection(UserCollection)
	update := bson.M{
		"$set": bson.M{
			"recovery_codes": codes,
			"updated_at":     bson.DateTime(time.Now().UnixNano() / int64(time.Millisecond)),
		},
	}

	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	return err
}

// ValidateAndConsumeRecoveryCode 验证并使用恢复码
func ValidateAndConsumeRecoveryCode(userID, code string) (bool, error) {
	// 将字符串类型的 userID 转换为 ObjectID
	objID, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return false, err
	}

	collection := client.Database(DatabaseName).Collection(UserCollection)
	var result struct {
		RecoveryCodes []string `bson:"recovery_codes"`
	}

	err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&result)
	if err != nil {
		return false, err
	}

	// 检查恢复码是否存在
	foundIndex := -1
	for i, c := range result.RecoveryCodes {
		if c == code {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		return false, nil
	}

	// 使用恢复码（从列表中移除）
	newCodes := append(result.RecoveryCodes[:foundIndex], result.RecoveryCodes[foundIndex+1:]...)
	update := bson.M{
		"$set": bson.M{
			"recovery_codes": newCodes,
			"updated_at":     bson.DateTime(time.Now().UnixNano() / int64(time.Millisecond)),
		},
	}

	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	if err != nil {
		return false, fmt.Errorf("使用恢复码失败: %w", err)
	}

	return true, nil
}
