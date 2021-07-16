package service

import (
	model2 "RedisGiftCode/internal/model"
	"RedisGiftCode/internal/utils"
	"testing"
)

func TestCreateGiftCodeService(t *testing.T) {
	utils.InitClient()
	giftContent := model2.GiftContentList{
		GoldCoins: 111,
		Diamonds:  222,
		Props:     333,
		Heroes:    444,
		Creeps:    555,
	}
	GiftCodeInfo := model2.GiftCodeInfo{
		GiftDes:        "desc",
		AvailableTimes: 100000,
		ValidPeriod:    4,
		User:           "tom",
		ContentList:    giftContent,
	}
	code, err := CreateGiftCodeService(GiftCodeInfo)
	t.Log(code, err)
}

func TestGetGiftCodeInfoService(t *testing.T) {
	utils.InitClient()
	GiftInfo, err := GetGiftCodeInfoService("A4UJTDLV")
	t.Log(GiftInfo, err)
}

func TestVerifyFiftCodeService(t *testing.T) {
	utils.InitClient()
	ContentInfo, err := VerifyFiftCodeService("A4UJTDLV", "tom")
	t.Log(ContentInfo, err)
}
