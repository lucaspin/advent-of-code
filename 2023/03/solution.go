package pkg202303

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var lines = []string{}

func A(input string) int64 {
	lines = parse(input)
	result := int64(0)

	for i, line := range lines {
		n := []byte{}
		adjacent := false

		for col := 0; col < len(line); {
			v := line[col]

			//
			// Check if it is a digit.
			// If it is, accumulate it until we find a symbol to complete it.
			if isNumber(v) {
				n = append(n, v)

				// Check if this digit is adjacent to any symbol.
				if isAdjacentToSymbol(lines, line, i, col) {
					adjacent = true
				}

				// If this is the last digit on the line,
				// we check things here.
				if col == len(line)-1 {
					if len(n) > 0 && adjacent {
						v := byteArrToNumber(n)
						result += v
					}
				}

				col++
				continue
			}

			//
			// Here, we are at a symbol
			// Check if a number was found before this symbol.
			// If so, check if it is adjacent
			if len(n) > 0 && adjacent {
				v := byteArrToNumber(n)
				result += v
			}

			n = []byte{}
			adjacent = false
			col++
		}
	}

	return result
}

func B(input string) int64 {
	lines = parse(input)
	result := int64(0)

	for i, line := range lines {
		for col := 0; col < len(line); {
			v := line[col]

			//
			// Check if it is a *.
			// If it is, try to find two numbers adjacent to it.
			if v == '*' {
				ns, err := findAdjacentNumbers(lines, line, i, col)
				if err == nil {
					result += ns[0] * ns[1]
					fmt.Printf("row=%d, col=%d: %v\n", i, col, ns)
				} else {
					fmt.Printf("not found: %v\n", err)
				}
			}

			col++
		}
	}

	return result
}

func findNumber(lines []string, row, col int) int64 {
	line := lines[row]

	middleDigit := lines[row][col]
	leftDigits := []byte{}
	rightDigits := []byte{}

	if col > 0 && isNumber(lines[row][col-1]) {
		leftDigits = append(leftDigits, lines[row][col-1])
		for i := 2; col-i >= 0 && isNumber(lines[row][col-i]); i++ {
			leftDigits = append(leftDigits, lines[row][col-i])
		}
	}

	if col < len(line) && isNumber(lines[row][col+1]) {
		rightDigits = append(rightDigits, lines[row][col+1])
		for i := 2; col+i < len(line) && isNumber(lines[row][col+i]); i++ {
			rightDigits = append(rightDigits, lines[row][col+i])
		}
	}

	digits := reverse(leftDigits)
	digits = append(digits, middleDigit)
	for _, d := range rightDigits {
		digits = append(digits, d)
	}

	return byteArrToNumber(digits)
}

func reverse(a []byte) []byte {
	b := make([]byte, len(a))
	for i, n := range a {
		b[len(a)-1-i] = n
	}

	return b
}

func findAdjacentNumbers(lines []string, line string, row, col int) ([]int64, error) {
	lastCol := len(line)
	lastRow := len(lines)

	ns := []int64{}

	if col > 0 && row > 0 && isNumber(lines[row-1][col-1]) {
		ns = append(ns, findNumber(lines, row-1, col-1))
	}

	if col > 0 && isNumber(lines[row][col-1]) {
		ns = append(ns, findNumber(lines, row, col-1))
	}

	if col > 0 && row < lastRow-1 && isNumber(lines[row+1][col-1]) {
		ns = append(ns, findNumber(lines, row+1, col-1))
	}

	if row > 0 && isNumber(lines[row-1][col]) {
		ns = append(ns, findNumber(lines, row-1, col))
	}

	if row < lastRow-1 && isNumber(lines[row+1][col]) {
		ns = append(ns, findNumber(lines, row+1, col))
	}

	if row > 0 && col < lastCol-1 && isNumber(lines[row-1][col+1]) {
		ns = append(ns, findNumber(lines, row-1, col+1))
	}

	if row < lastRow-1 && col < lastCol-1 && isNumber(lines[row+1][col+1]) {
		ns = append(ns, findNumber(lines, row+1, col+1))
	}

	if col < lastCol-1 && isNumber(lines[row][col+1]) {
		ns = append(ns, findNumber(lines, row, col+1))
	}

	ns = removeDuplicates(ns)
	if len(ns) == 2 {
		return ns, nil
	}

	return ns, fmt.Errorf("no 2 numbers: %v", ns)
}

func removeDuplicates(slice []int64) []int64 {
	allKeys := make(map[int64]bool)
	list := []int64{}
	for _, item := range slice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func byteArrToNumber(n []byte) int64 {
	v, _ := strconv.ParseInt(string(n[:]), 10, 64)
	return v
}

func isAdjacentToSymbol(lines []string, line string, row, col int) bool {
	lastCol := len(line)
	lastRow := len(lines)

	if col > 0 && row > 0 && lines[row-1][col-1] != '.' {
		return true
	}

	if col > 0 && lines[row][col-1] != '.' && !isNumber(lines[row][col-1]) {
		return true
	}

	if col > 0 && row < lastRow-1 && lines[row+1][col-1] != '.' {
		return true
	}

	if row > 0 && lines[row-1][col] != '.' {
		return true
	}

	if row < lastRow-1 && lines[row+1][col] != '.' {
		return true
	}

	if row > 0 && col < lastCol-1 && lines[row-1][col+1] != '.' {
		return true
	}

	if row < lastRow-1 && col < lastCol-1 && lines[row+1][col+1] != '.' {
		return true
	}

	if col < lastCol-1 && !isNumber(lines[row][col+1]) && lines[row][col+1] != '.' {
		return true
	}

	return false
}

func isNumber(char byte) bool {
	return char >= '0' && char <= '9'
}

func parse(input string) []string {
	f, _ := os.Open(input)
	scanner := bufio.NewScanner(f)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
