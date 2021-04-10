package main

import (
	"github.com/slack-go/slack"
	"log"
	"time"
)

func main() {

	// slack client
	api := slack.New(config.BotToken)

	// loop
	for {
		if message, exists := GetSlackNotifyMessage(config.SlackMessagesKey); exists {
			// build slack message
			msgOption := slack.MsgOptionCompose(
				slack.MsgOptionTS(message.ThreadTimestamp),
				slack.MsgOptionUsername(message.UserName),
				slack.MsgOptionIconURL(message.IconURL),
				slack.MsgOptionText(message.Text, true),
				slack.MsgOptionAsUser(false),
			)

			// send to slack
			_, _, err := api.PostMessage(
				message.ChannelID,
				msgOption,
			)
			if err != nil {
				log.Println(err)
			}
		}

		// for slack rate limits
		// https://api.slack.com/docs/rate-limits
		time.Sleep(time.Second * 2)
	}
}
