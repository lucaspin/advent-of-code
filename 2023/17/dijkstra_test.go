package pkg202317

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test__ShortestPath(t *testing.T) {
	graph := Graph{}
	graph.Add(Node{
		ID: "A",
		Edges: []Edge{
			{Destination: "B", Weight: 2},
			{Destination: "D", Weight: 8},
		},
	})

	graph.Add(Node{
		ID: "B",
		Edges: []Edge{
			{Destination: "A", Weight: 2},
			{Destination: "D", Weight: 5},
			{Destination: "E", Weight: 6},
		},
	})

	graph.Add(Node{
		ID: "D",
		Edges: []Edge{
			{Destination: "A", Weight: 8},
			{Destination: "B", Weight: 5},
			{Destination: "D", Weight: 3},
			{Destination: "F", Weight: 2},
		},
	})

	graph.Add(Node{
		ID: "E",
		Edges: []Edge{
			{Destination: "B", Weight: 6},
			{Destination: "D", Weight: 3},
			{Destination: "F", Weight: 1},
			{Destination: "C", Weight: 9},
		},
	})

	graph.Add(Node{
		ID: "F",
		Edges: []Edge{
			{Destination: "D", Weight: 2},
			{Destination: "E", Weight: 1},
			{Destination: "C", Weight: 3},
		},
	})

	graph.Add(Node{
		ID: "C",
		Edges: []Edge{
			{Destination: "E", Weight: 9},
			{Destination: "F", Weight: 3},
		},
	})

	require.Equal(t, 12, graph.ShortestPath("A", "C"))
}
