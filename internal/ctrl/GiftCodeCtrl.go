package ctrl

import (
	"RedisGiftCode/StructInfo"
	"RedisGiftCode/internal/handler"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func CreateGiftCode(c *gin.Context) {
	//获取参数
	info,err1 := c.GetRawData()
	if err1 != nil{
		fmt.Printf("")
	}
	var giftCodeInfo StructInfo.GiftCodeInfo
	json.Unmarshal(info,&giftCodeInfo)
	//var giftCodeInfo = StructInfo.GiftCodeInfo{}
	//c.ShouldBind(&giftCodeInfo)
	//调用Handler
	code := handler.CreateGiftCodeHandler(giftCodeInfo)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
	})
}

func GetGiftCodeInfoCtrl(c *gin.Context) {
	//获取参数
	code := c.Query("code")
	info := handler.GetFiftCodeInfoHandler(code)
	c.JSON(http.StatusOK, info)
}

func VerifyGiftCodeCtrl(c *gin.Context) {
	//获取参数
	code := c.Query("code")
	user := c.Query("user")
	info := handler.VerifyFiftCodeHandler(code,user)
	c.JSON(http.StatusOK, gin.H{
		"GoldCoins": info.ContentList.GoldCoins,
		"Diamonds": info.ContentList.Diamonds,
		"Props": info.ContentList.Props,
		"Heroes": info.ContentList.Heroes,
		"Creeps": info.ContentList.Creeps,
	})
}