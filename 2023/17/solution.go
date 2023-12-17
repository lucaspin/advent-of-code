package pkg202317

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func A(input string) int {
	data, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	grid := parseGrid(string(data))
	graph := buildGraph(grid)

	start := "0-0"
	end := fmt.Sprintf("%d-%d", len(grid)-1, len(grid[0])-1)
	fmt.Printf("start=%s end=%s\n", start, end)
	ans := graph.ShortestPathModified(start, end, false)
	return ans
}

func B(input string) int {
	data, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	grid := parseGrid(string(data))
	graph := buildGraph(grid)

	start := "0-0"
	end := fmt.Sprintf("%d-%d", len(grid)-1, len(grid[0])-1)
	fmt.Printf("start=%s end=%s\n", start, end)
	ans := graph.ShortestPathModified(start, end, true)
	return ans
}

func parseGrid(input string) [][]string {
	grid := [][]string{}
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, strings.Split(line, ""))
	}

	return grid
}

func buildGraph(grid [][]string) Graph {
	graph := Graph{}

	for rowIndex, row := range grid {
		for colIndex := range row {
			graph.Add(Node{
				ID:    fmt.Sprintf("%d-%d", rowIndex, colIndex),
				Edges: edgesForNode(grid, rowIndex, colIndex),
			})
		}
	}

	return graph
}

func edgesForNode(grid [][]string, row, col int) []Edge {
	edges := []Edge{}

	if row > 0 {
		edges = append(edges, Edge{
			Destination: fmt.Sprintf("%d-%d", row-1, col),
			Weight:      toInt(grid[row-1][col]),
			Direction:   "north",
		})
	}

	if row < len(grid)-1 {
		edges = append(edges, Edge{
			Destination: fmt.Sprintf("%d-%d", row+1, col),
			Weight:      toInt(grid[row+1][col]),
			Direction:   "south",
		})
	}

	if col > 0 {
		edges = append(edges, Edge{
			Destination: fmt.Sprintf("%d-%d", row, col-1),
			Weight:      toInt(grid[row][col-1]),
			Direction:   "west",
		})
	}

	if col < len(grid[row])-1 {
		edges = append(edges, Edge{
			Destination: fmt.Sprintf("%d-%d", row, col+1),
			Weight:      toInt(grid[row][col+1]),
			Direction:   "east",
		})
	}

	return edges
}

func toInt(v string) int {
	weight, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}

	return weight
}
