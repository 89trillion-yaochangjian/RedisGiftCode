package model

type GiftCodeInfo struct {
	GiftDes        string            `json:"gift_des"`  //礼品码描述
	CodeType       int               `json:"code_type"` // >n（n>0）--有n次 不限定用户，限定领取次数   -1--指定用户一次领取 -2--不限定用户，不限定次数
	CodeTypeDesc   string            `json:"code_type_desc"`
	ReceiveNum     int64             `json:"receive_num"`     //已经领取次数
	AvailableTimes int64             `json:"available_times"` //可以领取的次数
	ValidPeriod    int               `json:"valid_period"`    //有效期
	OverTime       int64             `json:"over_time"`       //过期时间
	Contents       int               `json:"contents"`        //礼包内容
	Code           string            `json:"code"`            //礼包码
	Creator        string            `json:"creator"`         //创建人
	CreatTime      int64             `json:"creat_time"`      //创建时间
	User           string            `json:"user"`            //指定用户
	ContentList    GiftContentList   //礼品内容列表
	ReceiveList    []ReceiveGiftList //领取列表
}
