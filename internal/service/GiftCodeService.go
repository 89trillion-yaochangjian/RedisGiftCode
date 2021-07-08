package service

import (
	"RedisGiftCode/StructInfo"
	"RedisGiftCode/internal/dao"
	"errors"
	"fmt"
	"time"
)

var receiveGiftList StructInfo.ReceiveGiftList

//管理后台调用 - 创建礼品码

func CreateGiftCodeService(giftCodeInfo StructInfo.GiftCodeInfo) (string, error) {
	CodeInfo, err := dao.CreateGiftCodeDao(giftCodeInfo)
	if err != nil {
		err = errors.New("创建礼品码异常")
		return "", err
	}
	return CodeInfo, err
}

//管理后台调用 - 查询礼品码信息

func GetGiftCodeInfoService(code string) (StructInfo.GiftCodeInfo, error) {
	//根据礼品码查询礼品信息
	CodeInfo, err := dao.GetGiftCodeInfoDao(code)
	if err != nil {
		err = errors.New("礼品码查询异常")
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

func VerifyFiftCodeService(code string, user string) (StructInfo.GiftCodeInfo, error) {
	CodeInfo, err := dao.GetGiftCodeInfoDao(code)
	if err != nil {
		err = errors.New("礼品码查询异常")
		return CodeInfo, err
	}
	switch CodeInfo.CodeType {
	case -1:
		if CodeInfo.ReceiveNum != 1 || CodeInfo.User != user {
			fmt.Printf("礼包已经领取过")
		}
	case 0:
		if CodeInfo.AvailableTimes > CodeInfo.ReceiveNum {
			//领取数加一
			CodeInfo.ReceiveNum = CodeInfo.ReceiveNum + 1
			//用户添加到领取列表，保存到Redis
			receiveGiftList.ReceiveTime = time.Now()
			receiveGiftList.ReceiveUser = user
			CodeInfo.ReceiveList = append(CodeInfo.ReceiveList, receiveGiftList)
			dao.VerifyFiftCodeDao(CodeInfo)
		}
	case -2:
		//领取数加一
		CodeInfo.ReceiveNum = CodeInfo.ReceiveNum + 1
		//用户添加到领取列表
		receiveGiftList.ReceiveTime = time.Now()
		receiveGiftList.ReceiveUser = user
		CodeInfo.ReceiveList = append(CodeInfo.ReceiveList, receiveGiftList)
		dao.VerifyFiftCodeDao(CodeInfo)
	}
	return CodeInfo, nil
}
