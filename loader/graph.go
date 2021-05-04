package loader

import (
	"errors"

	mapset "github.com/deckarep/golang-set"
)

type Node struct {
	name string
	deps []string
}

func (t *Node) Add(dep ...string) {
	for _, x := range dep {
		var found bool

		for _, n := range t.deps {
			if x == n {
				found = true
				break
			}
		}

		if !found {
			t.deps = append(t.deps, x)
		}
	}
}

func NewNode(name string, deps ...string) *Node {
	n := &Node{
		name: name,
		deps: make([]string, 0),
	}

	n.Add(deps...)
	return n
}

type Graph []*Node

func resolveGraph(graph Graph) (Graph, error) {
	nodeNames := make(map[string]*Node)
	nodeDependencies := make(map[string]mapset.Set)

	for _, node := range graph {
		nodeNames[node.name] = node

		dependencySet := mapset.NewSet()
		for _, dep := range node.deps {
			dependencySet.Add(dep)
		}
		nodeDependencies[node.name] = dependencySet
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
		for _, dep := range node.deps {
			t.reportf(t.reportStack.S(), "%s -> %s\n", node.name, dep)
		}
	}
}
