package pkg202312

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Record struct {
	Line   string
	Blocks []int
}

func (r *Record) Find(cache map[string]int, line string, blocks []int, lineIndex, blockIndex, blockCount int) int {
	key := fmt.Sprintf("i=%d;bi=%d;bc=%d", lineIndex, blockIndex, blockCount)
	if v, ok := cache[key]; ok {
		return v
	}

	if lineIndex == len(line) {
		if blockIndex == len(blocks)-1 && blocks[blockIndex] == blockCount {
			return 1
		} else if blockIndex == len(blocks) && blockCount == 0 {
			return 1
		}

		return 0
	}

	count := 0
	for _, c := range []byte{'.', '#'} {
		if line[lineIndex] == c || line[lineIndex] == '?' {
			if c == '.' && blockCount == 0 {
				count += r.Find(cache, line, blocks, lineIndex+1, blockIndex, blockCount)
			} else if c == '.' && blockCount > 0 && blockIndex < len(blocks) && blocks[blockIndex] == blockCount {
				count += r.Find(cache, line, blocks, lineIndex+1, blockIndex+1, 0)
			} else if c == '#' {
				count += r.Find(cache, line, blocks, lineIndex+1, blockIndex, blockCount+1)
			}
		}
	}

	cache[key] = count
	return count
}

func (r *Record) FindArrangements() int {
	var cache = map[string]int{}
	return r.Find(cache, r.Line, r.Blocks, 0, 0, 0)
}

func (r *Record) FindArrangementsExpanded() int {
	newLine := ""
	newBlocks := []int{}
	for i := 0; i < 5; i++ {
		newLine += r.Line

		// We don't add ? on the last iteration
		if i < 4 {
			newLine += "?"
		}

		newBlocks = append(newBlocks, r.Blocks...)
	}

	var cache = map[string]int{}
	return r.Find(cache, newLine, newBlocks, 0, 0, 0)
}

func A(input string) int {
	result := 0
	for _, record := range parseRecords(input) {
		fmt.Printf("Going through record %s\n", record.Line)
		arrangements := record.FindArrangements()
		fmt.Printf("Arrangements: %d\n", arrangements)
		result += arrangements
	}

	return result
}

func B(input string) int {
	result := 0
	for _, record := range parseRecords(input) {
		fmt.Printf("Going through record %s\n", record.Line)
		arrangements := record.FindArrangementsExpanded()
		fmt.Printf("Arrangements: %d\n", arrangements)
		result += arrangements
	}

	return result
}

func parseRecords(input string) []Record {
	records := []Record{}
	f, _ := os.Open(input)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		blocks := []int{}
		for _, b := range strings.Split(parts[1], ",") {
			v, _ := strconv.Atoi(b)
			blocks = append(blocks, v)
		}

		records = append(records, Record{Line: parts[0], Blocks: blocks})
	}

	return records
}
