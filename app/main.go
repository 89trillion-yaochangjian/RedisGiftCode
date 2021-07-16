package main

import (
	"RedisGiftCode/internal/router"
	"RedisGiftCode/internal/utils"
)

func main() {
	//// 初始化连接
	utils.InitClient()
	//调用路由
	router.GiftCodeRouter()
}
