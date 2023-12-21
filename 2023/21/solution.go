package pkg202321

import (
	"bufio"
	"os"
	"strings"
)

type Grid [][]*Point

type Set map[Point]bool

func (s *Set) Push(p Point) {
	(*s)[p] = true
}

func (s *Set) Values() []Point {
	vs := []Point{}
	for v, _ := range *s {
		vs = append(vs, v)
	}

	return vs
}

func (g *Grid) Walk(list Set, times int) []Point {
	for times > 0 {
		newList := Set{}
		for _, p := range list.Values() {
			if p.Row-1 > 0 {
				up := (*g)[p.Row-1][p.Col]
				if up.Type != "#" {
					newList.Push(*up)
				}
			}

			if p.Row+1 < len(*g) {
				down := (*g)[p.Row+1][p.Col]
				if down.Type != "#" {
					newList.Push(*down)
				}
			}

			if p.Col-1 > 0 {
				left := (*g)[p.Row][p.Col-1]
				if left.Type != "#" {
					newList.Push(*left)
				}
			}

			if p.Col+1 < len((*g)[p.Row]) {
				right := (*g)[p.Row][p.Col+1]
				if right.Type != "#" {
					newList.Push(*right)
				}
			}
		}

		list = newList
		times--
	}

	return list.Values()
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

func (g *Grid) Start(steps int) int {
	p := g.FindStartingPosition()
	s := Set{}
	s.Push(*p)
	pos := g.Walk(s, steps)
	return len(pos)
}

type Point struct {
	Row  int
	Col  int
	Type string
}

func A(input string, steps int) int {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	grid := parse(f)
	return grid.Start(steps)
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
