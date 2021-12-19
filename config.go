package main

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	SlackMessageCollectionName string `required:"true" split_words:"true"`
	BotToken                   string `required:"true" split_words:"true"`
	ProjectID                  string `required:"true" split_words:"true"`
}

var config Config

func init() {
	err := envconfig.Process("app", &config)
	if err != nil {
		log.Fatal(err.Error())
	}
	createFirestoreClient()
}
