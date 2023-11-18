package graph

import (
	"fmt"
	"testing"
)

func TestVertices(t *testing.T) {
	var g Graph[int, string] = (*adjacencyListGraph[int, string])(nil)
	g = NewAdjacencyListGraph[int, string]()

	g.AddVertex(0)
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)

	vertices := g.Vertices()
	if len(vertices) != 5 {
		t.Fatalf("expected 5 got %d", len(vertices))
	}

	for i := 0; i <= 4; i++ {
		if !g.ContainsVertex(i) {
			t.Fatalf("expected true for contains '%d'", i)
		}
	}

	if err := g.AddVertex(0); err == nil {
		t.Fatal("expected error adding duplicate vertex 0")
	}

	g.RemoveVertex(4)
	if g.ContainsVertex(4) {
		t.Fatal("expect grap doe snot contain 4")
	}
}

func TestEdges(t *testing.T) {
	var g Graph[int, string] = (*adjacencyListGraph[int, string])(nil)
	g = NewAdjacencyListGraph[int, string]()

	g.AddVertex(0)
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)

	e1, err := g.AddEdge(0, 1, "0:1")
	if e1.u != 0 || e1.v != 1 || err != nil {
		t.Fatalf("edge error")
	}
	if !g.ContainsEdge(*e1) {
		t.Fatalf("expected graph contains edge %v", e1)
	}

	g.AddEdge(0, 2, "0:2")
	g.AddEdge(0, 3, "0:3")

	_, err = g.AddEdge(0, 4, "0:4")
	if err == nil {
		t.Fatalf("expected edge error for unknown vertex 4")
	}

	_, err = g.AddEdge(4, 0, "4:0")
	if err == nil {
		t.Fatalf("expected edge error for unknown vertex 4")
	}

	eDuplicate, err := g.AddEdge(0, 1, "fnarrp")
	if *eDuplicate != *e1 || err != nil {
		t.Fatalf("expected duplicate edge to be same and no error")
	}

	g.RemoveEdge(*eDuplicate)
	if g.ContainsEdge(*eDuplicate) {
		t.Fatalf("unexpected %v in graph", eDuplicate)
	}

	edges := g.Edges()
	if len(edges) != 2 {
		t.Fatalf("expected 2 got %d", len(edges))
	}
}

func TestLargeFullyConnected(t *testing.T) {
	var g Graph[int, string] = (*adjacencyListGraph[int, string])(nil)
	g = NewAdjacencyListGraph[int, string]()
	numVertices := 999
	for i := 0; i < numVertices; i++ {
		g.AddVertex(i)
	}
	fmt.Print("vertices added")

	for i := 0; i < numVertices; i++ {
		for j := i; j < numVertices; j++ {
			if i != j {
				g.AddEdge(i, j, fmt.Sprintf("%d:%d", i, j))
			}
			
		}
	}
	fmt.Println("edges added")

	expectedVertices := len(g.Vertices())
	if expectedVertices != numVertices {
		t.Fatalf("expected %d got %d", numVertices, expectedVertices)
	}
	fmt.Println("vertice check pass")

	expectedEdges := (numVertices * (numVertices - 1) ) /2
	actualEdges := len(g.Edges())
	if actualEdges != expectedEdges {
		t.Fatalf("expected %d got %d", expectedEdges, actualEdges)
	}
	fmt.Println("all edges check pass")

	expectedEdges = numVertices - 1
	actualEdges = len(g.VectorEdges(0))
	if actualEdges != expectedEdges {
		t.Fatalf("expected %d got %d", expectedEdges, actualEdges)
	}
	fmt.Println("vector edges check pass")

}
