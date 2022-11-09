package db

import (
	"fmt"
	"github.com/go-redis/redis"
)

// 定义一个全局变量
var Redisdb *redis.Client

func init() {

	err := initRedis()
	if err != nil {
		fmt.Printf("connect redis failed! err : %v\n", err)
		return
	}
}

func initRedis() (err error) {
	Redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // 指定
		Password: "",
		DB:       0, // redis一共16个库，指定其中一个库即可
	})
	_, err = Redisdb.Ping().Result()
	return
}
