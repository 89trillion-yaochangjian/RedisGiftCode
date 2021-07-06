package handler

import (
	"RedisGiftCode/StructInfo"
	"RedisGiftCode/internal/service"
)


//管理后台调用 - 创建礼品码

func CreateGiftCodeHandler(giftCodeInfo StructInfo.GiftCodeInfo) (string) {
	codeRes := service.CreateGiftCodeService(giftCodeInfo)
	return codeRes
}


//管理后台调用 - 查询礼品码信息

func GetFiftCodeInfoHandler(code string) (StructInfo.GiftCodeInfo){
	giftCodeInfo := service.GetGiftCodeInfoService(code)
	return giftCodeInfo
}


//客户端调用 - 验证礼品码

func VerifyFiftCodeHandler(code string,user string) (StructInfo.GiftCodeInfo)  {
	giftCodeInfo := service.VerifyFiftCodeService(code,user)
	return giftCodeInfo
}