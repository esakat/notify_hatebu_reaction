package main

import (
	"cloud.google.com/go/firestore"
	"context"
	. "github.com/esakat/observe_my_hatebu/data"
	"google.golang.org/api/iterator"
	"log"
)

var ctx = context.Background()
var firestoreClient *firestore.Client

func createFirestoreClient() {
	var err error
	firestoreClient, err = firestore.NewClient(ctx, config.ProjectID)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
}

func GetSlackNotifyMessage() (slackMessages []SlackNotifyMessage) {

	iter := firestoreClient.Collection(config.SlackMessageCollectionName).Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("failed to iterate: %v", err)
		}

		var sm SlackMessage
		doc.DataTo(&sm)
		slackMessages = append(slackMessages, SlackNotifyMessage{
			ID:           doc.Ref.ID,
			SlackMessage: sm,
		})
	}

	return
}

func DeleteSlackNotifyMessage(id string) {
	_, err := firestoreClient.Collection(config.SlackMessageCollectionName).Doc(id).Delete(ctx)
	if err != nil {
		panic(err)
	}
}
