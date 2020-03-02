// Package pov implements reparenting a graph on a selected node.
package pov

import "fmt"

// Graph base type.
type Graph struct {
	tree map[string][]string
}

// New returns new graph.
func New() *Graph {
	return &Graph{
		tree: make(map[string][]string),
	}
}

// AddNode adds new node.
func (*Graph) AddNode(nodeLabel string) {}

// AddArc constructs rest of the tree.
func (g *Graph) AddArc(from, to string) {
	g.tree[from] = append(g.tree[from], to)
}

// ArcList dumps tree.
func (g *Graph) ArcList() []string {
	var edges []string

	for from, v := range g.tree {
		for _, to := range v {
			edges = append(edges, fmt.Sprintf("%s -> %s", from, to))
		}
	}

	return edges
}

// ChangeRoot recreates tree to new root.
func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	path := g.getPath(oldRoot, newRoot)

	for i := 0; i < len(path)-1; i++ {
		oldTo, oldFrom := path[i], path[i+1]
		g.removeArc(oldFrom, oldTo)
		g.AddArc(oldTo, oldFrom)
	}

	return g
}

func (g *Graph) getPath(from, to string) []string {
	if from == to {
		return []string{to}
	}

	for _, child := range g.tree[from] {
		if path := g.getPath(child, to); path != nil {
			return append(path, from)
		}
	}

	return nil
}

func (g *Graph) removeArc(from, to string) {
	for i, child := range g.tree[from] {
		if child == to {
			g.tree[from] = append(g.tree[from][0:i], g.tree[from][i+1:]...)
			return
		}
	}
}
