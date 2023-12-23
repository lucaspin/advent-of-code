package pkg202323

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Node struct {
	ID   string
	Row  int
	Col  int
	Type string
}

type Graph struct {
	nodes  map[string]Node
	edges  map[string][]Edge
	height int
	width  int
}

type Edge struct {
	Destination string
	Direction   string
}

func (s *Graph) GetNode(ID string) Node {
	return s.nodes[ID]
}

func (s *Graph) GetDestinations(nodeID string) []Edge {
	return s.edges[nodeID]
}

func (s *Graph) AddNode(node Node) {
	if s.nodes == nil {
		s.nodes = map[string]Node{}
	}

	s.nodes[node.ID] = node
}

func (s *Graph) AddDestination(src string, edge Edge) {
	if s.edges == nil {
		s.edges = map[string][]Edge{}
	}

	if _, ok := s.edges[src]; ok {
		new := append(s.edges[src], edge)
		s.edges[src] = new
	} else {
		s.edges[src] = []Edge{edge}
	}
}

func (g *Graph) AllNodes() []string {
	s := []string{}
	for k := range g.nodes {
		s = append(s, k)
	}

	return s
}

func (g *Graph) findStart() Node {
	for i := 0; i < g.width; i++ {
		n := g.GetNode(nodeID(0, i))
		if n.Type == "." {
			return n
		}
	}

	panic("should not happen")
}

func (g *Graph) findEnd() Node {
	for i := 0; i < g.width; i++ {
		n := g.GetNode(nodeID(g.height-1, i))
		if n.Type == "." {
			return n
		}
	}

	panic("should not happen")
}

type Location struct {
	Row  int
	Col  int
	Type string
}

type queue struct {
	items []queueItem
}

func (q *queue) Push(item queueItem) {
	q.items = append(q.items, item)
}

func (q *queue) Pop() queueItem {
	head := q.items[0]
	q.items = q.items[1:]
	return head
}

func (q *queue) Len() int {
	return len(q.items)
}

type queueItem struct {
	node Node
	seen NodeSet
}

type NodeSet map[Node]bool

func (n *NodeSet) Has(v Node) bool {
	_, ok := (*n)[v]
	return ok
}

func (n *NodeSet) Add(v Node) {
	(*n)[v] = true
}

func (n *NodeSet) Len() int {
	return len(*n)
}

func (n *NodeSet) Values() []Node {
	ls := []Node{}
	for k := range *n {
		ls = append(ls, k)
	}

	return ls
}

func (n *NodeSet) Copy() NodeSet {
	new := make(map[Node]bool, len(*n))
	for k, v := range *n {
		new[k] = v
	}

	return new
}

func _search(graph Graph, queue queue, target Node) []int {
	paths := []int{}

	for queue.Len() > 0 {
		current := queue.Pop()

		// Reached bottom row
		if current.node == target {
			paths = append(paths, current.seen.Len()-1)
			continue
		}

		next := graph.GetDestinations(current.node.ID)
		if len(next) == 0 {
			continue
		}

		for _, n := range next {
			dst := graph.GetNode(n.Destination)
			if current.seen.Has(dst) {
				fmt.Printf("")
				continue
			}

			newSeen := current.seen.Copy()
			newSeen.Add(dst)

			queue.Push(queueItem{
				node: dst,
				seen: newSeen,
			})
		}
	}

	return paths
}

func longest(graph Graph) int {
	start := graph.findStart()
	end := graph.findEnd()

	seen := NodeSet{}
	seen.Add(start)

	queue := queue{}
	queue.Push(queueItem{
		node: start,
		seen: seen,
	})

	return max(_search(graph, queue, end))
}

func max(list []int) int {
	max := 0
	for _, i := range list {
		if i > max {
			max = i
		}
	}

	return max
}

func (g *Graph) Condense() {
	// end := g.findEnd()
	// current := g.findStart()

	// for current != end {
	// 	g.GetDestinations()
	// }
}

