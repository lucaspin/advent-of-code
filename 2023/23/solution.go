package pkg202323

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

type Location struct {
	Row  int
	Col  int
	Type string
}

func findStart(locations [][]Location) Point {
	for _, c := range locations[0] {
		if c.Type == "." {
			return Point{Row: c.Row, Col: c.Col}
		}
	}

	panic("asdasds")
}

func findEnd(locations [][]Location) Point {
	for _, c := range locations[len(locations)-1] {
		if c.Type == "." {
			return Point{Row: c.Row, Col: c.Col}
		}
	}

	panic("asdasds")
}

func A(input string) int {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	locations := parse(f)
	graph := buildGraph(locations, false)
	return graph.Longest(findStart(locations), findEnd(locations))
}

func B(input string) int {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	locations := parse(f)
	graph := buildGraph(locations, true)
	graph.Simplify()
	return graph.Longest(findStart(locations), findEnd(locations))
}

type Point struct {
	Row int
	Col int
}

type Dest struct {
	p Point
	d int
}

type Graph map[Point][]Dest

func (g *Graph) Add(p Point, dest []Dest) {
	(*g)[p] = dest
}

func (g *Graph) AddDest(p Point, dest Dest) {
	(*g)[p] = append((*g)[p], dest)
}

func (g *Graph) Simplify() {
	points := make([]Point, 0)
	for k := range *g {
		points = append(points, k)
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i].Row < points[j].Row
	})

	for _, k := range points {
		v := (*g)[k]
		if len(v) == 2 {
			a := v[0]
			b := v[1]
			g.remove(a.p, k)
			g.remove(b.p, k)
			g.AddDest(a.p, Dest{p: b.p, d: a.d + b.d})
			g.AddDest(b.p, Dest{p: a.p, d: a.d + b.d})
			delete(*g, k)
		}
	}
}

func (g *Graph) remove(a, b Point) {
	new := []Dest{}
	for _, d := range (*g)[a] {
		if d.p != b {
			new = append(new, d)
		}
	}

	(*g)[a] = new
}

func buildGraph(locations [][]Location, part2 bool) *Graph {
	graph := Graph{}

	for row := 0; row < len(locations); row++ {
		for col := 0; col < len(locations[row]); col++ {
			c := locations[row][col]
			if c.Type != "#" {
				graph.Add(Point{Row: row, Col: col}, connections(locations, row, col, part2))
			}
		}
	}

	return &graph
}

type stack struct {
	items []stackItem
}

func (q *stack) Push(item stackItem) {
	q.items = append(q.items, item)
}

func (q *stack) Pop() stackItem {
	head := q.items[len(q.items)-1]
	q.items = q.items[0 : len(q.items)-1]
	return head
}

func (q *stack) Len() int {
	return len(q.items)
}

type stackItem struct {
	p Point
	d int
}

type PointSet map[Point]bool

func (n *PointSet) Has(v Point) bool {
	_, ok := (*n)[v]
	return ok
}

func (n *PointSet) Remove(v Point) {
	delete(*n, v)
}

func (n *PointSet) Add(v Point) {
	(*n)[v] = true
}

func (n *PointSet) Len() int {
	return len(*n)
}

func (n *PointSet) Values() []Point {
	ls := []Point{}
	for k := range *n {
		ls = append(ls, k)
	}

	return ls
}

func (n *PointSet) Copy() PointSet {
	new := make(map[Point]bool, len(*n))
	for k, v := range *n {
		new[k] = v
	}

	return new
}

func (g *Graph) Search(stack stack, target Point) int {
	seen := PointSet{}
	maximum := 0
	for stack.Len() > 0 {
		i := stack.Pop()

		if i.d == -1 {
			seen.Remove(i.p)
			continue
		}

		if i.p == target {
			maximum = max(maximum, i.d)
			continue
		}

		stack.Push(stackItem{p: i.p, d: -1})
		seen.Add(i.p)

		for _, d := range (*g)[i.p] {
			if seen.Has(d.p) {
				continue
			}

			seen.Add(i.p)
			stack.Push(stackItem{
				p: d.p,
				d: i.d + d.d,
			})
		}
	}

	return maximum
}

func (g *Graph) Longest(start, end Point) int {
	queue := stack{}
	queue.Push(stackItem{
		p: start,
		d: 0,
	})

	return g.Search(queue, end)
}

func connections(locations [][]Location, row, col int, part2 bool) []Dest {
	conns := []Dest{}
	if row > 0 {
		up := locations[row-1][col]
		if !part2 && (up.Type == "." || up.Type == "^") {
			conns = append(conns, Dest{p: Point{Row: up.Row, Col: up.Col}, d: 1})
		} else if part2 && up.Type != "#" {
			conns = append(conns, Dest{p: Point{Row: up.Row, Col: up.Col}, d: 1})
		}
	}

	if row < len(locations)-1 {
		down := locations[row+1][col]
		if !part2 && (down.Type == "." || down.Type == "v") {
			conns = append(conns, Dest{p: Point{Row: down.Row, Col: down.Col}, d: 1})
		} else if part2 && down.Type != "#" {
			conns = append(conns, Dest{p: Point{Row: down.Row, Col: down.Col}, d: 1})
		}
	}

	if col < len(locations[row])-1 {
		right := locations[row][col+1]
		if !part2 && (right.Type == "." || right.Type == ">") {
			conns = append(conns, Dest{p: Point{Row: right.Row, Col: right.Col}, d: 1})
		} else if part2 && right.Type != "#" {
			conns = append(conns, Dest{p: Point{Row: right.Row, Col: right.Col}, d: 1})
		}
	}

	if col > 0 {
		left := locations[row][col-1]
		if !part2 && (left.Type == "." || left.Type == "<") {
			conns = append(conns, Dest{p: Point{Row: left.Row, Col: left.Col}, d: 1})
		} else if part2 && left.Type != "#" {
			conns = append(conns, Dest{p: Point{Row: left.Row, Col: left.Col}, d: 1})
		}
	}

	return conns
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
