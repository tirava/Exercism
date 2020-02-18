// Package linkedlist implements simple linked list that uses Elements and a List.
package linkedlist

// Element base type.
type Element struct {
	data int
	next *Element
}

// List base type.
type List struct {
	head *Element
	size int
}

// New returns new list from slice.
func New(from []int) *List {
	return &List{}
}

// Size returns size of the list.
func (l *List) Size() int {
	return 0
}

// Push adds element to the list.
func (l *List) Push(elem int) {

}

// Pop removes element from the list.
func (l *List) Pop() (int, error) {
	return 0, nil
}

// Array returns slice from the list.
func (l *List) Array() []int {
	return nil
}

// Reverse reverses the list.
func (l *List) Reverse() *List {
	return &List{}
}
