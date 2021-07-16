package model

type GiftContentList struct {
	GoldCoins int `json:"gold_coins"` //金币
	Diamonds  int `json:"diamonds"`   //钻石
	Props     int `json:"props"`      //道具
	Heroes    int `json:"heroes"`     //英雄
	Creeps    int `json:"creeps"`     //小兵
}
