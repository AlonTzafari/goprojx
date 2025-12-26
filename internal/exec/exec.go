package exec

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/alontzafari/goprojx/internal/gitlib"
	"github.com/alontzafari/goprojx/internal/hashset"
	"github.com/alontzafari/goprojx/internal/loader"
	"github.com/dominikbraun/graph"
	"golang.org/x/tools/go/packages"
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
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	hs := hashset.New[string]()
	for i := 0; i < len(files); i++ {
		files[i] = filepath.Dir(filepath.Join(cwd, files[i]))
		hs.Add(files[i])
	}
	am, err := pkgs.AdjacencyMap()
	if err != nil {
		return err
	}

	changed := make([]string, 0)
	for k := range am {
		pkg, err := pkgs.Vertex(k)
		if err != nil {
			return err
		}

		pkgDir := filepath.Clean(pkg.Dir)
		if hs.Has(pkgDir) {
			changed = append(changed, k)
			continue
		}
	}

	fmt.Printf("detected changes in the following packages:\n%s\n\n", strings.Join(changed, "\n"))

	affectedIds := hashset.New(changed...)
	for _, s := range changed {
		graph.DFS(pkgs, s, func(s string) bool {
			affectedIds.Add(s)
			return false
		})
	}

	affectedPkgs := make([]*packages.Package, 0, affectedIds.Len())
	for pkgId := range affectedIds.Iter() {
		pkg, err := pkgs.Vertex(pkgId)
		if err != nil {
			return err
		}
		affectedPkgs = append(affectedPkgs, pkg)
	}

	affectedPkgsDisplay := make([]string, 0, len(affectedPkgs))
	for _, pkg := range affectedPkgs {
		affectedPkgsDisplay = append(affectedPkgsDisplay, fmt.Sprintf("%s - %s", pkg.ID, pkg.Name))
	}
	fmt.Printf("affected packages:\n%s\n\n", strings.Join(affectedPkgsDisplay, "\n"))

	for _, pkg := range affectedPkgs {
		fmt.Printf("running %s for package %s\n", command, pkg.ID)
		commandParts := strings.Split(command, " ")
		bin := commandParts[0]
		args := commandParts[1:]
		args = append(args, pkg.Dir)
		cmd := exec.CommandContext(ctx, bin, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return err
		}
		fmt.Printf("\n")
	}

	return nil
}
