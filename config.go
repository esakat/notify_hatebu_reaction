package main

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	RedisAddr        string `default:"localhost:6379" split_words:"true"`
	RedisDB          int    `default:"1" split_words:"true"`
	SlackMessagesKey string `default:"notify-queue" split_words:"true"`
	BotToken         string `required:"true" split_words:"true"`
}

var config Config

func init() {
	err := envconfig.Process("app", &config)
	if err != nil {
		log.Fatal(err.Error())
	}
	createRedisClient()
}
