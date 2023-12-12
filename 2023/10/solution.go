package pkg202310

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TileType int

const (
	VERTICAL TileType = iota
	HORIZONTAL
	NORTH_EAST
	NORTH_WEST
	SOUTH_WEST
	SOUTH_EAST
	GROUND
	START
)

func (t *TileType) In() []Direction {
	switch *t {
	case VERTICAL:
		return []Direction{UP, DOWN}
	case HORIZONTAL:
		return []Direction{RIGHT, LEFT}
	case NORTH_EAST:
		return []Direction{DOWN, LEFT}
	case NORTH_WEST:
		return []Direction{DOWN, RIGHT}
	case SOUTH_EAST:
		return []Direction{UP, LEFT}
	case SOUTH_WEST:
		return []Direction{UP, RIGHT}
	default:
		return []Direction{}
	}
}

func (t *TileType) Out() []Direction {
	switch *t {
	case VERTICAL:
		return []Direction{UP, DOWN}
	case HORIZONTAL:
		return []Direction{RIGHT, LEFT}
	case NORTH_EAST:
		return []Direction{UP, RIGHT}
	case NORTH_WEST:
		return []Direction{UP, LEFT}
	case SOUTH_EAST:
		return []Direction{DOWN, RIGHT}
	case SOUTH_WEST:
		return []Direction{DOWN, LEFT}
	default:
		return []Direction{}
	}
}

func (t *TileType) String() string {
	switch *t {
	case VERTICAL:
		return "VERTICAL"
	case HORIZONTAL:
		return "HORIZONTAL"
	case NORTH_EAST:
		return "NORTH_EAST"
	case NORTH_WEST:
		return "NORTH_WEST"
	case SOUTH_WEST:
		return "SOUTH_WEST"
	case SOUTH_EAST:
		return "SOUTH_EAST"
	case START:
		return "START"
	default:
		return "GROUND"
	}
}

type Direction int

const (
	UP Direction = iota
	DOWN
	RIGHT
	LEFT
)

type Tile struct {
	Type TileType
	Row  int
	Col  int
}

func (t *Tile) Equals(o *Tile) bool {
	if o == nil {
		return false
	}

	return t.Row == o.Row && t.Col == o.Col
}

func (t *Tile) String() string {
	return fmt.Sprintf("(%d,%d,%s)", t.Row, t.Col, t.Type.String())
}

func (t *Tile) Copy() *Tile {
	return &Tile{
		Row:  t.Row,
		Col:  t.Col,
		Type: t.Type,
	}
}

type Grid struct {
	Tiles [][]Tile
}

func contains(ds []Direction, d Direction) bool {
	for _, x := range ds {
		if x == d {
			return true
		}
	}

	return false
}

func (g *Grid) NextTile(current *Tile, previous *Tile) Tile {
	var top, bottom, right, left *Tile

	if current.Row > 0 {
		top = &g.Tiles[current.Row-1][current.Col]
	}

	if current.Row < len(g.Tiles)-1 {
		bottom = &g.Tiles[current.Row+1][current.Col]
	}

	if current.Col > 0 {
		left = &g.Tiles[current.Row][current.Col-1]
	}

	if current.Col < len(g.Tiles[current.Row])-1 {
		right = &g.Tiles[current.Row][current.Col+1]
	}

	for _, out := range current.Type.Out() {
		switch out {
		case UP:
			if top != nil && contains(top.Type.In(), out) && !top.Equals(previous) {
				return *top
			}
		case DOWN:
			if bottom != nil && contains(bottom.Type.In(), out) && !bottom.Equals(previous) {
				return *bottom
			}
		case LEFT:
			if left != nil && contains(left.Type.In(), out) && !left.Equals(previous) {
				return *left
			}
		case RIGHT:
			if right != nil && contains(right.Type.In(), out) && !right.Equals(previous) {
				return *right
			}
		}
	}

	panic("THIS SHOULD NOT HAPPEN")
}

func (g *Grid) FindStartingTile() *Tile {
	for _, row := range g.Tiles {
		for _, tile := range row {
			if tile.Type == START {
				return &tile
			}
		}
	}

	return nil
}

