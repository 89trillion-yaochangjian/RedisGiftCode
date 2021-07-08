package service

import (
	"RedisGiftCode/internal/dao"
	"RedisGiftCode/internal/structInfo"
	"RedisGiftCode/internal/utils"
	"encoding/json"
	"errors"
	"time"
)

//管理后台调用 - 创建礼品码

func CreateGiftCodeService(giftCodeInfo structInfo.GiftCodeInfo) (string, error) {
	code := utils.GetGiftCodeUtil()
	giftCodeInfo.Code = code
	//设置创建时间
	giftCodeInfo.CreatTime = time.Now()
	//设置过期时间
	validPeriod := giftCodeInfo.ValidPeriod
	jsonCodeInfo, err1 := json.Marshal(giftCodeInfo)
	if err1 != nil {
		return "", err1
	}
	CodeInfo, err := dao.CreateGiftCodeDao(code, jsonCodeInfo, validPeriod)
	if err != nil {
		return "", err
	}
	return CodeInfo, err
}

//管理后台调用 - 查询礼品码信息

func GetGiftCodeInfoService(code string) (structInfo.GiftCodeInfo, error) {
	//根据礼品码查询礼品信息
	CodeInfo, err := dao.GetGiftCodeInfoDao(code)
	if err != nil {
		//err = errors.New("礼品码查询异常")
		return CodeInfo, err
	}
	//显示礼包类型
	codeType := CodeInfo.CodeType
	if codeType > 0 {
		CodeInfo.CodeTypeDesc = "不指定用户限制兑换次数"
	} else if codeType == -1 {
		CodeInfo.CodeTypeDesc = "指定用户一次性消耗"
	} else if codeType == -2 {
		CodeInfo.CodeTypeDesc = "不限用户不限次数兑换"
	}
	return CodeInfo, err
}

//客户端调用 - 验证礼品码

func VerifyFiftCodeService(code string, user string) (structInfo.GiftCodeInfo, error) {
	CodeInfo, err := dao.GetGiftCodeInfoDao(code)
	if err != nil {
		return CodeInfo, err
	}
	switch CodeInfo.CodeType {
	case -1:
		if CodeInfo.ReceiveNum == 1 || CodeInfo.User != user {
			err = errors.New("礼包码已经领取过了")
			return CodeInfo, err
		}
		dao.VerifyFiftCodeDao(CodeInfo, user)
	case 0:
		if CodeInfo.AvailableTimes > CodeInfo.ReceiveNum {
			dao.VerifyFiftCodeDao(CodeInfo, user)
		}
	case -2:
		dao.VerifyFiftCodeDao(CodeInfo, user)
	}
	return CodeInfo, nil
}
