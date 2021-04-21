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
	// A map containing the node names and the actual node object
	nodeNames := make(map[string]*Node)

	// A map containing the nodes and their dependencies
	nodeDependencies := make(map[string]mapset.Set)

	// Populate the maps
	for _, node := range graph {
		nodeNames[node.name] = node

		dependencySet := mapset.NewSet()
		for _, dep := range node.deps {
			dependencySet.Add(dep)
		}
		nodeDependencies[node.name] = dependencySet
	}

	// Iteratively find and remove nodes from the graph which have no dependencies.
	// If at some point there are still nodes in the graph and we cannot find
	// nodes without dependencies, that means we have a circular dependency
	var resolved Graph
	for len(nodeDependencies) != 0 {
		// Get all nodes from the graph which have no dependencies
		readySet := mapset.NewSet()
		for name, deps := range nodeDependencies {
			if deps.Cardinality() == 0 {
				readySet.Add(name)
			}
		}

		// If there aren't any ready nodes, then we have a cicular dependency
		if readySet.Cardinality() == 0 {
			var g Graph
			for name := range nodeDependencies {
				g = append(g, nodeNames[name])
			}

			return g, errors.New("Circular dependency found")
		}

		// Remove the ready nodes and add them to the resolved graph
		for name := range readySet.Iter() {
			delete(nodeDependencies, name.(string))
			resolved = append(resolved, nodeNames[name.(string)])
		}

		// Also make sure to remove the ready nodes from the
		// remaining node dependencies as well
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
			t.report("%s -> %s\n", node.name, dep)
		}
	}
}
