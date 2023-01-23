package commands

import (
	"os"

	"github.com/owenrumney/pubsub/pkg/logger"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "pubsub",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logger.Configure(os.Stdout, debug)
	},
}

func RootCommand() *cobra.Command {
	rootCmd.PersistentFlags().StringVarP(&projectID, "project", "p", projectID, "GCP project ID")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", debug, "Enable debug logging")
	rootCmd.PersistentFlags().StringVarP(&topicName, "topic-name", "t", topicName, "Topic name")

	rootCmd.AddCommand(topicsCommand(), subscriptionCommand())

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	return rootCmd
}
