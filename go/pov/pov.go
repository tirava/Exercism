// Package pov implements reparenting a graph on a selected node.
package pov

// Graph base type.
type Graph struct{}

// New returns new graph.
func New() *Graph {
	return &Graph{}
}

// AddNode adds new node.
func (*Graph) AddNode(nodeLabel string) {

}

// AddArc constructs rest of the tree.
func (*Graph) AddArc(from, to string) {

}

// ArcList dumps tree.
func (*Graph) ArcList() []string {
	return nil
}

// ChangeRoot recreates tree to new root.
func (*Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	return &Graph{}
}
