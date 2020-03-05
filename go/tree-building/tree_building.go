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
		if v.ID != i || v.Parent > v.ID || v.ID > 0 && v.Parent == v.ID {
			return nil, errors.New("not in sequence or has bad parent")
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
