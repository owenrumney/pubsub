package commands

import (
	"github.com/owenrumney/pubsub/internal/app/pubsub"
	"github.com/spf13/cobra"
)

var subscriptionCmd = &cobra.Command{
	Use:   "subscriptions",
	Short: "List create and delete subscriptions",
}

var listSubscriptionCmd = &cobra.Command{
	Use:   "list",
	Short: "List subscriptions",
	RunE: func(cmd *cobra.Command, args []string) error {
		psClient, err := pubsub.New(projectID)
		if err != nil {
			return err
		}

		psClient.ListSubscriptions()
		return nil
	},
}

var createSubscriptionCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a subscription",
	RunE: func(cmd *cobra.Command, args []string) error {
		psClient, err := pubsub.New(projectID)
		if err != nil {
			return err
		}

		_, err = psClient.CreateSubscription(subscriptionName, topicName, pushEndpoint)
		return err
	},
}

var deleteSubscriptionCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a subscription",
	RunE: func(cmd *cobra.Command, args []string) error {
		psClient, err := pubsub.New(projectID)
		if err != nil {
			return err
		}

		return psClient.DeleteSubscription(subscriptionName)
	},
}

func subscriptionCommand() *cobra.Command {
	subscriptionCmd.AddCommand(listSubscriptionCmd)
	subscriptionCmd.AddCommand(createSubscriptionCmd)
	subscriptionCmd.AddCommand(deleteSubscriptionCmd)

	createSubscriptionCmd.Flags().StringVarP(&subscriptionName, "sub-name", "s", "", "Subscription name")
	createSubscriptionCmd.Flags().StringVarP(&pushEndpoint, "push-endpoint", "e", "", "The push endpoint to send messages to")
	createSubscriptionCmd.MarkFlagRequired("topic")

	deleteSubscriptionCmd.Flags().StringVarP(&subscriptionName, "sub-name", "s", "", "Subscription name")

	return subscriptionCmd
}
