package commands

import "os"

var (
	projectID        string = os.Getenv("PROJECT_ID")
	debug            bool
	topicName        string
	subscriptionName string
	pushEndpoint     string
	tailDuration     string
)
