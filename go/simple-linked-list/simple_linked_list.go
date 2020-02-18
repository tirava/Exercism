// Package linkedlist implements simple linked list that uses Elements and a List.
package linkedlist

import (
	"errors"
)

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
	list := &List{}

	for _, data := range from {
		list.Push(data)
	}

	return list
}

func (l *List) back(backIndex int) *Element {
	back := l.head

	for i := 0; i < l.size-(backIndex+1); i++ {
		back = back.next
	}

	return back
}

// Size returns size of the list.
func (l *List) Size() int {
	return l.size
}

// Push adds element to the list.
func (l *List) Push(data int) {
	elem := &Element{data: data}

	if l.size > 0 {
		l.back(0).next = elem
	} else {
		l.head = elem
	}

	l.size++
}

// Pop removes element from the list.
func (l *List) Pop() (int, error) {
	if l.size <= 0 {
		return 0, errors.New("empty list")
	}

	var data int
	if l.size == 1 {
		data = l.head.data
		l.head = nil
	} else {
		data = l.back(0).data
		l.back(1).next = nil
	}

	l.size--

	return data, nil
}

// Array returns slice from the list.
func (l *List) Array() []int {
	arr := make([]int, l.size)
	next := l.head

	for i := 0; i < l.size; i++ {
		arr[i] = next.data
		next = next.next
	}

	return arr
}

// Reverse reverses the list.
func (l *List) Reverse() *List {
	list := &List{}
	s := l.Array()

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	for _, v := range s {
		list.Push(v)
	}

	return list
}
