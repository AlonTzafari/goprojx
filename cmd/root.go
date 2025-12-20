package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goprojx",
	Short: "A go monorepo cli tool",
}

func Execute() error {
	return rootCmd.Execute()
}
