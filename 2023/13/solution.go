package pkg202313

import (
	"fmt"
	"os"
	"strings"
)

type Direction int

const (
	HORIZONTAL Direction = iota
	VERTICAL
	UNKNOWN
)

func (d *Direction) String() string {
	switch *d {
	case HORIZONTAL:
		return "horizontal"
	case VERTICAL:
		return "vertical"
	default:
		return "UNKNOWN"
	}
}

type Pattern struct {
	Rows []string
}

func (p *Pattern) ColAsString(colIndex int) string {
	s := strings.Builder{}
	for i := 0; i < len(p.Rows); i++ {
		s.WriteByte(p.Rows[i][colIndex])
	}

	return s.String()
}

func (p *Pattern) IsVerticalReflection(index int) bool {
	for i := 1; index-i >= 0 && index+i+1 < len(p.Rows[0]); i++ {
		if p.ColAsString(index-i) != p.ColAsString(index+i+1) {
			return false
		}
	}

	return true
}

func (p *Pattern) IsHorizontalReflection(index int) bool {
	for i := 1; index-i >= 0 && index+i+1 < len(p.Rows); i++ {
		if p.Rows[index-i] != p.Rows[index+i+1] {
			return false
		}
	}

	return true
}

func (p *Pattern) FindReflection() (int, Direction) {
	for i := 0; i < len(p.Rows)-1; i++ {
		if p.Rows[i] == p.Rows[i+1] && p.IsHorizontalReflection(i) {
			return i + 1, HORIZONTAL
		}
	}

	for i := 0; i < len(p.Rows[0])-1; i++ {
		if p.ColAsString(i) == p.ColAsString(i+1) && p.IsVerticalReflection(i) {
			return i + 1, VERTICAL
		}
	}

	return 0, UNKNOWN
}

func A(input string) int {
	sum := 0
	for _, p := range parsePatterns(input) {
		index, direction := p.FindReflection()
		switch direction {
		case HORIZONTAL:
			fmt.Printf("H %d\n", index)
			sum += 100 * index
		case VERTICAL:
			fmt.Printf("V %d\n", index)
			sum += index
		default:
			panic("should not happen")
		}
	}
	return sum
}

func parsePatterns(input string) []Pattern {
	data, err := os.ReadFile(input)
	if err != nil {
		panic("bad file")
	}

	patterns := []Pattern{}
	for _, pattern := range strings.Split(string(data), "\n\n") {
		patterns = append(patterns, Pattern{Rows: strings.Split(pattern, "\n")})
	}

	return patterns
}
