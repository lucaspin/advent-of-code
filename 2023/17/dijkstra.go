package pkg202317

import (
	"fmt"
	"math"

	"github.com/lucaspin/computer-science/pkg/binary_heap"
)

type Node struct {
	ID    string
	Edges []Edge
}

func (n *Node) Destinations() []string {
	ns := []string{}
	for _, e := range n.Edges {
		ns = append(ns, e.Destination)
	}

	return ns
}

type Edge struct {
	Destination string
	Direction   string
	Weight      int
}

type Graph map[string]Node

func (s *Graph) Get(ID string) Node {
	return (*s)[ID]
}

func (s *Graph) Add(node Node) {
	if s.Has(node) {
		return
	}

	(*s)[node.ID] = node
}

func (s *Graph) Has(node Node) bool {
	_, ok := (*s)[node.ID]
	return ok
}

func (g *Graph) AllNodes() []string {
	s := []string{}
	for k := range *g {
		s = append(s, k)
	}

	return s
}

func (g *Graph) ShortestPathModified(start, end string) (int, []string) {
	allNodes := g.AllNodes()

	// build initial tentative distances map
	distances := map[string]int{}
	previous := map[string]string{}
	for _, v := range allNodes {
		distances[v] = math.MaxInt32 // infinite
	}
	distances[start] = 0

	type N struct {
		ID             string
		Value          int
		Direction      string
		DirectionCount int
	}

	unvisited := binary_heap.NewBinaryHeap[N]([]N{}, func(a, b N) bool { return a.Value < b.Value })
	for _, k := range allNodes {
		unvisited.Push(N{ID: k, Value: distances[k]})
	}

	// iterate until there are no unvisited nodes
	for unvisited.Len() > 0 {
		// Pick the next node, which will be the unvisited node with the minimum difference.
		current := unvisited.Pop()

		// Same node can be added to the queue multiple times.
		// We are only interested on the smallest known distance for each node,
		// so we ignore it if it's greater than the current known smallest difference.
		if current.Value > distances[current.ID] {
			continue
		}

		fmt.Printf("current=%s count=%d dir=%s\n", current.ID, current.DirectionCount, current.Direction)
		node := g.Get(current.ID)

		// update distances for each adjacent node for the current node
		for _, e := range node.Edges {
			if current.Direction == e.Direction && current.DirectionCount >= 2 {
				fmt.Printf("%s -> %s - ignoring due to count=%d\n", current.ID, e.Destination, current.DirectionCount)
				continue
			}

			fmt.Printf("%s -> %s - OK\n", current.ID, e.Destination)
			dist := distances[current.ID] + e.Weight
			if dist < distances[e.Destination] {
				var dCount int
				if current.Direction == e.Direction {
					dCount = current.DirectionCount + 1
				} else {
					dCount = 0
				}

				distances[e.Destination] = dist
				previous[e.Destination] = node.ID
				unvisited.Push(N{
					ID:             e.Destination,
					Value:          dist,
					Direction:      e.Direction,
					DirectionCount: dCount,
				})
			}
		}
	}

	path := []string{}
	for i := end; i != start; {
		path = append(path, i)
		i = previous[i]
	}

	return distances[end], path
}

func (g *Graph) ShortestPath(start, end string) int {
	allNodes := g.AllNodes()

	// build initial tentative distances map
	distances := map[string]int{}
	for _, v := range allNodes {
		distances[v] = math.MaxInt32 // infinite
	}
	distances[start] = 0

	type N struct {
		ID    string
		Value int
	}

	unvisited := binary_heap.NewBinaryHeap[N]([]N{}, func(a, b N) bool { return a.Value < b.Value })
	for _, k := range allNodes {
		unvisited.Push(N{ID: k, Value: distances[k]})
	}

	// iterate until there are no unvisited nodes
	for unvisited.Len() > 0 {
		// Pick the next node, which will be the unvisited node with the minimum difference.
		current := unvisited.Pop()

		// Same node can be added to the queue multiple times.
		// We are only interested on the smallest known distance for each node,
		// so we ignore it if it's greater than the current known smallest difference.
		if current.Value > distances[current.ID] {
			continue
		}

		fmt.Printf("current=%s unvisited=%d\n", current.ID, unvisited.Len())
		node := g.Get(current.ID)

		// update distances for each adjacent node for the current node
		for _, e := range node.Edges {
			dist := distances[current.ID] + e.Weight
			if dist < distances[e.Destination] {
				distances[e.Destination] = dist
				unvisited.Push(N{ID: e.Destination, Value: dist})
			}
		}
	}

	return distances[end]
}
