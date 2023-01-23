package pubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
)

type pubsubClient struct {
	ctx    context.Context
	client *pubsub.Client
}

func New(projectID string) (*pubsubClient, error) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}

	return &pubsubClient{
		ctx:    ctx,
		client: client,
	}, nil
}
