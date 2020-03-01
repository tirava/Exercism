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
func (std *SearchTreeData) MapString(convert func(int) string) []string {
	result := make([]string, 0)
	in := make([]int, 0)

	for _, v := range *visitRecursive(std, &in) {
		result = append(result, convert(v))
	}

	return result
}

// MapInt maps tree to int slice.
func (std *SearchTreeData) MapInt(convert func(int) int) []int {
	result := make([]int, 0)

	return *visitRecursive(std, &result)
}

func visitRecursive(root *SearchTreeData, in *[]int) *[]int {
	if root.left != nil {
		visitRecursive(root.left, in)
	}

	*in = append(*in, root.data)

	if root.right != nil {
		visitRecursive(root.right, in)
	}

	return in
}
