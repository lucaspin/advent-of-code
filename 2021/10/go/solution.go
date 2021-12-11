package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	// fmt.Println("Part one...")
	// partOne()
	fmt.Println("Part two...")
	partTwo()
}

func partOne() {
	lines := readLines()
	illegalChars := []string{}

	for _, line := range lines {
		symbols := strings.Split(line, "")
		stack := []string{}
		for _, symbol := range symbols {
			if isOpeningSymbol(symbol) {
				stack = append(stack, symbol)
			} else {
				lastOpeningSymbol := stack[len(stack)-1]
				if isMatch(lastOpeningSymbol, symbol) {
					stack = stack[0 : len(stack)-1]
				} else {
					illegalChars = append(illegalChars, symbol)
					break
				}
			}
		}
	}

	points := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	syntaxErrorScore := 0
	for _, illegal := range illegalChars {
		syntaxErrorScore += points[illegal]
	}

	fmt.Println(syntaxErrorScore)
}

func partTwo() {
	lines := readLines()

	incompleteLines := [][]string{}

	for _, line := range lines {
		symbols := strings.Split(line, "")
		stack := []string{}
		for _, symbol := range symbols {
			if isOpeningSymbol(symbol) {
				stack = append(stack, symbol)
			} else {
				lastOpeningSymbol := stack[len(stack)-1]
				if isMatch(lastOpeningSymbol, symbol) {
					stack = stack[0 : len(stack)-1]
				} else {
					// discard corrupted lines
					stack = []string{}
					break
				}
			}
		}

		if len(stack) > 0 {
			incompleteLines = append(incompleteLines, stack)
		}
	}

	completions := [][]string{}
	for _, incompleteLine := range incompleteLines {
		remaining := []string{}
		for _, symbol := range reverse(incompleteLine) {
			remaining = append(remaining, close(symbol))
		}

		completions = append(completions, remaining)
	}

	scoreMap := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}

	scores := []int{}
	for _, completion := range completions {
		total := 0
		for _, symbol := range completion {
			total = total * 5
			total = total + scoreMap[symbol]
		}

		scores = append(scores, total)
	}

	sort.SliceStable(scores, func(i, j int) bool {
		return scores[i] > scores[j]
	})

	fmt.Println(scores)

	middle := len(scores) / 2
	fmt.Println(scores[middle])
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

func close(symbol string) string {
	switch symbol {
	case "{":
		return "}"
	case "(":
		return ")"
	case "[":
		return "]"
	case "<":
		return ">"
	}

	return ""
}

func isOpeningSymbol(symbol string) bool {
	switch symbol {
	case "(", "{", "[", "<":
		return true
	default:
		return false
	}
}

func isMatch(openingSymbol, closingSymbol string) bool {
	if openingSymbol == "(" && closingSymbol == ")" {
		return true
	}

	if openingSymbol == "[" && closingSymbol == "]" {
		return true
	}

	if openingSymbol == "{" && closingSymbol == "}" {
		return true
	}

	if openingSymbol == "<" && closingSymbol == ">" {
		return true
	}

	return false
}

func readLines() []string {
	bytes, _ := ioutil.ReadFile("../input.txt")
	return strings.Split(string(bytes), "\n")
}
