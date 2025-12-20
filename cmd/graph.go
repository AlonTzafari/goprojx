package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var graphCmd = &cobra.Command{
	Use: "graph",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("graph")
	},
}

func init() {
	rootCmd.AddCommand(graphCmd)
}
