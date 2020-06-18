package pov

import (
	"fmt"
	"sync"
)

type Node struct {
	label string
}

type Arc struct {
	from string
	to   string
}

type Graph struct {
	nodes []Node
	arcs  []Arc
	lock  sync.RWMutex
}

func New() *Graph { return &Graph{} }

// AddNode adds a new node to the graph
func (g *Graph) AddNode(nodeLabel string) {
	g.lock.Lock()
	defer g.lock.Unlock()
	g.nodes = append(g.nodes, Node{label: nodeLabel})
}

// AddArc adds an arc/edge to the graph
func (g *Graph) AddArc(from, to string) {
	g.lock.Lock()
	defer g.lock.Unlock()
	g.arcs = append(g.arcs, Arc{from: from, to: to})
}

// ArcList returns a dump of all the arcs in the graph
func (g *Graph) ArcList() []string {
	var arcList []string
	for _, arc := range g.arcs {
		arcList = append(arcList, fmt.Sprintf("%v -> %v", arc.from, arc.to))
	}
	return arcList
}

// Changeroot make the newRoot the newRoot and changes all directed arcs as well
func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	for i, v := range g.arcs {
		if v.from == newRoot {
			g.arcs[i].from, g.arcs[i].to = g.arcs[i].to, g.arcs[i].from
		}
	}
	return g
}
