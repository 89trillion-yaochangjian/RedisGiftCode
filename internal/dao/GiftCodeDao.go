package dao

import (
	"RedisGiftCode/internal/structInfo"
	"RedisGiftCode/internal/utils"
	"encoding/json"
	"time"
)

var receiveGiftList structInfo.ReceiveGiftList

func CreateGiftCodeDao(code string, jsonCodeInfo []byte, validPeriod int) (string, error) {
	//以礼品吗为key存到Redis,并设置过期时间
	err := utils.Rdb.Set(code, jsonCodeInfo, time.Duration(validPeriod)*time.Hour).Err()
	if err != nil {
		//err1 = errors.New("redis存储异常")
		return "", err
	}
	return code, nil
}

func GetGiftCodeInfoDao(code string) (structInfo.GiftCodeInfo, error) {

	CodeInfo := structInfo.GiftCodeInfo{}
	//根据礼品码查询礼品信息
	JsonCodeInfo, err1 := utils.Rdb.Get(code).Result()
	if err1 != nil {
		//err1 = errors.New("获取礼包信息失败")
		return CodeInfo, err1
	}
	//反序列化
	UnmarshalErr := json.Unmarshal([]byte(JsonCodeInfo), &CodeInfo)
	if UnmarshalErr != nil {
		return CodeInfo, UnmarshalErr
	}
	return CodeInfo, err1
}

func VerifyFiftCodeDao(giftCodeInfo structInfo.GiftCodeInfo, user string) (structInfo.GiftCodeInfo, error) {
	//领取数加一
	giftCodeInfo.ReceiveNum = giftCodeInfo.ReceiveNum + 1
	//用户添加到领取列表，保存到Redis
	receiveGiftList.ReceiveTime = time.Now()
	receiveGiftList.ReceiveUser = user
	giftCodeInfo.ReceiveList = append(giftCodeInfo.ReceiveList, receiveGiftList)
	code := giftCodeInfo.Code
	jsonCodeInfo, err1 := json.Marshal(giftCodeInfo)
	if err1 != nil {
		return giftCodeInfo, err1
	}
	utils.Rdb.Set(code, jsonCodeInfo, utils.Rdb.TTL(code).Val())
	return giftCodeInfo, nil
}
