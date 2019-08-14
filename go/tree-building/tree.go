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
	ID     int
	Parent int
}

// Build builds tree.
func Build(records []Record) (*Node, error) {

	if len(records) == 0 {
		return nil, nil
	}

	err := checkErrors(records)
	if err != nil {
		return nil, err
	}

	node := &Node{}
	calcNodes(records, node)

	return node, err
}

func checkErrors(records []Record) error {

	duplicates := make(map[int]int, len(records))

	for _, record := range records {

		if record.ID == 0 && record.Parent != 0 {
			return errors.New("root node has parent")
		}
		if record.Parent >= record.ID && record.ID != 0 {
			return errors.New("higher id parent of lower id")
		}
		if _, ok := duplicates[record.ID]; ok {
			return errors.New("duplicate found")
		}

		duplicates[record.ID]++
	}

	for i := 0; i < len(records); i++ {
		if _, ok := duplicates[i]; !ok {
			return errors.New("non-continuous")
		}
	}

	return nil
}

func calcNodes(records []Record, node *Node) {

	for _, record := range records {
		if node.ID == record.Parent && record.ID != record.Parent {
			newNode := &Node{ID: record.ID}
			node.Children = append(node.Children, newNode)
			calcNodes(records, newNode)
		}
	}

	sort.Slice(node.Children, func(i, j int) bool {
		return node.Children[i].ID < node.Children[j].ID
	})
}
