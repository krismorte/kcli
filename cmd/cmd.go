package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "k-cli",
	Short: "k-cli is a CLI helper developers.",
	Long:  "k-cli is a CLI helper developers",
}

// Execute start all command chain.
func Execute() {
	rootCmd.AddCommand(RepoCommand())

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
