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

func replace(line []string, index int, s string) []string {
	new := []string{}
	new = append(new, line[0:index]...)
	new = append(new, ".")
	new = append(new, line[index+1:]...)
	return new
}

func (r *Record) Find(arrangements *[]string, line []string, blocks []int, lineIndex, blockIndex, blockCount int) {
	if lineIndex == len(line) {
		if blockIndex == len(blocks)-1 && blocks[blockIndex] == blockCount {
			*arrangements = append(*arrangements, strings.Join(line, ""))
		} else if blockIndex == len(blocks) && blockCount == 0 {
			*arrangements = append(*arrangements, strings.Join(line, ""))
		} else {
		}

		return
	}

	if line[lineIndex] == "?" {
		r.Find(arrangements, replace(line, lineIndex, "."), blocks, lineIndex, blockIndex, blockCount)
		r.Find(arrangements, replace(line, lineIndex, "#"), blocks, lineIndex, blockIndex, blockCount)
		return
	}

	if line[lineIndex] == "." {
		if blockCount == 0 {
			r.Find(arrangements, line, blocks, lineIndex+1, blockIndex, 0)
			return
		}

		if blockIndex < len(blocks) && blockCount == blocks[blockIndex] {
			r.Find(arrangements, line, blocks, lineIndex+1, blockIndex+1, 0)
			return
		}
	}

	if line[lineIndex] == "#" {
		r.Find(arrangements, line, blocks, lineIndex+1, blockIndex, blockCount+1)
		return
	}
}

func (r *Record) FindArrangements() []string {
	arrangements := []string{}
	r.Find(&arrangements, strings.Split(r.Line, ""), r.Blocks, 0, 0, 0)
	return arrangements
}

// NOTE: don't call this, it takes FOREVER to run.
func (r *Record) FindArrangementsExpanded() []string {
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

	fmt.Printf("After expansion: %s, %v\n", newLine, newBlocks)

	arrangements := []string{}
	r.Find(&arrangements, strings.Split(newLine, ""), newBlocks, 0, 0, 0)
	return arrangements
}

func A(input string) int {
	result := 0
	for _, record := range parseRecords(input) {
		fmt.Printf("Going through record %s\n", record.Line)
		arrangements := record.FindArrangements()
		fmt.Printf("Arrangements: %d\n", len(arrangements))
		result += len(arrangements)
	}

	return result
}

func B(input string) int {
	result := 0
	for _, record := range parseRecords(input) {
		fmt.Printf("Going through record %s\n", record.Line)
		arrangements := record.FindArrangementsExpanded()
		fmt.Printf("Arrangements: %d\n", len(arrangements))
		// fmt.Printf("%v\n", arrangements[0:100])
		result += len(arrangements)
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
