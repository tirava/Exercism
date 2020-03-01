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
			if nextElem.right == nil {
				nextElem.right = newElem
				break
			}

			nextElem = nextElem.right
			continue
		}

		if nextElem.left == nil {
			nextElem.left = newElem
			break
		}

		nextElem = nextElem.left
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
