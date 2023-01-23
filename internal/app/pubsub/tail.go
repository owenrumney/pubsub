package pubsub

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/owenrumney/pubsub/pkg/logger"
)

func (t *pubsubClient) TailTopic(topicName, tailDuration string) error {

	subscriptionName := topicName + "-tail"
	t.CreateSubscription(subscriptionName, topicName, "")

	subscription := t.client.Subscription(subscriptionName)
	_, err := subscription.Exists(t.ctx)
	if err != nil {
		return err
	}

	duration, err := time.ParseDuration(tailDuration)
	if err != nil {
		logger.Error("Error parsing duration: %v", err)
		duration = 10 * time.Minute
	}

	logger.Info("Tailing topic: %s for %s", topicName, tailDuration)
	ctx, cancel := context.WithTimeout(t.ctx, duration)
	defer cancel()
	return subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		msg.Ack()
		fmt.Println(msg)
	})

}
