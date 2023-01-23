package pubsub

import (
	"fmt"

	"github.com/owenrumney/pubsub/pkg/logger"
	"google.golang.org/api/iterator"
)

func (t *pubsubClient) ListTopics() error {
	topics := t.client.Topics(t.ctx)
	for {
		topic, err := topics.Next()
		if err == iterator.Done {
			return nil
		}

		if err != nil {
			logger.Error("Error listing topics: %v", err)
		}

		fmt.Println(topic)
	}
}

func (t *pubsubClient) CreateTopic(topicName string) error {
	if topicName == "" {
		return fmt.Errorf("topic name is required")
	}

	logger.Debug("Creating topic: %s", topicName)
	topic := t.client.Topic(topicName)
	exists, err := topic.Exists(t.ctx)
	if err != nil {
		logger.Error("Error checking if topic exists: %v", err)
	}
	if exists {
		logger.Warn("Topic already exists: %s", topicName)
		return nil
	}

	_, err = t.client.CreateTopic(t.ctx, topicName)
	if err != nil {
		return nil
	}
	logger.Info("Topic created: %s", topicName)
	return nil
}

func (t *pubsubClient) DeleteTopic(topicName string) error {
	if topicName == "" {
		return fmt.Errorf("topic name is required")
	}

	logger.Debug("Deleting topic: %s", topicName)
	topic := t.client.Topic(topicName)
	exists, err := topic.Exists(t.ctx)
	if err != nil {
		logger.Error("Error checking if topic exists: %v", err)
	}
	if !exists {
		logger.Warn("Topic does not exist: %s", topicName)
		return nil
	}

	if err := topic.Delete(t.ctx); err != nil {
		return err
	}
	logger.Info("Topic deleted: %s", topicName)
	return nil
}
