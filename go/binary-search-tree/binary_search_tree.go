// Package binarysearchtree implements inserting and searching for numbers in a binary tree.
package binarysearchtree

// SearchTreeData base struct.
type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

// Bst returns new tree.
func Bst(root int) *SearchTreeData {
	return &SearchTreeData{
		data: root,
	}
}

// Insert inserts value into tree.
func (std *SearchTreeData) Insert(elem int) {
	newElem := Bst(elem)
	nextElem := std

	for {
		if elem > nextElem.data {
			nextElem = nextElem.right
			//if std.right == nil {
			//	std.right = newElem
			//}
			//std.right = nextElem
		} else {
			nextElem = nextElem.left
		}

		if elem <= std.data && std.left == nil {
			std.left = newElem
			continue
		}

		if
	}
}

// MapString maps tree to string slice.
func (std *SearchTreeData) MapString(func(int) string) []string {
	return nil
}

// MapInt maps tree to int slice.
func (std *SearchTreeData) MapInt(func(int) int) []int {
	return nil
}
