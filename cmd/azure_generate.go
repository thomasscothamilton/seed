package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	name        string
	generateCmd = &cobra.Command{
		Use: "generate",
		Run: func(cmd *cobra.Command, args []string) {
			out := cmd.OutOrStdout()
			// Print the final resolved value from binding cobra flags and viper config
			fmt.Fprintln(out, "your name is:", name)
		},
	}
)

func init() {
	generateCmd.Flags().StringVar(&name, "name", "n", name)
	azureCmd.AddCommand(generateCmd)
}
