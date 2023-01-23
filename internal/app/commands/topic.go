package commands

import (
	"github.com/owenrumney/pubsub/internal/app/pubsub"
	"github.com/spf13/cobra"
)

var topicsCmd = &cobra.Command{
	Use:   "topics",
	Short: "List create and delete topics",
}

var listTopicsCmd = &cobra.Command{
	Use:   "list",
	Short: "List topics",
	RunE: func(cmd *cobra.Command, args []string) error {
		psClient, err := pubsub.New(projectID)
		if err != nil {
			return err
		}

		psClient.ListTopics()
		return nil
	},
}

var createTopicCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a topic",
	RunE: func(cmd *cobra.Command, args []string) error {
		psClient, err := pubsub.New(projectID)
		if err != nil {
			return err
		}

		return psClient.CreateTopic(topicName)
	},
}

var deleteTopicCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a topic",
	RunE: func(cmd *cobra.Command, args []string) error {
		psClient, err := pubsub.New(projectID)
		if err != nil {
			return err
		}

		return psClient.DeleteTopic(topicName)
	},
}

func topicsCommand() *cobra.Command {
	topicsCmd.AddCommand(listTopicsCmd)
	topicsCmd.AddCommand(createTopicCmd)
	topicsCmd.AddCommand(deleteTopicCmd)

	return topicsCmd
}
