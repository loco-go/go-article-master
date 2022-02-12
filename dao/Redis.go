package dao

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	Rdb *redis.Client
)

// 初始化连接
func InitClient() (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "123.56.192.106:6379",
		Password: "$Sd59656422^&",
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = Rdb.Ping(ctx).Result()
	return err
}
