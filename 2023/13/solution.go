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

func (p *Pattern) IsVerticalReflection(index int, considerSmudge bool) bool {
	curSmudges := 0
	for i := 1; index-i >= 0 && index+i+1 < len(p.Rows[0]); i++ {
		if !considerSmudge || curSmudges > 0 {
			if p.ColAsString(index-i) != p.ColAsString(index+i+1) {
				return false
			}
		} else {
			diff := p.Diff(p.ColAsString(index-i), p.ColAsString(index+i+1))
			if diff > 1 {
				return false
			}

			if diff == 1 {
				curSmudges += 1
			}
		}
	}

	if considerSmudge {
		return curSmudges == 1
	}

	return true
}

func (p *Pattern) IsHorizontalReflection(index int, considerSmudge bool) bool {
	curSmudges := 0
	for i := 1; index-i >= 0 && index+i+1 < len(p.Rows); i++ {
		if !considerSmudge || curSmudges > 0 {
			if p.Rows[index-i] != p.Rows[index+i+1] {
				return false
			}
		} else {
			diff := p.Diff(p.Rows[index-i], p.Rows[index+i+1])
			if diff > 1 {
				return false
			}

			if diff == 1 {
				curSmudges += 1
			}
		}
	}

	if considerSmudge {
		return curSmudges == 1
	}

	return true
}

func (p *Pattern) Diff(a, b string) int {
	diff := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff += 1
		}
	}

	return diff
}

func (p *Pattern) FindReflection(part2 bool) (int, Direction) {
	for i := 0; i < len(p.Rows)-1; i++ {
		if !part2 {
			if p.Rows[i] == p.Rows[i+1] && p.IsHorizontalReflection(i, false) {
				return i + 1, HORIZONTAL
			}
		} else {
			diff := p.Diff(p.Rows[i], p.Rows[i+1])
			if diff <= 1 && p.IsHorizontalReflection(i, diff == 0) {
				return i + 1, HORIZONTAL
			}
		}
	}

	for i := 0; i < len(p.Rows[0])-1; i++ {
		if !part2 {
			if p.ColAsString(i) == p.ColAsString(i+1) && p.IsVerticalReflection(i, false) {
				return i + 1, VERTICAL
			}
		} else {
			diff := p.Diff(p.ColAsString(i), p.ColAsString(i+1))
			if diff <= 1 && p.IsVerticalReflection(i, diff == 0) {
				return i + 1, VERTICAL
			}
		}
	}

	return 0, UNKNOWN
}

func A(input string) int {
	sum := 0
	for _, p := range parsePatterns(input) {
		index, direction := p.FindReflection(false)
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

func B(input string) int {
	sum := 0
	for _, p := range parsePatterns(input) {
		index, direction := p.FindReflection(true)
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
