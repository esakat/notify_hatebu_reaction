package main

import (
	"context"
	"encoding/json"
	. "github.com/esakat/observe_my_hatebu/data"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

func createRedisClient() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: "",
		DB:       config.RedisDB,
	})
}

func GetSlackNotifyMessage(redisKey string) (slackMessage SlackMessage, exists bool) {
	sm, err := rdb.RPop(ctx, redisKey).Result()
	if err == redis.Nil {
		exists = false
		return
	} else if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(sm), &slackMessage)
	if err != nil {
		panic(err)
	}
	exists = true

	return
}
