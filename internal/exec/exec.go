package exec

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/alontzafari/goprojx/internal/analyzer"
	"github.com/alontzafari/goprojx/internal/gitlib"
	"github.com/alontzafari/goprojx/internal/loader"
)

func OnAffected(ctx context.Context, command, head, base string) error {
	t0 := time.Now()
	defer func() {
		fmt.Printf("elapsed %s\n", time.Since(t0).String())
	}()

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

	for _, pkgId := range affectedPkgs {
		pkg, err := pkgs.Vertex(pkgId)
		if err != nil {
			return err
		}
		fmt.Printf("running %s for package %s\n", command, pkg.ID)
		commandParts := strings.Split(command, " ")
		bin := commandParts[0]
		args := commandParts[1:]
		args = append(args, pkg.Dir)
		cmd := exec.CommandContext(ctx, bin, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			return err
		}
		fmt.Printf("\n")
	}

	return nil
}
