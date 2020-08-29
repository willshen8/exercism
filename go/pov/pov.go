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

// ChangeRoot make the newRoot the newRoot and changes all directed arcs as well
func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	fromNodes := make([]string, 0, len(g.arcs))
	for k := range g.arcs {
		fromNodes = append(fromNodes, k)
	}
	var arcsToBeChanged []Arc
	currNode := newRoot
	for currNode != oldRoot {
		for _, fromNode := range fromNodes {
			for _, toNode := range g.arcs[fromNode] {
				if currNode == toNode {
					arcsToBeChanged = append(arcsToBeChanged, Arc{fromNode, toNode})
					currNode = fromNode
				}
			}
		}
	}
	for _, arc := range arcsToBeChanged {
		g.RemoveArc(arc.from, arc.to)
		g.AddArc(arc.to, arc.from)
	}
	return g
}
