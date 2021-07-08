package handler

import (
	"RedisGiftCode/internal/service"
	"RedisGiftCode/internal/structInfo"
)

//管理后台调用 - 创建礼品码

func CreateGiftCodeHandler(giftCodeInfo structInfo.GiftCodeInfo) (string, *structInfo.Response) {
	codeInfo, err := service.CreateGiftCodeService(giftCodeInfo)
	if err != nil {
		return codeInfo, err
	}
	return codeInfo, nil
}

//管理后台调用 - 查询礼品码信息

func GetFiftCodeInfoHandler(code string) (structInfo.GiftCodeInfo, *structInfo.Response) {
	giftCodeInfo, err := service.GetGiftCodeInfoService(code)
	if err != nil {
		return giftCodeInfo, err
	}
	return giftCodeInfo, nil
}

//客户端调用 - 验证礼品码

func VerifyFiftCodeHandler(code string, user string) (structInfo.GiftCodeInfo, *structInfo.Response) {
	giftCodeInfo, err := service.VerifyFiftCodeService(code, user)
	if err != nil {
		return giftCodeInfo, err
	}
	return giftCodeInfo, nil
}
