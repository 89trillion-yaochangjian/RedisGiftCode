package dao

import (
	"RedisGiftCode/internal/config"
	"RedisGiftCode/internal/model"
	"RedisGiftCode/internal/status"
	"encoding/json"
	"time"
)

var receiveGiftList model.ReceiveGiftList

//创建礼品码

func CreateGiftCodeDao(code string, jsonCodeInfo []byte, validPeriod int) (string, *status.Response) {
	//以礼品吗为key存到Redis,并设置过期时间
	err := config.Rdb.Set(code, jsonCodeInfo, time.Duration(validPeriod)*time.Hour).Err()
	if err != nil {
		return "", status.RedisErr
	}
	return code, nil
}

//管理后台调用 - 查询礼品码信息

func GetGiftCodeInfoDao(code string) (model.GiftCodeInfo, *status.Response) {

	CodeInfo := model.GiftCodeInfo{}
	//根据礼品码查询礼品信息
	JsonCodeInfo, err1 := config.Rdb.Get(code).Result()
	if err1 != nil {
		return CodeInfo, status.CodeTimeOver
	}
	//反序列化
	err := json.Unmarshal([]byte(JsonCodeInfo), &CodeInfo)
	if err != nil {
		return CodeInfo, status.MarshalErr
	}
	return CodeInfo, nil
}

//客户端调用 - 验证礼品码

func VerifyFiftCodeDao(giftCodeInfo model.GiftCodeInfo, user string) (model.GiftCodeInfo, *status.Response) {
	//用户添加到领取列表，保存到Redis
	receiveGiftList.ReceiveTime = time.Now()
	receiveGiftList.ReceiveUser = user
	giftCodeInfo.ReceiveList = append(giftCodeInfo.ReceiveList, receiveGiftList)
	code := giftCodeInfo.Code
	jsonCodeInfo, err1 := json.Marshal(giftCodeInfo)
	if err1 != nil {
		return giftCodeInfo, status.MarshalErr
	}
	err := config.Rdb.Set(code, jsonCodeInfo, config.Rdb.TTL(code).Val())
	if err != nil {
		return giftCodeInfo, status.RedisErr
	}
	//领取数加一
	count := config.Rdb.Incr(giftCodeInfo.Code + "count")
	giftCodeInfo.ReceiveNum = count.Val()
	return giftCodeInfo, nil
}
