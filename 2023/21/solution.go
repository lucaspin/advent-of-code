package pkg202321

import (
	"bufio"
	"os"
	"strings"
)

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
	Point    Point
	Distance int64
}

type Grid [][]*Point

func (g *Grid) Walk(q queue) map[Point]int64 {
	distances := map[Point]int64{}

	for q.Len() > 0 {
		i := q.Pop()

		if _, ok := distances[i.Point]; ok {
			continue
		}

		distances[i.Point] = i.Distance

		if i.Point.Row > 0 {
			up := (*g)[i.Point.Row-1][i.Point.Col]
			if up.Type != "#" {
				q.Push(queueItem{Point: *up, Distance: i.Distance + 1})
			}
		}

		if i.Point.Row < len(*g)-1 {
			down := (*g)[i.Point.Row+1][i.Point.Col]
			if down.Type != "#" {
				q.Push(queueItem{Point: *down, Distance: i.Distance + 1})
			}
		}

		if i.Point.Col > 0 {
			left := (*g)[i.Point.Row][i.Point.Col-1]
			if left.Type != "#" {
				q.Push(queueItem{Point: *left, Distance: i.Distance + 1})
			}
		}

		if i.Point.Col < len((*g)[i.Point.Row])-1 {
			right := (*g)[i.Point.Row][i.Point.Col+1]
			if right.Type != "#" {
				q.Push(queueItem{Point: *right, Distance: i.Distance + 1})
			}
		}
	}

	return distances
}

func (g *Grid) FindStartingPosition() *Point {
	for _, row := range *g {
		for _, col := range row {
			if col.Type == "S" {
				return col
			}
		}
	}

	panic("no starting position")
}

func (g *Grid) Start() map[Point]int64 {
	p := g.FindStartingPosition()
	q := queue{}
	q.Push(queueItem{Point: *p, Distance: 0})
	return g.Walk(q)
}

type Point struct {
	Row  int
	Col  int
	Type string
}

func countCond(m map[Point]int64, cond func(k Point, v int64) bool) int64 {
	c := int64(0)
	for k, v := range m {
		if cond(k, v) {
			c++
		}
	}

	return c
}

func A(input string, steps int) int64 {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	grid := parse(f)
	distances := grid.Start()
	return countCond(distances, func(k Point, v int64) bool {
		return v <= 64 && v%2 == 0
	})
}

func B(input string, steps int) int64 {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	grid := parse(f)
	distances := grid.Start()

	gridDimension := len(grid)
	n := int64((steps - (gridDimension / 2)) / gridDimension)

	odd := countCond(distances, func(k Point, v int64) bool { return v%2 == 1 })
	even := countCond(distances, func(k Point, v int64) bool { return v%2 == 0 })
	cornerOdd := countCond(distances, func(k Point, v int64) bool { return v%2 == 1 && v > 65 })
	cornerEven := countCond(distances, func(k Point, v int64) bool { return v%2 == 0 && v > 65 })

	return ((n + 1) * (n + 1) * odd) + (n * n * even) - ((n + 1) * cornerOdd) + (n * cornerEven)
}

func parse(f *os.File) Grid {
	scanner := bufio.NewScanner(f)

	grid := Grid{}

	r := 0
	for scanner.Scan() {
		row := []*Point{}
		for c, v := range strings.Split(scanner.Text(), "") {
			row = append(row, &Point{Row: r, Col: c, Type: v})
		}

		grid = append(grid, row)
		r++
	}

	return grid
}
