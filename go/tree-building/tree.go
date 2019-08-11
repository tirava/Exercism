// Package tree implements the tree building logic for an unsorted set of records.
package tree

type Node struct {
	ID       int
	Children []*Node
}

type Record struct {
	ID     int
	Parent int
}

// Build builds tree.
func Build(records []Record) (*Node, error) {

	return nil, nil
}
