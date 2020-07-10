package pov

import (
	"fmt"
	"sync"
)

type Node struct {
	label string
}

type Graph struct {
	nodes []*Node
	arcs  map[string][]string
	lock  sync.RWMutex
}

type Arc struct {
	from, to string
}

func New() *Graph {
	return &Graph{arcs: make(map[string][]string)}
}

// AddNode adds a new node to the graph
func (g *Graph) AddNode(nodeLabel string) {
	g.lock.Lock()
	defer g.lock.Unlock()
	g.nodes = append(g.nodes, &Node{label: nodeLabel})
}

// AddArc adds an arc/edge to the graph
func (g *Graph) AddArc(from, to string) {
	g.lock.Lock()
	defer g.lock.Unlock()
	g.arcs[from] = append(g.arcs[from], to)
}

// RemoveArc removes an arc/edge from the graph
func (g *Graph) RemoveArc(from, to string) {
	g.lock.Lock()
	defer g.lock.Unlock()
	toArcs := g.arcs[from]
	for i, v := range toArcs {
		if v == to {
			g.arcs[from][i] = g.arcs[from][len(g.arcs[from])-1]
			g.arcs[from] = g.arcs[from][:len(g.arcs[from])-1]
		}
	}
}

// ArcList returns a dump of all the arcs in the graph
func (g *Graph) ArcList() []string {
	var arcList []string
	for from, arcs := range g.arcs {
		for _, to := range arcs {
			arcList = append(arcList, fmt.Sprintf("%v -> %v", from, to))
		}
	}
	return arcList
}

// FindPathVertices recursively finds a list of vertices between two vertices: from and to
// and stores the results in a map called path
func (g *Graph) FindPathVertices(from, to string) (path []Arc) {
	keys := make([]string, 0, len(g.arcs))
	for k := range g.arcs {
		keys = append(keys, k)
	}

	var found string
	if _, ok := g.arcs[to]; !ok {
		return path
	}
	for _, k := range keys {
		for _, arc := range g.arcs[k] {
			if arc == to {
				path = append(path, Arc{k, to})
				found = k
				fmt.Println("Found=", found)
			}
		}
	}

	// if found != from {
	// 	path = g.FindPathVertices(from, found)
	// }
	return path
}

// ChangeRoot make the newRoot the newRoot and changes all directed arcs as well
func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	// first find all the nodes start from oldRoot to newRoot
	path := g.FindPathVertices(oldRoot, newRoot)
	fmt.Println("Path:", path)
	if verticesTo, ok := g.arcs[oldRoot]; ok {
		for _, v := range verticesTo {
			if v == newRoot {
				g.AddArc(newRoot, oldRoot)
				g.RemoveArc(oldRoot, newRoot)
			}
		}
	}
	return g
}

// FindPathBetweenTwoVertices returns a list of all nodes between two vertices
// func (g *Graph) FindPathBetweenTwoVertices(from, to string) []string {
// 	var pathVertices []string
// 	for k, _ := range g.arcs[{
// 		pathVertices = append(pathVertices, k)
// 	}
// }
