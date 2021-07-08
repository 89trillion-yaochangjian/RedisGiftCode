package dao

import (
	"RedisGiftCode/StructInfo"
	"RedisGiftCode/internal/utils"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

func CreateGiftCodeDao(giftCodeInfo StructInfo.GiftCodeInfo) (string, error) {
	code := utils.GetGiftCodeUtil()
	giftCodeInfo.Code = code
	//设置创建时间
	giftCodeInfo.CreatTime = time.Now()
	//设置过期时间
	validPeriod := giftCodeInfo.ValidPeriod
	jsonCodeInfo, err1 := json.Marshal(giftCodeInfo)
	if err1 != nil {
		err1 = errors.New("序列化异常")
		return "", err1
	}
	//以礼品吗为key存到Redis,并设置过期时间
	err := utils.Rdb.Set(code, jsonCodeInfo, time.Duration(validPeriod)*time.Hour).Err()
	if err != nil {
		err1 = errors.New("redis存储异常")
		return "", err
	}
	return code, nil
}

func GetGiftCodeInfoDao(code string) (StructInfo.GiftCodeInfo, error) {
	if len(code) != 8 {
		fmt.Printf("礼品码长度不为8")
	}
	CodeInfo := StructInfo.GiftCodeInfo{}
	//根据礼品码查询礼品信息
	JsonCodeInfo, err1 := utils.Rdb.Get(code).Result()
	if err1 != nil {
		err1 = errors.New("获取礼包信息失败")
		return CodeInfo, err1
	}

	//反序列化
	UnmarshalErr := json.Unmarshal([]byte(JsonCodeInfo), &CodeInfo)
	if UnmarshalErr != nil {
		UnmarshalErr = errors.New("序列化失败")
		return CodeInfo, UnmarshalErr
	}
	return CodeInfo, err1
}

func VerifyFiftCodeDao(giftCodeInfo StructInfo.GiftCodeInfo) (StructInfo.GiftCodeInfo, error) {
	code := giftCodeInfo.Code
	jsonCodeInfo, err1 := json.Marshal(giftCodeInfo)
	if err1 != nil {
		err1 = errors.New("序列化失败")
		return giftCodeInfo, err1
	}
	utils.Rdb.Set(code, jsonCodeInfo, utils.Rdb.TTL(code).Val())
	return giftCodeInfo, nil
}
