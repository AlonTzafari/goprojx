package cmd

import (
	"github.com/alontzafari/goprojx/internal/exec"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use: "exec",

	RunE: func(cmd *cobra.Command, args []string) error {
		head, err := cmd.Flags().GetString("head")
		if err != nil {
			return err
		}
		base, err := cmd.Flags().GetString("base")
		if err != nil {
			return err
		}

		return exec.OnAffected(cmd.Context(), "go test", head, base)
	},
}

func init() {
	execCmd.Flags().String("head", "HEAD", "head commit hash")
	execCmd.Flags().String("base", "main", "base commit hash")
	rootCmd.AddCommand(execCmd)
}
