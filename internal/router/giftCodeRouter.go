package router

import (
	"RedisGiftCode/internal/ctrl"
	"github.com/gin-gonic/gin"
)

func GiftCodeRouter() {
	r := gin.Default()
	r.POST("/CreateGiftCode", ctrl.CreateGiftCode)
	r.GET("/GetGiftCodeInfo", ctrl.GetGiftCodeInfoCtrl)
	r.GET("/VerifyGiftCode", ctrl.VerifyGiftCodeCtrl)
	r.Run(":8080")
}
