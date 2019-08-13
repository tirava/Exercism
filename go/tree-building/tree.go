// Package tree implements the tree building logic for an unsorted set of records.
package tree

import (
	"sort"
)

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

	if len(records) == 0 {
		return nil, nil
	}

	node := &Node{}

	calcNodes(records, node)

	//fmt.Println(record)
	//if record.ID == 0 && record.Parent !=0 {
	//	return nil, errors.New("error building tree")
	//}
	//}

	return node, nil
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
