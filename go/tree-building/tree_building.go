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

// NewNode returns an node with empty value
func NewNode() *Node { return &Node{} }

// Build takes an array of records and returns the root node that contains all its children
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })
	treeMap := make(map[int]*Node, len(records))
	for i, v := range records {
		if i == 0 && v.ID != 0 && v.Parent != 0 {
			return nil, errors.New("Error with the root node record")
		}
		if v.ID != 0 && v.ID == v.Parent || v.ID < v.Parent || i != v.ID {
			return nil, errors.New("Record ID error")
		}
		if found := treeMap[v.ID]; found != nil {
			return nil, errors.New("Duplicate record found")
		}
		newNode := NewNode()
		newNode.ID = v.ID
		treeMap[v.ID] = newNode
		if v.ID != 0 {
			parentNode := treeMap[v.Parent]
			parentNode.Children = append(parentNode.Children, newNode)
		}
	}
	return treeMap[0], nil
}
