package StructInfo

type GiftContentList struct{
	GoldCoins int `json:"gold_coins"`   //金币
	//GoldCoinsCount int   //金币数量
	Diamonds int `json:"diamonds"`   //钻石
	//DiamondsCount int    //钻石数量
	Props int `json:"props"`	 //道具
	//PropsCount int		 //道具数量
	Heroes int `json:"heroes"`	 //英雄
	//HeroesCount int      //英雄数量
	Creeps int `json:"creeps"`    //小兵
	//CreepsCount int      //小兵数量
}