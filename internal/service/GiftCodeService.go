package service

import (
	"RedisGiftCode/StructInfo"
	"RedisGiftCode/internal/utils"
	"encoding/json"
	"fmt"
	"time"
)
var receiveGiftList StructInfo.ReceiveGiftList
//管理后台调用 - 创建礼品码

func CreateGiftCodeService(giftCodeInfo StructInfo.GiftCodeInfo) (string) {
	code := utils.GetGiftCodeUtil()
	giftCodeInfo.Code = code
	//设置创建时间
	giftCodeInfo.CreatTime = time.Now()
	//设置过期时间
	validPeriod := giftCodeInfo.ValidPeriod
	jsonCodeInfo,err1 :=json.Marshal(giftCodeInfo)
	if err1!=nil {
		fmt.Printf("Serialization giftCodeInfo failed, err:%v\n", err1)
	}
	//以礼品吗为key存到Redis,并设置过期时间
	err := utils.Rdb.Set(code, jsonCodeInfo, time.Duration(validPeriod)*time.Hour).Err()
	utils.Rdb.Do()
	if err != nil {
		fmt.Printf("create giftCode failed, err:%v\n", err)
	}
	return code
}


//管理后台调用 - 查询礼品码信息

func GetGiftCodeInfoService(code string) (StructInfo.GiftCodeInfo){
	if len(code) != 8{
		fmt.Printf("礼品码长度不为8")
		return StructInfo.GiftCodeInfo{}
	}
	//根据礼品码查询礼品信息
	JsonCodeInfo, err1 :=utils.Rdb.Get(code).Result()
	if err1 != nil {
		fmt.Printf("get giftCodeInfo failed, err:%v\n", err1)
	}
	CodeInfo := StructInfo.GiftCodeInfo{}
	//反序列化
	UnmarshalErr := json.Unmarshal([]byte(JsonCodeInfo),&CodeInfo)
	if UnmarshalErr != nil {
		fmt.Printf("Deserialization giftCodeInfo failed, err:%v\n", UnmarshalErr)
	}
	codeType :=CodeInfo.CodeType
	if codeType > 0 {
		CodeInfo.CodeTypeDesc = "不指定用户限制兑换次数"
	}else if codeType == -1 {
		CodeInfo.CodeTypeDesc = "指定用户一次性消耗"
	}else if codeType == -2 {
		CodeInfo.CodeTypeDesc = "不限用户不限次数兑换"
	}
    return CodeInfo
}


//客户端调用 - 验证礼品码

func VerifyFiftCodeService(code string,user string) (StructInfo.GiftCodeInfo)  {
	if len(code) != 8{
		fmt.Printf("礼品码长度不为8")
		return StructInfo.GiftCodeInfo{}
	}
	JsonCodeInfo, GteInfoRrr :=utils.Rdb.Get(code).Result()
	if GteInfoRrr != nil {
		fmt.Printf("code无效或已过期, err:%v\n", GteInfoRrr)
	}

	CodeInfo := StructInfo.GiftCodeInfo{}
	UnmarshalErr := json.Unmarshal([]byte(JsonCodeInfo),&CodeInfo)
	if UnmarshalErr != nil {
		fmt.Printf("Deserialization giftCodeInfo failed, err:%v\n", UnmarshalErr)
	}
	//获取当前客户
	//user := "tom"
	switch CodeInfo.CodeType {
	case -1:
		if CodeInfo.ReceiveNum != 1||CodeInfo.User!=user{
			fmt.Printf("礼包已经领取过")
		}
	case 0:
		if CodeInfo.AvailableTimes>CodeInfo.ReceiveNum{
			//领取数加一
			CodeInfo.ReceiveNum = CodeInfo.ReceiveNum + 1
			//用户添加到领取列表
			receiveGiftList.ReceiveTime = time.Now()
			receiveGiftList.ReceiveUser = user
			CodeInfo.ReceiveList = append(CodeInfo.ReceiveList,receiveGiftList)
			jsonCodeInfo,err1 :=json.Marshal(CodeInfo)
			if err1!=nil {
				fmt.Printf("Serialization giftCodeInfo failed, err:%v\n", err1)
			}
			time := utils.Rdb.TTL(code)
			utils.Rdb.Set(code, jsonCodeInfo,time.Val())
		}
	case -2:
		//领取数加一
		CodeInfo.ReceiveNum = CodeInfo.ReceiveNum + 1
		//用户添加到领取列表
		receiveGiftList.ReceiveTime = time.Now()
		receiveGiftList.ReceiveUser = user
		CodeInfo.ReceiveList = append(CodeInfo.ReceiveList,receiveGiftList)
		jsonCodeInfo,err1 :=json.Marshal(CodeInfo)
		if err1!=nil {
			fmt.Printf("Serialization giftCodeInfo failed, err:%v\n", err1)
		}
		a := utils.Rdb.TTL(code)
		utils.Rdb.Set(code, jsonCodeInfo,a.Val())
	}
	return CodeInfo
}