package handler

import (
	"RedisGiftCode/StructInfo"
	"RedisGiftCode/internal/service"
	"errors"
)

//管理后台调用 - 创建礼品码

func CreateGiftCodeHandler(giftCodeInfo StructInfo.GiftCodeInfo) (StructInfo.MesInfo, error) {
	codeRes, err := service.CreateGiftCodeService(giftCodeInfo)
	if err != nil {
		return StructInfo.MesInfo{Msg: "创建礼包码失败", ER: err}, nil
	}
	return StructInfo.MesInfo{Msg: "创建礼包码成功", Data: codeRes}, nil
}

//管理后台调用 - 查询礼品码信息

func GetFiftCodeInfoHandler(code string) (StructInfo.MesInfo, error) {
	giftCodeInfo, err := service.GetGiftCodeInfoService(code)
	if err != nil {
		return StructInfo.MesInfo{Msg: "礼品码无效或已过期", ER: err}, nil
	}
	return StructInfo.MesInfo{Msg: "查询礼品码信息成功", Data: giftCodeInfo}, nil
}

//客户端调用 - 验证礼品码

func VerifyFiftCodeHandler(code string, user string) (StructInfo.GiftCodeInfo, error) {
	giftCodeInfo, err := service.VerifyFiftCodeService(code, user)
	if err != nil {
		err = errors.New("礼品码无效或已过期")
		return giftCodeInfo, err
	}
	return giftCodeInfo, nil
}
