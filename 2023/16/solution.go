package pkg202316

import (
	"bufio"
	"os"
)

type Direction int

const (
	RIGHT Direction = iota
	LEFT
	UP
	DOWN
)

func (d *Direction) String() string {
	switch *d {
	case UP:
		return "UP"
	case DOWN:
		return "DOWN"
	case LEFT:
		return "LEFT"
	case RIGHT:
		return "RIGHT"
	}

	return "???"
}

type Beam struct {
	Row       int
	Col       int
	Direction Direction
}

func (b *Beam) Move(grid *Grid) []Beam {
	if b.Row < 0 || b.Col < 0 {
		return []Beam{}
	}

	if b.Row > len(grid.Points)-1 || b.Col > len(grid.Points[b.Row])-1 {
		return []Beam{}
	}

	gridPoint := grid.Points[b.Row][b.Col]
	gridPoint.Energized = true

	switch b.Direction {
	case RIGHT:
		if (gridPoint.Type == '.' || gridPoint.Type == '-') && b.Col < len(grid.Points[b.Row])-1 {
			return []Beam{
				{
					Row:       b.Row,
					Col:       b.Col + 1,
					Direction: RIGHT,
				},
			}
		}

		if gridPoint.Type == '/' && b.Row > 0 {
			return []Beam{
				{
					Row:       b.Row - 1,
					Col:       b.Col,
					Direction: UP,
				},
			}
		}

		if gridPoint.Type == '\\' && b.Row < len(grid.Points)-1 {
			return []Beam{
				{
					Row:       b.Row + 1,
					Col:       b.Col,
					Direction: DOWN,
				},
			}
		}

		if gridPoint.Type == '|' {
			bs := []Beam{}

			if b.Row > 0 {
				bs = append(bs, Beam{
					Row:       b.Row - 1,
					Col:       b.Col,
					Direction: UP,
				})
			}

			if b.Row < len(grid.Points) {
				bs = append(bs, Beam{
					Row:       b.Row + 1,
					Col:       b.Col,
					Direction: DOWN,
				})
			}

			return bs
		}
	case LEFT:
		if (gridPoint.Type == '.' || gridPoint.Type == '-') && b.Col > 0 {
			return []Beam{
				{
					Row:       b.Row,
					Col:       b.Col - 1,
					Direction: LEFT,
				},
			}
		}

		if gridPoint.Type == '/' && b.Row < len(grid.Points)-1 {
			return []Beam{
				{
					Row:       b.Row + 1,
					Col:       b.Col,
					Direction: DOWN,
				},
			}
		}

		if gridPoint.Type == '\\' && b.Row > 0 {
			return []Beam{
				{
					Row:       b.Row - 1,
					Col:       b.Col,
					Direction: UP,
				},
			}
		}

		if gridPoint.Type == '|' {
			bs := []Beam{}

			if b.Row > 0 {
				bs = append(bs, Beam{
					Row:       b.Row - 1,
					Col:       b.Col,
					Direction: UP,
				})
			}

			if b.Row < len(grid.Points) {
				bs = append(bs, Beam{
					Row:       b.Row + 1,
					Col:       b.Col,
					Direction: DOWN,
				})
			}

			return bs
		}
	case UP:
		if (gridPoint.Type == '.' || gridPoint.Type == '|') && b.Row > 0 {
			return []Beam{
				{
					Row:       b.Row - 1,
					Col:       b.Col,
					Direction: UP,
				},
			}
		}

		if gridPoint.Type == '/' && b.Col < len(grid.Points)-1 {
			return []Beam{
				{
					Row:       b.Row,
					Col:       b.Col + 1,
					Direction: RIGHT,
				},
			}
		}

		if gridPoint.Type == '\\' && b.Col > 0 {
			return []Beam{
				{
					Row:       b.Row,
					Col:       b.Col - 1,
					Direction: LEFT,
				},
			}
		}

		if gridPoint.Type == '-' {
			bs := []Beam{}

			if b.Col > 0 {
				bs = append(bs, Beam{
					Row:       b.Row,
					Col:       b.Col - 1,
					Direction: LEFT,
				})
			}

			if b.Row < len(grid.Points[b.Row]) {
				bs = append(bs, Beam{
					Row:       b.Row,
					Col:       b.Col + 1,
					Direction: RIGHT,
				})
			}

			return bs
		}
	case DOWN:
		if (gridPoint.Type == '.' || gridPoint.Type == '|') && b.Row < len(grid.Points)-1 {
			return []Beam{
				{
					Row:       b.Row + 1,
					Col:       b.Col,
					Direction: DOWN,
				},
			}
		}

		if gridPoint.Type == '/' && b.Col > 0 {
			return []Beam{
				{
					Row:       b.Row,
					Col:       b.Col - 1,
					Direction: LEFT,
				},
			}
		}

		if gridPoint.Type == '\\' && b.Col < len(grid.Points[b.Row])-1 {
			return []Beam{
				{
					Row:       b.Row,
					Col:       b.Col + 1,
					Direction: RIGHT,
				},
			}
		}

		if gridPoint.Type == '-' {
			bs := []Beam{}

			if b.Col > 0 {
				bs = append(bs, Beam{
					Row:       b.Row,
					Col:       b.Col - 1,
					Direction: LEFT,
				})
			}

			if b.Row < len(grid.Points[b.Row])-1 {
				bs = append(bs, Beam{
					Row:       b.Row,
					Col:       b.Col + 1,
					Direction: RIGHT,
				})
			}

			return bs
		}
	}

	return []Beam{}
}

