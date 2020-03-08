package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(diffCmd)
}

var diffCmd = &cobra.Command{
	Use: "diff",
	Short: "diff",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Diff")


	},
}

