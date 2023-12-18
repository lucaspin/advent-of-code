package pkg202318

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test__SolutionA(t *testing.T) {
	require.Equal(t, 62.0, A("example.txt"))
	require.Equal(t, 61865.0, A("input.txt"))
}

func Test__SolutionB(t *testing.T) {
	require.Equal(t, 952408144115.0, B("example.txt"))
	require.Equal(t, 4.0343619199142e+13, B("input.txt"))
}

func Test__Shoelace(t *testing.T) {
	p := &Polygon{Vertices: []Vertice{}}
	p.AddVertice(Vertice{X: 1, Y: 6})
	p.AddVertice(Vertice{X: 8, Y: 5})
	p.AddVertice(Vertice{X: 4, Y: 4})
	p.AddVertice(Vertice{X: 7, Y: 2})
	p.AddVertice(Vertice{X: 3, Y: 1})
	require.Equal(t, 16.5, p.Area())
}
