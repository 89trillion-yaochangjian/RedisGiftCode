package main

import (
	"RedisGiftCode/internal/router"
	"RedisGiftCode/internal/utils"
	"fmt"
)

func main() {
	utils.InitClient()
	rdb := utils.Rdb
	fmt.Println(rdb)
	router.GiftCodeRouter()
}
