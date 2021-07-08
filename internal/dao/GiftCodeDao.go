package dao

import (
	"RedisGiftCode/internal/structInfo"
	"RedisGiftCode/internal/utils"
	"encoding/json"
	"time"
)

var receiveGiftList structInfo.ReceiveGiftList

func CreateGiftCodeDao(code string, jsonCodeInfo []byte, validPeriod int) (string, *structInfo.Response) {
	//以礼品吗为key存到Redis,并设置过期时间
	err := utils.Rdb.Set(code, jsonCodeInfo, time.Duration(validPeriod)*time.Hour).Err()
	if err != nil {
		return "", structInfo.RedisErr
	}
	return code, nil
}

func GetGiftCodeInfoDao(code string) (structInfo.GiftCodeInfo, *structInfo.Response) {

	CodeInfo := structInfo.GiftCodeInfo{}
	//根据礼品码查询礼品信息
	JsonCodeInfo, err1 := utils.Rdb.Get(code).Result()
	if err1 != nil {
		return CodeInfo, structInfo.RedisErr
	}
	//反序列化
	err := json.Unmarshal([]byte(JsonCodeInfo), &CodeInfo)
	if err != nil {
		return CodeInfo, structInfo.MarshalErr
	}
	return CodeInfo, nil
}

func VerifyFiftCodeDao(giftCodeInfo structInfo.GiftCodeInfo, user string) (structInfo.GiftCodeInfo, *structInfo.Response) {
	//领取数加一
	giftCodeInfo.ReceiveNum = giftCodeInfo.ReceiveNum + 1
	//用户添加到领取列表，保存到Redis
	receiveGiftList.ReceiveTime = time.Now()
	receiveGiftList.ReceiveUser = user
	giftCodeInfo.ReceiveList = append(giftCodeInfo.ReceiveList, receiveGiftList)
	code := giftCodeInfo.Code
	jsonCodeInfo, err1 := json.Marshal(giftCodeInfo)
	if err1 != nil {
		return giftCodeInfo, structInfo.MarshalErr
	}
	err := utils.Rdb.Set(code, jsonCodeInfo, utils.Rdb.TTL(code).Val())
	if err != nil {
		return giftCodeInfo, structInfo.RedisErr
	}
	return giftCodeInfo, nil
}
