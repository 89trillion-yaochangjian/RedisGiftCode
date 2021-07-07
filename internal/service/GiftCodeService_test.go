package service

import (
	"RedisGiftCode/StructInfo"
	"RedisGiftCode/internal/utils"
	"testing"
)

func TestCreateGiftCodeService(t *testing.T) {
	utils.InitClient()
	giftContent := StructInfo.GiftContentList{
		GoldCoins: 111,
		Diamonds:  222,
		Props:     333,
		Heroes:    444,
		Creeps:    555,
	}
	GiftCodeInfo := StructInfo.GiftCodeInfo{
		GiftDes:        "desc",
		AvailableTimes: 100000,
		ValidPeriod:    4,
		User:           "tom",
		ContentList:    giftContent,
	}
	code := CreateGiftCodeService(GiftCodeInfo)
	t.Log(code)
}

func TestGetGiftCodeInfoService(t *testing.T) {
	utils.InitClient()
	GiftInfo := GetGiftCodeInfoService("A4UJTDLV")
	t.Log(GiftInfo)
}

func TestVerifyFiftCodeService(t *testing.T) {
	utils.InitClient()
	ContentInfo := VerifyFiftCodeService("A4UJTDLV", "tom")
	t.Log(ContentInfo)
}
