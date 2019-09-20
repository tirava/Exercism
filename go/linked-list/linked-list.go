// Package linkedlist implements a doubly linked list.
package linkedlist

import "errors"

// ErrEmptyList is an error to return on an operation working on an empty list
var ErrEmptyList = errors.New("list is empty")

// Element is a node of the doubly linked list
type Element struct {
	Val        interface{}
	prev, next *Element
}

// List is a structure representing a doubly linked list
type List struct {
	top, bottom *Element
}

// Next returns a pointer to the next element of the node
func (e *Element) Next() *Element { return e.next }

// Prev returns a pointer to the previous element of the node
func (e *Element) Prev() *Element { return e.prev }

// NewList creates a new doubly linked list with the provided input elements
func NewList(args ...interface{}) *List {
	l := new(List)
	for _, item := range args {
		l.PushBack(item)
	}
	return l
}

// PushFront pushes an element to the front of the list
func (l *List) PushFront(v interface{}) {
	node := &Element{Val: v, next: l.top}
	if l.bottom == nil && l.top == nil {
		l.top, l.bottom = node, node
		return
	}
	l.top.prev = node
	l.top = node
}

// PushBack pushes an element to the back of the list
func (l *List) PushBack(v interface{}) {
	node := &Element{Val: v, prev: l.bottom}
	if l.bottom == nil && l.top == nil {
		l.top, l.bottom = node, node
		return
	}
	l.bottom.next = node
	l.bottom = node
}

func (l *List) popCornerCase() (interface{}, error) {
	if l.top == nil {
		return nil, ErrEmptyList
	}
	val := l.top.Val
	l.top, l.bottom = nil, nil
	return val, nil
}

// PopFront pops an element from the back of the list
func (l *List) PopFront() (interface{}, error) {
	if l.top == l.bottom {
		return l.popCornerCase()
	}
	val := l.top.Val
	l.top = l.top.next
	l.top.prev = nil
	return val, nil
}

// PopBack pops an element from the back of the list
func (l *List) PopBack() (interface{}, error) {
	if l.top == l.bottom {
		return l.popCornerCase()
	}
	val := l.bottom.Val
	l.bottom = l.bottom.prev
	l.bottom.next = nil
	return val, nil
}

// Reverse returns a list with the elements reversed
func (l *List) Reverse() *List {
	revL := l
	itNode := revL.top
	for itNode != nil {
		itNode.prev, itNode.next = itNode.next, itNode.prev
		itNode = itNode.prev
	}
	l.top, l.bottom = l.bottom, l.top
	return revL
}

// First return a pointer to the element at the front of the list
func (l *List) First() *Element { return l.top }

// Last return a pointer to the element at the back of the list
func (l *List) Last() *Element { return l.bottom }

//var ErrEmptyList = errors.New("the linked list is empty")
//
//type Element struct {
//	prev, next *Element
//	Val        interface{}
//}
//
//type List struct {
//	head, tail *Element
//}
//
//func (el *Element) Next() *Element {
//	if el == nil {
//		return nil
//	}
//	return el.next
//}
//
//func (el *Element) Prev() *Element {
//	if el == nil {
//		return nil
//	}
//	return el.prev
//}
//
//// it bugs me that this isn't named linkedlist.New, but this is required for the tests
//func NewList(elements ...interface{}) *List {
//	l := &List{}
//	for _, v := range elements {
//		l.PushBack(v)
//	}
//	return l
//}
//
//func (l *List) First() *Element {
//	return l.head
//}
//
//func (l *List) Last() *Element {
//	return l.tail
//}
//
//func (l *List) PushFront(v interface{}) {
//	el := &Element{next: l.head, Val: v}
//	if l.head != nil {
//		l.head.prev = el
//	} else {
//		l.tail = el
//	}
//	l.head = el
//}
//
//func (l *List) PopFront() (interface{}, error) {
//	if l.head == nil {
//		return 0, ErrEmptyList
//	}
//	v := l.head.Val
//	l.head = l.head.next
//	if l.head == nil {
//		l.tail = nil
//	} else {
//		l.head.prev = nil
//	}
//	return v, nil
//}
//
//func (l *List) PushBack(v interface{}) {
//	el := &Element{prev: l.tail, Val: v}
//	if l.tail != nil {
//		l.tail.next = el
//	} else {
//		l.head = el
//	}
//	l.tail = el
//}
//
//func (l *List) PopBack() (interface{}, error) {
//	if l.tail == nil {
//		return 0, ErrEmptyList
//	}
//	v := l.tail.Val
//	l.tail = l.tail.prev
//	if l.tail == nil {
//		l.head = nil
//	} else {
//		l.tail.next = nil
//	}
//	return v, nil
//}
//
//func (l *List) Reverse() {
//	for el := l.head; el != nil; el = el.prev {
//		el.next, el.prev = el.prev, el.next
//	}
//	l.head, l.tail = l.tail, l.head
//}
