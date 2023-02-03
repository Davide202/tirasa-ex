package function

import (
	"context"
	"encoding/json"

	"fmt"
	"log"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
)

type User struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	//Info    string `json:"info"`
}

func (u User) String() string {
	return fmt.Sprintf("\nName: %v, \nSurname: %v", u.Name, u.Surname)
}

func init() {
	functions.CloudEvent("HelloPubSub", helloPubSub)
}

// helloPubSub consumes a CloudEvent message and extracts the Pub/Sub message.
func helloPubSub(ctx context.Context, e event.Event) error {
	var err error
	var user User
	log.Printf("Received Message from Pub/Sub")

	//VERSION 1

	log.Printf("Received: %v", e.String())
	log.Printf("ID: %v", e.ID())
	log.Printf("Extensions: %v", e.Extensions())
	log.Printf("Extensions: %v", e.Data())

	if err = e.DataAs(&user); err != nil {
		return fmt.Errorf("event.DataAs: %v", err)
	}
	log.Printf("User: %v", user.String())

	err = json.Unmarshal(e.Data(), &user)
	if err != nil {
		return fmt.Errorf("event.DataAs: %v", err)
	}
	log.Printf("User: %v", user.String())
	/*
		// VERSION 2
		//topicID := os.Getenv("TOPIC")
		projectID := os.Getenv("PROJECT_ID")
		subscriptionID := os.Getenv("SUBSCRIPTION")
		log.Printf("Project: %v - Subscription: %v \n", projectID, subscriptionID)

		//ctx = context.Background()
		client, err := pubsub.NewClient(ctx, projectID)
		if err != nil {
			log.Printf("Error creating client: %v\n", err)
		}
		sub := client.Subscription(subscriptionID)
		if err != nil {
			log.Printf("Error creating subscription: %v\n", err)
		}
		sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
			log.Printf("Got message: %s", m.Data)
			err = json.Unmarshal(m.Data, &user)
			if err != nil {
				log.Printf("Error Unmashal User %v", err.Error())
			}
			log.Printf("User: %v", user.String())
			m.Ack()
		})
	*/
	return nil
}
