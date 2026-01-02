package cmd

import (
	"github.com/alontzafari/goprojx/internal/list"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	RunE: func(cmd *cobra.Command, args []string) error {
		head, err := cmd.Flags().GetString("head")
		if err != nil {
			return err
		}
		base, err := cmd.Flags().GetString("base")
		if err != nil {
			return err
		}

		return list.Affected(cmd.Context(), head, base)
	},
}

func init() {
	listCmd.Flags().String("head", "HEAD", "head commit hash")
	listCmd.Flags().String("base", "main", "base commit hash")
	rootCmd.AddCommand(listCmd)
}
