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

	// check for valid node where ID!=ParentID
	m := make(map[int]int, len(records))
	var recordIds []int
	for i := 0; i < len(records); i++ {
		if records[i].ID != 0 && records[i].ID == records[i].Parent {
			return nil, errors.New("Record ID and parent ID can't be the same")
		}

		if records[i].ID < records[i].Parent {
			return nil, errors.New(("Record ID can't be higher than its parent ID"))
		}
		// check for duplicates
		m[records[i].ID]++
		recordIds = append(recordIds, records[i].ID)
		if m[records[i].ID] > 1 {
			return nil, errors.New("Duplicate record found")
		}
	}
	sort.Ints(recordIds)

	// check for continuous records
	for j := 0; j < len(recordIds); j++ {
		if j != recordIds[j] {
			return nil, errors.New("Non-continuous record found")
		}
	}

	root := NewNode()
	found, numOfRoots := ContainsRecord(records, 0)
	if !found {
		return nil, errors.New("No root node")
	}
	if numOfRoots > 1 {
		return nil, errors.New("Duplicate root found")
	}

	for _, value := range records {
		if value.ID == 0 {
			if value.Parent != 0 {
				return nil, errors.New("Root should not have a parent")
			}
		}
	}

	// find all the unique parent IDs and then build tree based on them
	var parentIDs []int
	for _, value := range records {
		if !ContainsParentID(parentIDs, value.Parent) && value.ID != 0 {
			parentIDs = append(parentIDs, value.Parent)
		}
	}
	sort.Ints(parentIDs)

	for _, parentID := range parentIDs {
		parentNode := FindNodeByID(root, parentID)
		if parentNode != nil {
			var childIds []int
			for _, value := range records {
				if value.Parent == parentID && value.ID != 0 {
					childIds = append(childIds, value.ID)

				}
			}
			// sort the ids
			sort.Ints(childIds)

			for _, id := range childIds {
				var newNode = NewNode()
				newNode.ID = id
				parentNode.Children = append(parentNode.Children, newNode)
			}
		}
	}

	return root, nil
}

// FindNodeByID returns a pointer to a node with ID=id from the root node
func FindNodeByID(root *Node, id int) *Node {
	if root.ID == id {
		return root
	} else if root.Children != nil {
		for _, v := range root.Children {
			if found := FindNodeByID(v, id); found != nil {
				return found
			}
		}
	}
	return nil // return nil if not found
}

// ContainsParentID checks whether an id is a parent
func ContainsParentID(parentIDs []int, id int) bool {
	for _, value := range parentIDs {
		if value == id {
			return true
		}
	}
	return false
}

// ContainsRecord returns a bool if records contains a particular record ID
func ContainsRecord(records []Record, recordID int) (bool, int) {
	var numberOfRecords int
	for _, value := range records {
		if value.ID == recordID {
			numberOfRecords++
		}
	}
	if numberOfRecords >= 1 {
		return true, numberOfRecords
	}
	return false, 0
}
