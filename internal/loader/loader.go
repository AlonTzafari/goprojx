package loader

import (
	"errors"

	"github.com/alontzafari/goprojx/internal/hashset"
	"github.com/dominikbraun/graph"
	"golang.org/x/tools/go/packages"
)

func Load() (graph.Graph[string, *packages.Package], error) {
	cfg := &packages.Config{
		Mode: packages.LoadImports,
	}

	pkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		return nil, err
	}
	g := graph.New(hash, graph.Directed(), graph.Acyclic(), graph.PreventCycles())

	hs := hashset.New[string]()

	for _, pkg := range pkgs {
		g.AddVertex(pkg)
		hs.Add(hash(pkg))
	}

	for _, pkg := range pkgs {
		for k := range pkg.Imports {
			if !hs.Has(k) {
				continue
			}

			err := g.AddEdge(k, hash(pkg))
			if err == graph.ErrEdgeCreatesCycle {
				return nil, errors.New("Cyclic imports detected")
			}
		}
	}
	return g, nil

}

func hash(pkg *packages.Package) string {
	return pkg.ID
}
