package linkedlist

import (
	"errors"
	"fmt"
)

var ErrEmptyList = errors.New("Empty list")

type Node struct {
	Val  interface{}
	prev *Node
	next *Node
}

func (n *Node) String() string {
	return fmt.Sprintf("{Value:%d, Prev:%s, Next:%s}", n.Val, n.prev, n.next)
}

type List struct {
	head *Node
	tail *Node
}

func (e *Node) Next() *Node {
	return e.next
}

func (e *Node) Prev() *Node {
	return e.prev
}

func NewList(args ...interface{}) *List {
	newList := &List{}
	for _, v := range args {
		newList.PushBack(v)
	}

	currNode := newList.head
	for currNode != nil {
		currNode = currNode.next
	}
	return newList
}

// PushFront adds a node to the front of the list
func (l *List) PushFront(v interface{}) {
	newNode := &Node{Val: v, next: l.head}
	if l.head == nil {
		l.tail = newNode
	} else {
		l.head.prev = newNode
	}
	l.head = newNode
}

// PushBack adds a node to the back of the list
func (l *List) PushBack(v interface{}) {
	newNode := &Node{Val: v, prev: l.tail}
	if l.tail == nil {
		l.head = newNode
	} else {
		l.tail.next = newNode
	}
	l.tail = newNode
}

// PopFront removes a node from the front of the list
func (l *List) PopFront() (interface{}, error) {
	if l.head == nil {
		return nil, ErrEmptyList
	}
	dataPopped := l.head.Val
	if l.head.next == nil {
		l.head = nil
		l.tail = nil
		return dataPopped, nil
	}
	l.head.Next().prev = nil
	l.head = l.head.next
	return dataPopped, nil
}

// PopBack removes the last node in the list
func (l *List) PopBack() (interface{}, error) {
	if l.head == nil {
		return nil, ErrEmptyList
	}
	dataPopped := l.Last().Val
	if l.tail.Prev() == nil {
		l.head = nil
		l.tail = nil
		return dataPopped, nil
	}
	l.tail.prev.next = nil
	l.tail = l.tail.prev
	return dataPopped, nil
}

//Reverse flips the order of the list
func (l *List) Reverse() *List {
	if l.head == nil {
		return nil
	}
	for n := l.head; n != nil; n = n.prev {
		n.prev, n.next = n.next, n.prev
	}
	l.head, l.tail = l.tail, l.head
	return l
}

// First return the head of the doubly linked list
func (l *List) First() *Node {
	if l == nil {
		return nil
	}
	return l.head
}

// Last returns the tail of the list
func (l *List) Last() *Node {
	if l == nil {
		return nil
	}
	return l.tail
}
