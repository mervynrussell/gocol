package graph

import (
	"container/list"
	"fmt"
)

type internalEdge[T comparable] struct {
	u T
	v T
}

func newInternalEdge[T comparable](u T, v T) internalEdge[T] {
	return internalEdge[T]{u, v}
}

// Undirected graph based on adjacency list implementation
type adjacencyListGraph[T comparable, D comparable] struct {
	vertices map[T]*list.List
	edges    map[internalEdge[T]]D
}

func NewAdjacencyListGraph[T comparable, D comparable]() *adjacencyListGraph[T, D] {
	return &adjacencyListGraph[T, D]{vertices: make(map[T]*list.List), edges: make(map[internalEdge[T]]D)}
}

func (g *adjacencyListGraph[T, D]) Vertices() []T {
	return mapKeys(g.vertices)
}

func (g *adjacencyListGraph[T, D]) AddEdge(u T, v T, d D) (*Edge[T, D], error) {
	if !g.ContainsVertex(u) {
		return nil, fmt.Errorf("unknown vertex %v:%T", u, u)
	}

	if !g.ContainsVertex(v) {
		return nil, fmt.Errorf("unknown vertex %v:%T", v, v)
	}

	edge := newInternalEdge(u, v)
	if _, ok := g.edges[edge]; !ok {
		g.edges[edge] = d
		g.vertices[u].PushBack(edge)
		g.vertices[v].PushBack(edge)
	} 
	r := NewEdge(u, v, g.edges[edge])
	return &r, nil
}

func (g *adjacencyListGraph[T, D]) Edges() []Edge[T, D] {
	edges := make([]Edge[T, D], len(g.edges))

	for i, e := range mapKeys(g.edges) {
		edges[i] = NewEdge[T, D](e.u, e.v, g.edges[e])
	}
	return edges
}

func (g *adjacencyListGraph[T, D]) VectorEdges(v T) []Edge[T, D] {
	l := g.vertices[v]
	edges := make([]Edge[T, D], l.Len())
	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		iEdge := e.Value.(internalEdge[T])
        edges[i] = NewEdge(iEdge.u, iEdge.v, g.edges[iEdge])
		i++
    }
	return edges
}

func (g *adjacencyListGraph[T, D]) AddVertex(v T) error {
	if g.ContainsVertex(v) {
		return fmt.Errorf("graph already contains vertex %v", v)
	}
	g.vertices[v] = list.New()
	return nil
}

func (g *adjacencyListGraph[T, D]) RemoveVertex(v T) {
	delete(g.vertices, v)
}

func (g *adjacencyListGraph[T, D]) RemoveEdge(e Edge[T, D]) {
	iEdge := newInternalEdge(e.u, e.v)
	element := elementFromList[internalEdge[T]](*g.vertices[e.u], iEdge)
	if element != nil {
		g.vertices[e.u].Remove(element)
	}
	
	element = elementFromList[internalEdge[T]](*g.vertices[e.v], iEdge)
	if element != nil {
		g.vertices[e.v].Remove(element)
	}

	delete(g.edges, newInternalEdge(e.u, e.v))
}

func (g *adjacencyListGraph[T, D]) ContainsVertex(v T) bool {
	_, ok := g.vertices[v]
	return ok
}

func (g *adjacencyListGraph[T, D]) ContainsEdge(e Edge[T, D]) bool {
	ie := newInternalEdge(e.u, e.v) 
	_, ok := g.edges[ie]
	return ok
}
