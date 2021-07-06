package service

import (
	"RedisGiftCode/StructInfo"
	"testing"
)

func TestCreateGiftCodeService(t *testing.T) {
	giftContent := StructInfo.GiftContentList{
		GoldCoins:111,
		Diamonds:222,
		Props:333,
		Heroes:444,
		Creeps:555,
	}
	GiftCodeInfo := StructInfo.GiftCodeInfo{
		GiftDes:"desc",
		AvailableTimes:100000,
		ValidPeriod:4,
		User: "tom",
		ContentList:giftContent,
	}
	code := CreateGiftCodeService(GiftCodeInfo)
	t.Log(code)
}

func TestGetGiftCodeInfoService(t *testing.T) {
	GiftInfo := GetGiftCodeInfoService("JI310XOC")
	t.Log(GiftInfo)
}

func TestVerifyFiftCodeService(t *testing.T) {
	ContentInfo := VerifyFiftCodeService("JI310XOC","tom")
	t.Log(ContentInfo)
}