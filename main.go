package main

import (
	"github.com/esakat/observe_my_hatebu/data"
	"github.com/slack-go/slack"
	"log"
	"time"
)

type SlackNotifyMessage struct {
	ID           string
	SlackMessage data.SlackMessage
}

func main() {

	// slack client
	api := slack.New(config.BotToken)

	slackMessages := GetSlackNotifyMessage()

	log.Println(len(slackMessages))

	for _, slackMessage := range slackMessages {
		// build slack message
		message := slackMessage.SlackMessage
		msgOption := slack.MsgOptionCompose(
			slack.MsgOptionTS(message.ThreadTimestamp),
			slack.MsgOptionUsername(message.UserName),
			slack.MsgOptionIconURL(message.IconURL),
			slack.MsgOptionText(message.Text, true),
			slack.MsgOptionAsUser(false),
		)

		log.Println(message.Text)

		// send to slack
		_, _, err := api.PostMessage(
			message.ChannelID,
			msgOption,
		)
		if err != nil {
			log.Println(err)
		} else {
			// if succeeded, delete from firestore
			DeleteSlackNotifyMessage(slackMessage.ID)
		}

		// for slack rate limits
		// https://api.slack.com/docs/rate-limits
		time.Sleep(time.Second * 2)
	}

}
