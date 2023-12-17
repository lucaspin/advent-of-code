package pkg202317

import (
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

type StringSet map[string]bool

func (s *StringSet) Has(v string) bool {
	_, ok := (*s)[v]
	return ok
}

func (s *StringSet) Add(v string) {
	(*s)[v] = true
}

type NSet map[N]bool

func (n *NSet) Has(v N) bool {
	_, ok := (*n)[v]
	return ok
}

func (n *NSet) Add(v N) {
	(*n)[v] = true
}

func (g *Graph) ShortestPath(start, end string) int {
	type N struct {
		ID    string
		Value int
	}

	distances := map[string]int{}
	distances[start] = 0

	queue := binary_heap.NewBinaryHeap[N]([]N{}, func(a, b N) bool { return a.Value < b.Value })
	queue.Push(N{ID: start, Value: distances[start]})

	visited := StringSet{}

	for queue.Len() > 0 {
		current := queue.Pop()
		if current.Value > distances[current.ID] {
			continue
		}

		if visited.Has(current.ID) {
			continue
		}

		visited.Add(current.ID)

		// update distances for each adjacent node for the current node
		for _, e := range g.Get(current.ID).Edges {
			dist := distances[current.ID] + e.Weight
			if v, ok := distances[e.Destination]; ok {
				if dist < v {
					distances[e.Destination] = dist
					queue.Push(N{ID: e.Destination, Value: dist})
				}
			} else {
				distances[e.Destination] = dist
				queue.Push(N{ID: e.Destination, Value: dist})
			}
		}
	}

	return distances[end]
}

type N struct {
	ID             string
	Direction      string
	DirectionCount int
}

func (g *Graph) ShortestPathModified(start, end string, part2 bool) int {
	n0 := N{
		ID:             start,
		Direction:      "east",
		DirectionCount: 0,
	}

	n1 := N{
		ID:             start,
		Direction:      "south",
		DirectionCount: 0,
	}

	distances := map[N]int{}
	distances[n0] = 0
	distances[n1] = 0

	queue := binary_heap.NewBinaryHeap[N]([]N{}, func(a, b N) bool { return distances[a] < distances[b] })
	queue.Push(n0)
	queue.Push(n1)

	visited := NSet{}

	for queue.Len() > 0 {
		current := queue.Pop()
		if visited.Has(*current) {
			continue
		}

		visited.Add(*current)

		// update distances for each adjacent node for the current node
		for _, e := range filter(g.Get(current.ID).Edges, *current, part2) {
			dist := distances[*current] + e.Weight
			n := buildN(*current, e)

			if v, ok := distances[n]; ok {
				if dist < v {
					distances[n] = dist
					queue.Push(n)
				}
			} else {
				distances[n] = dist
				queue.Push(n)
			}
		}
	}

	return min(distances, end, part2)
}

func filterForPart2(edges []Edge, current N) []Edge {
	es := []Edge{}
	for _, e := range edges {
		switch e.Direction {
		case "south":
			if current.Direction == "south" {
				if current.DirectionCount < 10 {
					es = append(es, e)
				}
			} else if current.Direction != "north" && current.DirectionCount >= 4 {
				es = append(es, e)
			}
		case "north":
			if current.Direction == "north" {
				if current.DirectionCount < 10 {
					es = append(es, e)
				}
			} else if current.Direction != "south" && current.DirectionCount >= 4 {
				es = append(es, e)
			}
		case "east":
			if current.Direction == "east" {
				if current.DirectionCount < 10 {
					es = append(es, e)
				}
			} else if current.Direction != "west" && current.DirectionCount >= 4 {
				es = append(es, e)
			}
		case "west":
			if current.Direction == "west" {
				if current.DirectionCount < 10 {
					es = append(es, e)
				}
			} else if current.Direction != "east" && current.DirectionCount >= 4 {
				es = append(es, e)
			}
		}
	}

	// fmt.Printf("current=%v edges=%v\n", current, es)

	return es
}

func filter(edges []Edge, current N, part2 bool) []Edge {
	if part2 {
		return filterForPart2(edges, current)
	}

	es := []Edge{}
	for _, e := range edges {
		switch e.Direction {
		case "south":
			if current.Direction == "east" || current.Direction == "west" || (current.Direction == "south" && current.DirectionCount < 3) {
				es = append(es, e)
			}
		case "north":
			if current.Direction == "east" || current.Direction == "west" || (current.Direction == "north" && current.DirectionCount < 3) {
				es = append(es, e)
			}
		case "east":
			if current.Direction == "north" || current.Direction == "south" || (current.Direction == "east" && current.DirectionCount < 3) {
				es = append(es, e)
			}
		case "west":
			if current.Direction == "north" || current.Direction == "south" || (current.Direction == "west" && current.DirectionCount < 3) {
				es = append(es, e)
			}
		}
	}

	return es
}

func min(distances map[N]int, ID string, part2 bool) int {
	if part2 {
		min := math.MaxInt32
		for k, v := range distances {
			if k.ID == ID && k.DirectionCount >= 4 {
				if v < min {
					min = v
				}
			}
		}

		return min
	}

	min := math.MaxInt32
	for k, v := range distances {
		if k.ID == ID {
			if v < min {
				min = v
			}
		}
	}

	return min
}

func buildN(current N, e Edge) N {
	n := N{
		ID:        e.Destination,
		Direction: e.Direction,
	}

	if current.Direction == e.Direction {
		n.DirectionCount = current.DirectionCount + 1
	} else {
		n.DirectionCount = 1
	}

	return n
}
