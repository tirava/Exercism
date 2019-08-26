// Package tree implements the tree building logic for an unsorted set of records.
package tree

import (
	"errors"
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
	duplicates := make(map[int]int, rlen)

	for i, record := range records {

		if record.ID >= rlen {
			return nil, errors.New("non-continuous")
		}
		if i|record.Parent == 0 {
			continue
		}
		if i == 0 && record.Parent != 0 {
			return nil, errors.New("root node has parent")
		}
		if record.ID <= record.Parent {
			return nil, errors.New("higher id parent of lower id")
		}
		if _, ok := duplicates[record.ID]; ok {
			return nil, errors.New("duplicate found")
		}
		duplicates[record.ID]++

		nodes[i] = Node{ID: record.ID}
		nodes[record.Parent].Children = append(nodes[record.Parent].Children, &nodes[i])
	}

	return &nodes[0], nil
}
