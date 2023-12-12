package pkg202311

import (
	"bufio"
	"math"
	"math/rand"
	"os"
	"strings"
)

type Point struct {
	Row    int
	Col    int
	Galaxy bool
	ID     int
}

type Grid struct {
	Points [][]Point
}

func (g *Grid) EmptyRow(rowIndex int) bool {
	for _, p := range g.Points[rowIndex] {
		if p.Galaxy {
			return false
		}
	}

	return true
}

func (g *Grid) EmptyCol(colIndex int) bool {
	for i := 0; i < len(g.Points); i++ {
		if g.Points[i][colIndex].Galaxy {
			return false
		}
	}

	return true
}

func (g *Grid) findGalaxies() []Point {
	ps := []Point{}
	for _, row := range g.Points {
		for _, p := range row {
			if p.Galaxy {
				ps = append(ps, p)
			}
		}
	}

	return ps
}

func pairs(points []Point) [][]Point {
	in := func(pairs [][]Point, p, q Point) bool {
		for _, pair := range pairs {
			if pair[0].ID == p.ID && pair[1].ID == q.ID {
				return true
			}

			if pair[0].ID == q.ID && pair[1].ID == p.ID {
				return true
			}
		}

		return false
	}

	pairs := [][]Point{}
	for _, p := range points {
		for _, q := range points {
			if p.Row == q.Row && p.Col == q.Col {
				continue
			}

			if !in(pairs, p, q) {
				pairs = append(pairs, []Point{p, q})
			}
		}
	}

	return pairs
}

func (g *Grid) emptyRowsBetween(a, b Point) int {
	if a.Row > b.Row {
		a, b = b, a
	}

	empty := 0
	for i := a.Row; i <= b.Row; i++ {
		if g.EmptyRow(i) {
			empty += 1
		}
	}

	return empty
}

func (g *Grid) emptyColsBetween(a, b Point) int {
	if a.Col > b.Col {
		a, b = b, a
	}

	empty := 0
	for i := a.Col; i <= b.Col; i++ {
		if g.EmptyCol(i) {
			empty += 1
		}
	}

	return empty
}

func expandDistance(distance, empty, extra int) int {
	if empty == 0 {
		return distance
	}

	return distance + (empty * extra)
}

func Solution(input string, timesLarger int) int {
	f, _ := os.Open(input)
	grid := parseGrid(f)

	galaxies := grid.findGalaxies()

	result := 0
	for _, pair := range pairs(galaxies) {
		xd := xdistance(pair[0], pair[1])
		yd := ydistance(pair[0], pair[1])
		emptyRows := grid.emptyRowsBetween(pair[0], pair[1])
		emptyCols := grid.emptyColsBetween(pair[0], pair[1])
		d := expandDistance(xd, emptyRows, timesLarger-1) + expandDistance(yd, emptyCols, timesLarger-1)
		result += d
	}

	return result
}

func xdistance(a, b Point) int {
	return int(math.Abs(float64(a.Row - b.Row)))
}

func ydistance(a, b Point) int {
	return int(math.Abs(float64(a.Col - b.Col)))
}

func distance(a, b Point) int {
	return int(math.Abs(float64(a.Row-b.Row)) + math.Abs(float64(a.Col-b.Col)))
}

func emptyCol(grid [][]Point, col int) bool {
	for i := 0; i < len(grid); i++ {
		if grid[i][col].Galaxy {
			return false
		}
	}

	return true
}

func emptyRow(grid [][]Point, row int) bool {
	for _, col := range grid[row] {
		if col.Galaxy {
			return false
		}
	}

	return true
}

func parseGrid(f *os.File) Grid {
	grid := Grid{Points: [][]Point{}}

	row := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		r := []Point{}
		for col, c := range strings.Split(scanner.Text(), "") {
			p := Point{Row: row, Col: col, Galaxy: isGalaxy(c)}
			if p.Galaxy {
				p.ID = rand.Int()
			}

			r = append(r, p)
		}

		grid.Points = append(grid.Points, r)
		row += 1
	}

	return grid
}

func isGalaxy(s string) bool {
	return s == "#"
}
