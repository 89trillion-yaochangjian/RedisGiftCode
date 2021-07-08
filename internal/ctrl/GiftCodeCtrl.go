package ctrl

import (
	"RedisGiftCode/StructInfo"
	"RedisGiftCode/internal/handler"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateGiftCode(c *gin.Context) {
	//获取参数
	info, err1 := c.GetRawData()
	if err1 != nil {
		c.JSON(http.StatusOK, StructInfo.MesInfo{Msg: "获取参数失败", Data: err1})
	}
	var giftCodeInfo StructInfo.GiftCodeInfo
	json.Unmarshal(info, &giftCodeInfo)
	//var giftCodeInfo = StructInfo.GiftCodeInfo{}
	//c.ShouldBind(&giftCodeInfo)
	//调用Handler
	code, err := handler.CreateGiftCodeHandler(giftCodeInfo)
	if err != nil {
		c.JSON(http.StatusOK, StructInfo.MesInfo{Msg: "创建礼包码失败", ER: err})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
	})
}

func GetGiftCodeInfoCtrl(c *gin.Context) {
	//获取参数
	code := c.Query("code")
	info, err := handler.GetFiftCodeInfoHandler(code)
	if err != nil {
		c.JSON(http.StatusOK, StructInfo.MesInfo{Msg: "查询礼品码失败", ER: err})
	}
	c.JSON(http.StatusOK, info)
}

func VerifyGiftCodeCtrl(c *gin.Context) {
	//获取参数
	code := c.Query("code")
	user := c.Query("user")
	info, err := handler.VerifyFiftCodeHandler(code, user)
	if err != nil {
		c.JSON(http.StatusOK, StructInfo.MesInfo{Msg: "礼品码验证失败", ER: err})
	}
	c.JSON(http.StatusOK, gin.H{
		"GoldCoins": info.ContentList.GoldCoins,
		"Diamonds":  info.ContentList.Diamonds,
		"Props":     info.ContentList.Props,
		"Heroes":    info.ContentList.Heroes,
		"Creeps":    info.ContentList.Creeps,
	})
}
