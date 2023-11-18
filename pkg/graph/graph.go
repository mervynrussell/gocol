package graph

type Graph[T comparable, D comparable] interface {
	Vertices() []T
	Edges() []Edge[T, D]
	VectorEdges(v T) []Edge[T, D]
	AddEdge(u T, v T, d D) (*Edge[T, D], error)
	AddVertex(v T) error
	RemoveVertex(v T)
	RemoveEdge(e Edge[T, D])
	ContainsVertex(v T) bool
	ContainsEdge(e Edge[T, D]) bool
}

type Edge[T comparable, D comparable] struct {
	u T
	v T
	d D
}

func NewEdge[T comparable, D comparable](u T, v T, d D) Edge[T, D] {
	return Edge[T, D]{u: u, v: v, d: d}
}