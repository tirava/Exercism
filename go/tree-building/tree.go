// Package tree implements the tree building logic for an unsorted set of records.
package tree

import "sort"

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

	//println("-----------")

	for _, record := range records {
		// find all unique parents
		// and append new Node fo every
		// sort

		if node.ID == record.Parent && record.ID != record.Parent {
			node.Children = append(node.Children, &Node{ID: record.ID})
		}

		sort.Slice(node.Children, func(i, j int) bool {
			return node.Children[i] < node.Children[j]
		})

		//fmt.Println(record)
		//if record.ID == 0 && record.Parent !=0 {
		//	return nil, errors.New("error building tree")
		//}
	}

	return node, nil
}
