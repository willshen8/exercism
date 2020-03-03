package tree

import (
	"errors"
	"sort"
)

// Record is a struct containing int fields ID and Parent
type Record struct {
	ID     int
	Parent int
}

type Records []Record

// Node is a struct containing int field ID and []*Node field Children.
type Node struct {
	ID       int
	Children []*Node
}

// Len() returns the length of the records
// It implements the sort interface in Go
func (r Records) Len() int {
	return len(r)
}

// Swap is an implementation of the sort interface method
// It implements the sort interfaces in Go
func (r Records) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

// Less reports whether the element with index i should sort before the element with index j.
// We sort it based on lower Parent IDs first, so they can be inserted in order
// It implements the sort interfaces in Go
func (r Records) Less(i, j int) bool {
	return r[i].ID < r[j].ID
}

// NewNode returns an node with empty value
func NewNode() *Node { return &Node{} }

// Build takes an array of records and returns the root node that contains all its children
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	//sort the records based on increasing ID order
	sort.Sort(Records(records))
	// stores the records in a map for easy retrival
	treeMap := make(map[int]*Node, len(records))
	var root *Node
	for i, v := range records {
		if i == 0 && v.ID != 0 {
			return nil, errors.New("No root node exist")
		}
		if i == 0 && v.Parent != 0 {
			return nil, errors.New("Root should not have a parent")
		}
		if v.ID != 0 && v.ID == v.Parent {
			return nil, errors.New("Record ID and parent ID can't be the same")
		}
		if i != v.ID {
			return nil, errors.New("Non-continuous record found")
		}
		if v.ID < v.Parent {
			return nil, errors.New(("Record ID can't be lower than its parent ID"))
		}
		if found := treeMap[v.ID]; found != nil {
			return nil, errors.New("Duplicate record found")
		}
		newNode := NewNode()
		newNode.ID = v.ID
		treeMap[v.ID] = newNode
		if v.ID == 0 {
			root = newNode
		} else if v.ID != 0 {
			parentNode := treeMap[v.Parent]
			parentNode.Children = append(parentNode.Children, newNode)
		}
	}
	return root, nil
}
