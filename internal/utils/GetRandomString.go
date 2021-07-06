package utils

import (
	"math/rand"
	"time"
)

// 随机生成指定位数的大写字母和数字的组合

func  GetGiftCodeUtil() string {
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 8; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