func (g *Grid) DetermineTypeThroughNeighbours(tile Tile) TileType {
	var top, bottom, right, left *Tile

	if tile.Row > 0 {
		top = &g.Tiles[tile.Row-1][tile.Col]
	}

	if tile.Row < len(g.Tiles)-1 {
		bottom = &g.Tiles[tile.Row+1][tile.Col]
	}

	if tile.Col > 0 {
		left = &g.Tiles[tile.Row][tile.Col-1]
	}

	if tile.Col < len(g.Tiles[tile.Row])-1 {
		right = &g.Tiles[tile.Row][tile.Col+1]
	}

	if top != nil && bottom != nil && contains(top.Type.In(), UP) && contains(bottom.Type.In(), DOWN) {
		return VERTICAL
	}

	if right != nil && left != nil && contains(right.Type.In(), RIGHT) && contains(left.Type.In(), LEFT) {
		return HORIZONTAL
	}

	if top != nil && right != nil && contains(top.Type.In(), UP) && contains(right.Type.In(), RIGHT) {
		return NORTH_EAST
	}

	if top != nil && left != nil && contains(top.Type.In(), UP) && contains(left.Type.In(), LEFT) {
		return NORTH_WEST
	}

	if bottom != nil && left != nil && contains(bottom.Type.In(), DOWN) && contains(left.Type.In(), LEFT) {
		return SOUTH_WEST
	}

	if bottom != nil && right != nil && contains(bottom.Type.In(), DOWN) && contains(right.Type.In(), RIGHT) {
		return SOUTH_EAST
	}

	panic("this should not happen")
}

func A(input string) int {
	f, _ := os.Open(input)
	grid := parseGrid(f)

	steps := 0
	startingTile := grid.FindStartingTile()
	currentTile := Tile{
		Row:  startingTile.Row,
		Col:  startingTile.Col,
		Type: grid.DetermineTypeThroughNeighbours(*startingTile),
	}

	grid.Tiles[startingTile.Row][startingTile.Col] = *currentTile.Copy()

	fmt.Printf("Starting tile is %s\n", currentTile.String())

	var previous *Tile

	for steps == 0 || currentTile.Row != startingTile.Row || currentTile.Col != startingTile.Col {
		next := grid.NextTile(&currentTile, previous)
		previous = currentTile.Copy()
		currentTile = next
		steps += 1
	}

	return steps / 2
}

func B(input string) int {
	f, _ := os.Open(input)
	grid := parseGrid(f)
	startingTile := grid.FindStartingTile()
	currentTile := Tile{
		Row:  startingTile.Row,
		Col:  startingTile.Col,
		Type: grid.DetermineTypeThroughNeighbours(*startingTile),
	}

	grid.Tiles[startingTile.Row][startingTile.Col] = *currentTile.Copy()

	var loop = []Tile{}
	var previous *Tile

	for len(loop) == 0 || currentTile.Row != startingTile.Row || currentTile.Col != startingTile.Col {
		loop = append(loop, currentTile)
		next := grid.NextTile(&currentTile, previous)
		previous = currentTile.Copy()
		currentTile = next
	}

	inside := []Tile{}
	for _, row := range grid.Tiles {
		for _, tile := range row {

			// we don't consider tiles that are part of the loop
			if in(loop, tile) != nil {
				continue
			}

			// Find vertical intersections
			intersecs := 0
			for r := tile.Row; r < len(grid.Tiles); r++ {
				if x := in(loop, grid.Tiles[r][tile.Col]); x != nil && (x.Type == HORIZONTAL || x.Type == NORTH_WEST || x.Type == SOUTH_WEST) {
					intersecs += 1
				}
			}

			// a point is inside the shape defined by the loop,
			// if we have a number of odd intersections
			if intersecs%2 == 1 {
				inside = append(inside, tile)
			}
		}
	}

	for _, i := range inside {
		fmt.Printf("%s\n", i.String())
	}

	return len(inside)
}

func in(loop []Tile, tile Tile) *Tile {
	for _, t := range loop {
		if t.Row == tile.Row && t.Col == tile.Col {
			return &t
		}
	}

	return nil
}

func parseGrid(f *os.File) Grid {
	grid := Grid{
		Tiles: [][]Tile{},
	}

	rowIndex := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := parseRow(rowIndex, scanner.Text())
		grid.Tiles = append(grid.Tiles, row)
		rowIndex += 1
	}

	return grid
}

func parseRow(rowIndex int, line string) []Tile {
	row := []Tile{}
	for colIndex, c := range strings.Split(line, "") {
		row = append(row, Tile{
			Row:  rowIndex,
			Col:  colIndex,
			Type: parseTileType(c),
		})
	}

	return row
}

func parseTileType(c string) TileType {
	switch c {
	case "|":
		return VERTICAL
	case "-":
		return HORIZONTAL
	case "F":
		return SOUTH_EAST
	case "J":
		return NORTH_WEST
	case "L":
		return NORTH_EAST
	case "7":
		return SOUTH_WEST
	case "S":
		return START
	default:
		return GROUND
	}
}
