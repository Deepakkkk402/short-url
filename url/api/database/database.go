package database

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
)

var ctx = context.Background()

func CreateClient(dbNo int) *redis.Client {
	//	rdb :=  redis
	rdb _ := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASS"),
		DB:       dbNo,
	})
	return rdb

}
