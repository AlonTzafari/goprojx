package cmd

import (
	"os"

	"github.com/alontzafari/goprojx/internal/loader"
	"github.com/dominikbraun/graph/draw"
	"github.com/spf13/cobra"
)

var graphCmd = &cobra.Command{
	Use: "graph",
	RunE: func(cmd *cobra.Command, args []string) error {
		pkgs, err := loader.Load()
		if err != nil {
			return err
		}

		f, err := os.Create("./graph.gv")
		if err != nil {
			return err
		}

		return draw.DOT(pkgs, f)
	},
}

func init() {
	rootCmd.AddCommand(graphCmd)
}
