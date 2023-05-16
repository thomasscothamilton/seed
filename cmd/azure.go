package cmd

import (
	"github.com/spf13/cobra"
)

var (
	azureCmd = &cobra.Command{
		Use: "azure",
		Run: func(cmd *cobra.Command, args []string) {
			// code you wish you had...
		},
	}
)

func init() {
	rootCmd.AddCommand(azureCmd)
}
