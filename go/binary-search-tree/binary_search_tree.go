package binarysearchtree

// SearchTreeData represents a node in the binary search tree
type SearchTreeData struct {
	data  int
	left  *SearchTreeData
	right *SearchTreeData
}

// Bst creates a new binary search tree with root value specifies in the argument
func Bst(key int) *SearchTreeData {
	return &SearchTreeData{data: key}
}

// Insert adds a new node into the bst
func (root *SearchTreeData) Insert(key int) {
	newNode := &SearchTreeData{data: key}
	if key <= root.data {
		if root.left == nil {
			root.left = newNode
		} else {
			root.left.Insert(key)
		}
	} else if key > root.data {
		if root.right == nil {
			root.right = newNode
		} else {
			root.right.Insert(key)
		}
	}
}

// MapString prints out all the nodes as string
func (root *SearchTreeData) MapString(stringfy func(int) string) []string {
	var results []string
	if root.left != nil {
		results = append(results, root.left.MapString(stringfy)...)
	}
	results = append(results, stringfy(root.data))
	if root.right != nil {
		results = append(results, root.right.MapString(stringfy)...)
	}
	return results
}

// MapInt prints out all the nodes as integers
func (root *SearchTreeData) MapInt(f func(int) int) []int {
	var results []int
	if root.left != nil {
		results = append(results, root.left.MapInt(f)...)
	}
	results = append(results, f(root.data))
	if root.right != nil {
		results = append(results, root.right.MapInt(f)...)
	}
	return results
}
