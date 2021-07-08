package ctrl

import (
	"RedisGiftCode/internal/handler"
	"RedisGiftCode/internal/structInfo"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

//管理后台调用 - 创建礼品码

func CreateGiftCode(c *gin.Context) {
	//获取参数
	info, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, structInfo.ParamErr)
		return
	}
	var giftCodeInfo structInfo.GiftCodeInfo
	json.Unmarshal(info, &giftCodeInfo)
	// 0--不限定用户，限定领取次数   -1--指定用户一次领取  -2--不限定用户，不限定次数
	if giftCodeInfo.CodeType != -1 && giftCodeInfo.CodeType != 0 && giftCodeInfo.CodeType != -2 {
		c.JSON(http.StatusBadRequest, structInfo.CodeTypeErr)
		return
	}
	//指定用户一次领取参数判断
	if giftCodeInfo.CodeType == -1 && len(giftCodeInfo.User) == 0 {
		c.JSON(http.StatusBadRequest, structInfo.CodeUserErr)
		return
	}

	code, err := handler.CreateGiftCodeHandler(giftCodeInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structInfo.CreateErr)
		return
	}
	c.JSON(http.StatusOK, structInfo.OK.WithData(code))
}

//管理后台调用 - 查询礼品码信息

func GetGiftCodeInfoCtrl(c *gin.Context) {
	//获取参数
	code := c.Query("code")
	if len(code) != 8 {
		c.JSON(http.StatusBadRequest, structInfo.CodeLenErr)
		return
	}
	info, err := handler.GetFiftCodeInfoHandler(code)
	if err != nil {
		c.JSON(http.StatusOK, structInfo.FindCodeErr)
		return
	}
	c.JSON(http.StatusOK, structInfo.OK.WithData(info))
}

//客户端调用 - 验证礼品码

func VerifyGiftCodeCtrl(c *gin.Context) {
	//获取参数
	code := c.Query("code")
	if len(code) != 8 {
		c.JSON(http.StatusBadRequest, structInfo.CodeLenErr)
		return
	}
	user := c.Query("user")
	if len(user) == 0 {
		c.JSON(http.StatusBadRequest, structInfo.CodeUserErr)
		return
	}
	info, err := handler.VerifyFiftCodeHandler(code, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structInfo.VerifyCodeErr)
		return
	}
	res := structInfo.GiftContentList{
		GoldCoins: info.ContentList.GoldCoins,
		Diamonds:  info.ContentList.Diamonds,
		Props:     info.ContentList.Props,
		Heroes:    info.ContentList.Heroes,
		Creeps:    info.ContentList.Creeps,
	}
	c.JSON(http.StatusOK, structInfo.OK.WithData(res))
}