type Point struct {
	Row       int
	Col       int
	Type      byte
	Energized bool
}

type Grid struct {
	Points [][]*Point
}

func in(beams []Beam, beam Beam) bool {
	for _, b := range beams {
		if b.Row == beam.Row && b.Col == beam.Col && beam.Direction == b.Direction {
			return true
		}
	}

	return false
}

func (g *Grid) FlashBeam(startingBeam Beam) {
	g.Reset()
	currentBeams := []Beam{startingBeam}
	seen := []Beam{}

	for len(currentBeams) > 0 {
		new := []Beam{}
		for _, beam := range currentBeams {
			for _, b := range beam.Move(g) {
				if !in(seen, b) {
					new = append(new, b)
				}
			}
		}

		currentBeams = new
		seen = append(seen, new...)
	}
}

func (g *Grid) Energized() int {
	n := 0
	for _, row := range g.Points {
		for _, col := range row {
			if col.Energized {
				n += 1
			}
		}
	}

	return n
}

func (g *Grid) Reset() {
	for _, row := range g.Points {
		for _, col := range row {
			col.Energized = false
		}
	}
}

func A(input string) int {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	grid := parseGrid(file)
	grid.FlashBeam(Beam{Row: 0, Col: 0, Direction: RIGHT})
	return grid.Energized()
}

func B(input string) int {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	grid := parseGrid(file)

	max := 0

	// first row
	for i := 0; i < len(grid.Points[0]); i++ {
		if i == 0 {
			grid.FlashBeam(Beam{Row: 0, Col: i, Direction: RIGHT})
			n := grid.Energized()
			if n > max {
				max = n
			}

			grid.FlashBeam(Beam{Row: 0, Col: i, Direction: DOWN})
			n = grid.Energized()
			if n > max {
				max = n
			}

			continue
		}

		if i == len(grid.Points[0])-1 {
			grid.FlashBeam(Beam{Row: 0, Col: i, Direction: LEFT})
			n := grid.Energized()
			if n > max {
				max = n
			}

			grid.FlashBeam(Beam{Row: 0, Col: i, Direction: DOWN})
			n = grid.Energized()
			if n > max {
				max = n
			}

			continue
		}

		grid.FlashBeam(Beam{Row: 0, Col: i, Direction: DOWN})
		n := grid.Energized()
		if n > max {
			max = n
		}
	}

	// last row
	lastRow := len(grid.Points) - 1
	for i := 0; i < len(grid.Points[lastRow]); i++ {
		if i == 0 {
			grid.FlashBeam(Beam{Row: lastRow, Col: i, Direction: RIGHT})
			n := grid.Energized()
			if n > max {
				max = n
			}

			grid.FlashBeam(Beam{Row: lastRow, Col: i, Direction: UP})
			n = grid.Energized()
			if n > max {
				max = n
			}

			continue
		}

		if i == len(grid.Points[lastRow])-1 {
			grid.FlashBeam(Beam{Row: 0, Col: i, Direction: LEFT})
			n := grid.Energized()
			if n > max {
				max = n
			}

			grid.FlashBeam(Beam{Row: 0, Col: i, Direction: UP})
			n = grid.Energized()
			if n > max {
				max = n
			}

			continue
		}

		grid.FlashBeam(Beam{Row: 0, Col: i, Direction: UP})
		n := grid.Energized()
		if n > max {
			max = n
		}
	}

	// first col
	for i := 1; i < len(grid.Points)-1; i++ {
		grid.FlashBeam(Beam{Row: i, Col: 0, Direction: RIGHT})
		n := grid.Energized()
		if n > max {
			max = n
		}
	}

	// last col
	lastCol := len(grid.Points[0]) - 1
	for i := 1; i < lastCol; i++ {
		grid.FlashBeam(Beam{Row: i, Col: lastCol, Direction: LEFT})
		n := grid.Energized()
		if n > max {
			max = n
		}
	}

	return max
}

func parseGrid(file *os.File) Grid {
	scanner := bufio.NewScanner(file)
	points := [][]*Point{}
	i := 0
	for scanner.Scan() {
		row := []*Point{}
		line := scanner.Text()
		for col := 0; col < len(line); col++ {
			row = append(row, &Point{Row: i, Col: col, Type: line[col]})
		}

		points = append(points, row)
		i++
	}

	return Grid{Points: points}
}
