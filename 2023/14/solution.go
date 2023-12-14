package pkg202314

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

type Grid [][]Point

func (g *Grid) MoveUp(row, col int) {
	g1 := *g
	for i := row; i > 0; i-- {
		if g1[i-1][col].Value == "#" {
			break
		}

		if g1[i-1][col].Value == "." {
			g1[i-1][col].Value = "O"
			g1[i][col].Value = "."
		}
	}
}

func (g *Grid) MoveLeft(row, col int) {
	g1 := *g
	for i := col; i > 0; i-- {
		if g1[row][i-1].Value == "#" {
			break
		}

		if g1[row][i-1].Value == "." {
			g1[row][i-1].Value = "O"
			g1[row][i].Value = "."
		}
	}
}

func (g *Grid) MoveDown(row, col int) {
	g1 := *g
	for i := row; i < len(g1)-1; i++ {
		if g1[i+1][col].Value == "#" {
			break
		}

		if g1[i+1][col].Value == "." {
			g1[i+1][col].Value = "O"
			g1[i][col].Value = "."
		}
	}
}

func (g *Grid) MoveRight(row, col int) {
	g1 := *g
	for i := col; i < len(g1[row])-1; i++ {
		if g1[row][i+1].Value == "#" {
			break
		}

		if g1[row][i+1].Value == "." {
			g1[row][i+1].Value = "O"
			g1[row][i].Value = "."
		}
	}
}

func (g *Grid) Cycle() {
	g.Tilt(func(i, j int) { g.MoveUp(i, j) })             // north
	g.Tilt(func(i, j int) { g.MoveLeft(i, j) })           // west
	g.TiltBackwards(func(i, j int) { g.MoveDown(i, j) })  // south
	g.TiltBackwards(func(i, j int) { g.MoveRight(i, j) }) // east
}

func (g *Grid) Tilt(f func(i, j int)) {
	g1 := *g
	for i := 0; i < len(g1); i++ {
		for j := 0; j < len(g1[i]); j++ {
			if g1[i][j].Value == "." || g1[i][j].Value == "#" {
				continue
			}

			f(i, j)
		}
	}
}

func (g *Grid) TiltBackwards(f func(i, j int)) {
	g1 := *g
	for i := len(g1) - 1; i >= 0; i-- {
		for j := len(g1[i]) - 1; j >= 0; j-- {
			if g1[i][j].Value == "." || g1[i][j].Value == "#" {
				continue
			}

			f(i, j)
		}
	}
}

func (g *Grid) Hash() string {
	hash := md5.Sum([]byte(g.Print()))
	return hex.EncodeToString(hash[:])
}

func (g *Grid) Print() string {
	s := strings.Builder{}
	for _, row := range *g {
		for _, col := range row {
			s.WriteString(col.Value)
		}

		s.WriteString("\n")
	}

	return s.String()
}

func (g *Grid) TotalLoad() int {
	total := 0

	g1 := *g
	for rowIndex, row := range g1 {
		for _, col := range row {
			if col.Value == "O" {
				total += len(g1) - rowIndex
			}
		}
	}

	return total
}

type Point struct {
	Row   int
	Col   int
	Value string
}

func A(input string) int {
	grid := parseGrid(input)
	grid.Tilt(func(i, j int) { grid.MoveUp(i, j) })
	return grid.TotalLoad()
}

func get(elems []Cache, hash string) *Cache {
	for _, el := range elems {
		if el.Hash == hash {
			return &el
		}
	}

	return nil
}

type Cache struct {
	Hash      string
	Iteration int64
}

func B(input string) int {
	grid := parseGrid(input)
	configs := []Cache{}
	remaining := int64(0)

	for i := int64(0); i < 1000000000; i++ {
		if e := get(configs, grid.Hash()); e != nil {
			fmt.Printf("Found cycle at %d, repeating config previously found at %d\n", i, e.Iteration)
			remaining = (1000000000 - e.Iteration) % (i - e.Iteration)
			break
		} else {
			configs = append(configs, Cache{Hash: grid.Hash(), Iteration: i})
			grid.Cycle()
		}
	}

	fmt.Printf("Remaining: %d\n", remaining)
	for j := 0; j < int(remaining); j++ {
		grid.Cycle()
	}

	return grid.TotalLoad()
}

func parseGrid(input string) Grid {
	data, err := os.ReadFile(input)
	if err != nil {
		panic("error reading file")
	}

	grid := [][]Point{}

	for rowIndex, row := range strings.Split(string(data), "\n") {
		r := []Point{}
		for colIndex, v := range strings.Split(row, "") {
			r = append(r, Point{Row: rowIndex, Col: colIndex, Value: v})
		}

		grid = append(grid, r)
	}

	return grid
}
