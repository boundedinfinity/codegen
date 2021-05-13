package loader

import (
	"errors"

	mapset "github.com/deckarep/golang-set"
)

type Node struct {
	Name string   `json:"name,omitempty" yaml:"name,omitempty"`
	Deps []string `json:"deps,omitempty" yaml:"deps,omitempty"`
}

func (t *Node) Add(deps ...*Node) {
	for _, dep := range deps {
		var found bool

		for _, n := range t.Deps {
			if dep.Name == n {
				found = true
				break
			}
		}

		if !found {
			t.Deps = append(t.Deps, dep.Name)
		}
	}
}

func NewNode(name string) *Node {
	n := &Node{
		Name: name,
		Deps: make([]string, 0),
	}

	return n
}

type Graph []*Node

func resolveGraph(graph Graph) (Graph, error) {
	nodeNames := make(map[string]*Node)
	nodeDependencies := make(map[string]mapset.Set)

	for _, node := range graph {
		nodeNames[node.Name] = node

		dependencySet := mapset.NewSet()
		for _, dep := range node.Deps {
			dependencySet.Add(dep)
		}
		nodeDependencies[node.Name] = dependencySet
	}

	var resolved Graph
	for len(nodeDependencies) != 0 {
		readySet := mapset.NewSet()

		for name, deps := range nodeDependencies {
			if deps.Cardinality() == 0 {
				readySet.Add(name)
			}
		}

		if readySet.Cardinality() == 0 {
			var g Graph
			for name := range nodeDependencies {
				g = append(g, nodeNames[name])
			}

			return g, errors.New("Circular dependency found")
		}

		for name := range readySet.Iter() {
			delete(nodeDependencies, name.(string))
			resolved = append(resolved, nodeNames[name.(string)])
		}

		for name, deps := range nodeDependencies {
			diff := deps.Difference(readySet)
			nodeDependencies[name] = diff
		}
	}

	return resolved, nil
}

func (t *Loader) displayGraph(graph Graph) {
	for _, node := range graph {
		for _, dep := range node.Deps {
			t.reportf([]string{}, "%s -> %s\n", node.Name, dep)
		}
	}
}
