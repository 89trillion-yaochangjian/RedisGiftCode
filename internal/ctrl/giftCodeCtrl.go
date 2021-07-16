package ctrl

import (
	"RedisGiftCode/internal/model"
	"RedisGiftCode/internal/service"
	"RedisGiftCode/internal/status"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

//管理后台调用 - 创建礼品码

func CreateGiftCode(c *gin.Context) {
	//获取参数
	info, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, status.ParamErr)
		return
	}
	var giftCodeInfo model.GiftCodeInfo
	json.Unmarshal(info, &giftCodeInfo)
	// 0--不限定用户，限定领取次数   -1--指定用户一次领取  -2--不限定用户，不限定次数
	if giftCodeInfo.CodeType != -1 && giftCodeInfo.CodeType != 0 && giftCodeInfo.CodeType != -2 {
		c.JSON(http.StatusBadRequest, status.CodeTypeErr)
		return
	}
	//指定用户一次领取参数判断
	if giftCodeInfo.CodeType == -1 && len(giftCodeInfo.User) == 0 {
		c.JSON(http.StatusBadRequest, status.CodeUserErr)
		return
	}

	code, err1 := service.CreateGiftCodeService(giftCodeInfo)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, err1)
		return
	}
	c.JSON(http.StatusOK, status.OK.WithData(code))
}

//管理后台调用 - 查询礼品码信息

func GetGiftCodeInfoCtrl(c *gin.Context) {
	//获取参数
	code := c.Query("code")
	if len(code) != 8 {
		c.JSON(http.StatusBadRequest, status.CodeLenErr)
		return
	}
	info, err := service.GetGiftCodeInfoService(code)
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}
	c.JSON(http.StatusOK, status.OK.WithData(info))
}

//客户端调用 - 验证礼品码

func VerifyGiftCodeCtrl(c *gin.Context) {
	//获取参数
	code := c.Query("code")
	if len(code) != 8 {
		c.JSON(http.StatusBadRequest, status.CodeLenErr)
		return
	}
	user := c.Query("user")
	if len(user) == 0 {
		c.JSON(http.StatusBadRequest, status.CodeUserErr)
		return
	}
	info, err := service.VerifyFiftCodeService(code, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	res := model.GiftContentList{
		GoldCoins: info.ContentList.GoldCoins,
		Diamonds:  info.ContentList.Diamonds,
		Props:     info.ContentList.Props,
		Heroes:    info.ContentList.Heroes,
		Creeps:    info.ContentList.Creeps,
	}
	c.JSON(http.StatusOK, status.OK.WithData(res))
}
