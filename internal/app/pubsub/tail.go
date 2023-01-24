package pubsub

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/owenrumney/pubsub/pkg/logger"
)

func (t *pubsubClient) TailTopic(topicName, tailDuration string) error {

	subscriptionName := topicName + "-tail"
	logger.Info("Creating tail subscription: %s", subscriptionName)
	subscription, err := t.CreateSubscription(subscriptionName, topicName, "")
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(t.ctx, 10*time.Second)
	defer cancel()
	_, err = subscription.Exists(ctx)
	if err != nil {
		return err
	}

	duration, err := time.ParseDuration(tailDuration)
	if err != nil {
		logger.Error("Error parsing duration: %v", err)
		duration = 10 * time.Minute
	}

	logger.Info("Tailing topic: %s for %s", topicName, tailDuration)
	ctx, cancel = context.WithTimeout(t.ctx, duration)
	defer cancel()
	return subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		msg.Ack()
		eventHeader := fmt.Sprintf("Published: %s ID: %s", msg.PublishTime, msg.ID)
		fmt.Printf("%s\n%s\n", eventHeader, strings.Repeat("-", len(eventHeader)))
		for attr, val := range msg.Attributes {
			fmt.Printf("%s:\t\t%v\n", attr, val)
		}

		var content map[string]interface{}
		err := json.Unmarshal(msg.Data, &content)
		if err != nil {
			fmt.Printf("%s\n", string(msg.Data))
		}
		body, err := json.MarshalIndent(content, "", "  ")
		if err != nil {
			logger.Error("Error marshalling message: %v", err)
			fmt.Printf("%s\n", string(msg.Data))
		}
		fmt.Printf("%v\n", string(body))
	},
	)
}
