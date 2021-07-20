package main

import (
	"RedisGiftCode/internal/config"
	"RedisGiftCode/internal/router"
)

func main() {
	//// 初始化连接
	config.InitClient()
	//调用路由
	router.GiftCodeRouter()
}