func A(input string) int {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	now := time.Now()
	locations := parse(f)
	fmt.Printf("Parsed locations in %v\n", time.Since(now))

	now = time.Now()
	graph := buildGraph(locations, false)
	fmt.Printf("Built graph in %v\n", time.Since(now))

	now = time.Now()
	graph.Condense()
	fmt.Printf("Condensed graph in %v\n", time.Since(now))

	now = time.Now()
	longest := longest(graph)
	fmt.Printf("Found longest in %v\n", time.Since(now))

	return longest
}

func buildGraph(locations [][]Location, part2 bool) Graph {
	graph := Graph{
		height: len(locations),
		width:  len(locations[0]),
	}

	for row := 0; row < len(locations); row++ {
		for col := 0; col < len(locations[row]); col++ {
			c := locations[row][col]
			node := Node{
				ID:   nodeID(row, col),
				Row:  row,
				Col:  col,
				Type: c.Type,
			}

			graph.AddNode(node)
			for _, c := range connections(locations, row, col, part2) {
				graph.AddDestination(node.ID, c)
			}
		}
	}

	return graph
}

func nodeID(row, col int) string {
	return fmt.Sprintf("row=%d;col=%d", row, col)
}

func connections(locations [][]Location, row, col int, part2 bool) []Edge {
	conns := []Edge{}
	if row > 0 {
		up := locations[row-1][col]
		if !part2 && (up.Type == "." || up.Type == "^") {
			conns = append(conns, Edge{Destination: nodeID(up.Row, up.Col), Direction: "up"})
		} else if part2 && up.Type != "#" {
			conns = append(conns, Edge{Destination: nodeID(up.Row, up.Col), Direction: "up"})
		}
	}

	if row < len(locations)-1 {
		down := locations[row+1][col]
		if !part2 && (down.Type == "." || down.Type == "v") {
			conns = append(conns, Edge{Destination: nodeID(down.Row, down.Col), Direction: "down"})
		} else if part2 && down.Type != "#" {
			conns = append(conns, Edge{Destination: nodeID(down.Row, down.Col), Direction: "down"})
		}
	}

	if col < len(locations[row])-1 {
		right := locations[row][col+1]
		if !part2 && (right.Type == "." || right.Type == ">") {
			conns = append(conns, Edge{Destination: nodeID(right.Row, right.Col), Direction: "right"})
		} else if part2 && right.Type != "#" {
			conns = append(conns, Edge{Destination: nodeID(right.Row, right.Col), Direction: "right"})
		}
	}

	if col > 0 {
		left := locations[row][col-1]
		if !part2 && (left.Type == "." || left.Type == "<") {
			conns = append(conns, Edge{Destination: nodeID(left.Row, left.Col), Direction: "left"})
		} else if part2 && left.Type != "#" {
			conns = append(conns, Edge{Destination: nodeID(left.Row, left.Col), Direction: "left"})
		}
	}

	return conns
}

func B(input string) int {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	now := time.Now()
	locations := parse(f)
	fmt.Printf("Parsed locations in %v\n", time.Since(now))

	now = time.Now()
	graph := buildGraph(locations, true)
	fmt.Printf("Built graph in %v\n", time.Since(now))

	now = time.Now()
	graph.Condense()
	fmt.Printf("Condensed graph in %v\n", time.Since(now))

	now = time.Now()
	longest := longest(graph)
	fmt.Printf("Found longest in %v\n", time.Since(now))
	return longest
}

func parse(f *os.File) [][]Location {
	locations := [][]Location{}
	scanner := bufio.NewScanner(f)
	row := 0
	for scanner.Scan() {
		r := []Location{}
		for col, t := range strings.Split(scanner.Text(), "") {
			r = append(r, Location{Row: row, Col: col, Type: t})
		}
		locations = append(locations, r)
		row++
	}

	return locations
}
