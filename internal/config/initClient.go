package config

// 声明一个全局的rdb变量
import (
	"github.com/go-redis/redis"
)

var Rdb *redis.Client

// 初始化连接

func InitClient() (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err = Rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
