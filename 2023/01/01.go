package pkg202301

import (
	"fmt"
	"strconv"
	"strings"
)

func First(lines []string) int64 {
	total := int64(0)

	for _, line := range lines {
		digits := []int64{}

		for _, character := range strings.Split(line, "") {
			if isIn([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}, character) {
				n, _ := strconv.ParseInt(character, 10, 64)
				digits = append(digits, n)
			}
		}

		fmt.Printf("Line: %s, digits: %v\n", line, digits)
		total += (digits[0] * 10) + digits[len(digits)-1]
	}

	return total
}

func Second(lines []string) int64 {
	newLines := []string{}
	for _, line := range lines {
		newLines = append(newLines, replace(line))
	}

	return First(newLines)
}

func replace(line string) string {
	line = strings.ReplaceAll(line, "one", "o1e")
	line = strings.ReplaceAll(line, "two", "t2o")
	line = strings.ReplaceAll(line, "three", "t3e")
	line = strings.ReplaceAll(line, "four", "f4")
	line = strings.ReplaceAll(line, "five", "f5e")
	line = strings.ReplaceAll(line, "six", "s6")
	line = strings.ReplaceAll(line, "seven", "s7n")
	line = strings.ReplaceAll(line, "eight", "e8t")
	line = strings.ReplaceAll(line, "nine", "n9e")
	return strings.ReplaceAll(line, "zero", "0o")
}

func isIn(items []string, item string) bool {
	for _, i := range items {
		if i == item {
			return true
		}
	}

	return false
}
