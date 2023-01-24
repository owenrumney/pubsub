package pubsub

import (
	"fmt"

	"cloud.google.com/go/pubsub"
	"github.com/owenrumney/pubsub/pkg/logger"
	"google.golang.org/api/iterator"
)

func (p *pubsubClient) ListSubscriptions() error {

	subscriptions := p.client.Subscriptions(p.ctx)
	for {
		subscription, err := subscriptions.Next()
		if err == iterator.Done {
			return nil
		}

		if err != nil {
			logger.Error("Error listing subscriptions: %v", err)
		}

		fmt.Println(subscription)
	}
}

func (t *pubsubClient) CreateSubscription(subscriptionName, topicName, pushEndpoint string) (*pubsub.Subscription, error) {
	if topicName == "" {
		return nil, fmt.Errorf("topic name is required")
	}
	if subscriptionName == "" {
		return nil, fmt.Errorf("subscription name is required")
	}

	logger.Debug("Creating subscription: %s", subscriptionName)
	subscription := t.client.Subscription(subscriptionName)
	exists, err := subscription.Exists(t.ctx)
	if err != nil {
		logger.Error("Error checking if subscription exists: %v", err)
	}
	if exists {
		logger.Warn("Topic already exists: %s", topicName)
		return subscription, nil
	}

	subscriptionConfig := pubsub.SubscriptionConfig{
		Topic: t.client.Topic(topicName),
	}

	if pushEndpoint != "" {
		subscriptionConfig.PushConfig = pubsub.PushConfig{
			Endpoint: pushEndpoint,
		}
	}

	subscription, err = t.client.CreateSubscription(t.ctx, subscriptionName, subscriptionConfig)
	if err != nil {
		return nil, err
	}
	logger.Info("Subscription created: %s", subscriptionName)
	return subscription, nil
}

func (t *pubsubClient) DeleteSubscription(subscriptionName string) error {
	logger.Debug("Deleting subscription: %s", subscriptionName)
	subscription := t.client.Subscription(subscriptionName)
	exists, err := subscription.Exists(t.ctx)
	if err != nil {
		logger.Error("Error checking if subscription exists: %v", err)
	}
	if !exists {
		logger.Warn("Subscription does not exist: %s", subscriptionName)
		return nil
	}

	if err := subscription.Delete(t.ctx); err != nil {
		return err
	}
	logger.Info("Subscription deleted: %s", subscriptionName)
	return nil
}
