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
		if i == 0 && record.Parent != 0 {
			return nil, errors.New("root node has parent")
		}
		if _, ok := duplicates[record.ID]; ok {
			return nil, errors.New("duplicate found")
		}
		duplicates[record.ID]++

		nodes[i] = Node{ID: record.ID}

		if i != 0 {
			if record.ID <= record.Parent {
				return nil, errors.New("higher id parent of lower id")
			}
			nodes[record.Parent].Children = append(nodes[record.Parent].Children, &nodes[i])
		}
	}

	return &nodes[0], nil
}

// Build builds tree.
//func Build(records []Record) (*Node, error) {
//
//	if len(records) == 0 {
//		return nil, nil
//	}
//
//	err := checkErrors(records)
//	if err != nil {
//		return nil, err
//	}
//
//	node := &Node{}
//	calcNodes(records, node)
//
//	return node, err
//}
//
//func checkErrors(records []Record) error {
//
//	duplicates := make(map[int]int, len(records))
//
//	for _, record := range records {
//
//		if record.ID == 0 && record.Parent != 0 {
//			return errors.New("root node has parent")
//		}
//		if record.Parent >= record.ID && record.ID != 0 {
//			return errors.New("higher id parent of lower id")
//		}
//		if _, ok := duplicates[record.ID]; ok {
//			return errors.New("duplicate found")
//		}
//
//		duplicates[record.ID]++
//	}
//
//	for i := 0; i < len(records); i++ {
//		if _, ok := duplicates[i]; !ok {
//			return errors.New("non-continuous")
//		}
//	}
//
//	return nil
//}
//
//func calcNodes(records []Record, node *Node) {
//
//	for _, record := range records {
//		if node.ID == record.Parent && record.ID != record.Parent {
//			newNode := &Node{ID: record.ID}
//			node.Children = append(node.Children, newNode)
//			calcNodes(records, newNode)
//		}
//	}
//
//	sort.Slice(node.Children, func(i, j int) bool {
//		return node.Children[i].ID < node.Children[j].ID
//	})
//}
