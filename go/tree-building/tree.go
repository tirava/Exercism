// Package tree implements the tree building logic for an unsorted set of records.
package tree

import (
	"fmt"
	"sort"
)

// Node is the base struct for nodes.
type Node struct {
	ID       int
	Children []*Node
}

// Record is the base struct for input records.
type Record struct {
	ID, Parent int
}

// Build builds the tree.
func Build(records []Record) (*Node, error) {

	rlen := len(records)

	if rlen == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	nodes := make([]Node, rlen)

	for i, record := range records {

		if record.ID >= rlen {
			return nil, fmt.Errorf("non-continuous (%d)", record.ID)
		}
		if i == 0 && record.Parent != 0 {
			return nil, fmt.Errorf("root node should not have a parent (%d)", record.Parent)
		}
		if i != 0 && record.Parent >= i {
			return nil, fmt.Errorf("parent id (%d) should be lower than its own id (%d)", record.Parent, i)
		}
		if i != record.ID {
			return nil, fmt.Errorf("duplicate found (%d)", record.ID)
		}

		nodes[i] = Node{ID: record.ID}
		if i != 0 {
			nodes[record.Parent].Children = append(nodes[record.Parent].Children, &nodes[i])
		}
	}

	return &nodes[0], nil
}
