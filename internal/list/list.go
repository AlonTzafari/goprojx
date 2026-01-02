package list

import (
	"context"
	"os"

	"github.com/alontzafari/goprojx/internal/analyzer"
	"github.com/alontzafari/goprojx/internal/gitlib"
	"github.com/alontzafari/goprojx/internal/loader"
)

func Affected(ctx context.Context, head, base string) error {
	pkgs, err := loader.Load()
	if err != nil {
		return err
	}

	files, err := gitlib.GetDiffs(ctx, head, base)
	if err != nil {
		return err
	}

	changed, err := analyzer.GetPkgIdsFromFiles(pkgs, files)
	if err != nil {
		return err
	}

	affectedPkgs, err := analyzer.GetDependents(pkgs, changed)
	if err != nil {
		return err
	}

	for i, pkg := range affectedPkgs {
		line := pkg
		if i < len(affectedPkgs)-1 {
			line += "\n"
		}
		_, err := os.Stdout.WriteString(line)
		if err != nil {
			return err
		}
	}

	return nil
}
