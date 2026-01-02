package analyzer

import (
	"path/filepath"

	"github.com/alontzafari/goprojx/internal/hashset"
	"github.com/dominikbraun/graph"
	"golang.org/x/tools/go/packages"
)

type Project = graph.Graph[string, *packages.Package]

func GetPkgIdsFromFiles(proj Project, files []string) ([]string, error) {
	hs := hashset.New[string]()
	for _, file := range files {
		hs.Add(filepath.Dir(file))
	}

	am, err := proj.AdjacencyMap()
	if err != nil {
		return nil, err
	}

	pkgIds := make([]string, 0)
	for k := range am {
		pkg, err := proj.Vertex(k)
		if err != nil {
			return nil, err
		}

		pkgDir := filepath.Clean(pkg.Dir)
		if hs.Has(pkgDir) {
			pkgIds = append(pkgIds, k)
		}
	}

	return pkgIds, nil
}

func GetDependents(proj Project, ids []string) ([]string, error) {
	affectedIds := hashset.New(ids...)
	for _, id := range ids {
		err := graph.DFS(proj, id, func(s string) bool {
			affectedIds.Add(s)
			return false
		})
		if err != nil {
			return nil, err
		}
	}

	return affectedIds.ToSlice(), nil
}
