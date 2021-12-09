package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	X     int
	Y     int
	Value int
}

func (p *Point) IsEqual(other Point) bool {
	return p.X == other.X && p.Y == other.Y && p.Value == other.Value
}

func main() {
	fmt.Println("Part one...")
	partOne()
	fmt.Println("Part two...")
	partTwo()
}

func partOne() {
	riskLevelSum := 0
	grid := readGrid()
	for _, lowPoint := range findLowPoints(grid) {
		riskLevelSum += lowPoint.Value + 1
	}

	fmt.Println(riskLevelSum)
}

func partTwo() {
	grid := readGrid()
	lowPoints := findLowPoints(grid)

	basins := [][]Point{}
	for _, lowPoint := range lowPoints {
		basin := visit(grid, []Point{}, []Point{lowPoint})
		basins = append(basins, basin)
	}

	sort.SliceStable(basins, func(i, j int) bool {
		return len(basins[i]) > len(basins[j])
	})

	result := 1
	for _, basin := range basins[:3] {
		result *= len(basin)
	}

	fmt.Println(result)
}

func visit(grid [][]int, visitedPoints []Point, pointsToVisit []Point) []Point {
	if len(pointsToVisit) == 0 {
		return visitedPoints
	}

	newPointsToVisit := visitPoints(grid, visitedPoints, pointsToVisit)
	visitedPoints = appendIfNotPresent(visitedPoints, pointsToVisit)
	return visit(grid, visitedPoints, newPointsToVisit)
}

func visitPoints(grid [][]int, visitedPoints []Point, pointsToVisit []Point) []Point {
	foundPoints := []Point{}
	for _, pointToVisit := range pointsToVisit {
		if !isPresent(visitedPoints, pointToVisit) {
			foundPoints = appendIfNotPresent(foundPoints, findIncrementalPoints(grid, pointToVisit))
		}
	}

	return foundPoints
}

func findIncrementalPoints(grid [][]int, point Point) []Point {
	x := point.X
	y := point.Y
	rowSize := len(grid[y])
	gridSize := len(grid)
	incrementalPoints := []Point{}

	switch {
	case isTopLeftCorner(x, y):
		incrementalPoints = append(incrementalPoints, fromRight(grid, point)...)
		incrementalPoints = append(incrementalPoints, fromDown(grid, point)...)
	case isTopRightCorner(x, y, rowSize):
		incrementalPoints = append(incrementalPoints, fromLeft(grid, point)...)
		incrementalPoints = append(incrementalPoints, fromDown(grid, point)...)
	case isBottomLeftCorner(x, y, gridSize):
		incrementalPoints = append(incrementalPoints, fromTop(grid, point)...)
		incrementalPoints = append(incrementalPoints, fromRight(grid, point)...)
	case isBottomRightCorner(x, y, gridSize, rowSize):
		incrementalPoints = append(incrementalPoints, fromTop(grid, point)...)
		incrementalPoints = append(incrementalPoints, fromLeft(grid, point)...)
	case isTopEdge(x, y, rowSize):
		incrementalPoints = append(incrementalPoints, fromRight(grid, point)...)
		incrementalPoints = append(incrementalPoints, fromLeft(grid, point)...)
		incrementalPoints = append(incrementalPoints, fromDown(grid, point)...)
	case isBottomEdge(x, y, gridSize, rowSize):
		incrementalPoints = append(incrementalPoints, fromRight(grid, point)...)
		incrementalPoints = append(incrementalPoints, fromLeft(grid, point)...)
		incrementalPoints = append(incrementalPoints, fromTop(grid, point)...)
	case isLeftEdge(x, y, gridSize):
		incrementalPoints = append(incrementalPoints, fromDown(grid, point)...)
		incrementalPoints = append(incrementalPoints, fromTop(grid, point)...)
		incrementalPoints = append(incrementalPoints, fromRight(grid, point)...)
	case isRightEdge(x, y, gridSize, rowSize):
		incrementalPoints = append(incrementalPoints, fromDown(grid, point)...)
		incrementalPoints = append(incrementalPoints, fromTop(grid, point)...)
		incrementalPoints = append(incrementalPoints, fromLeft(grid, point)...)
	default:
		incrementalPoints = append(incrementalPoints, fromDown(grid, point)...)
		incrementalPoints = append(incrementalPoints, fromTop(grid, point)...)
		incrementalPoints = append(incrementalPoints, fromLeft(grid, point)...)
		incrementalPoints = append(incrementalPoints, fromRight(grid, point)...)
	}

	return incrementalPoints
}

func appendIfNotPresent(points []Point, toAppend []Point) []Point {
	finalPoints := points
	for _, point := range toAppend {
		if !isPresent(finalPoints, point) {
			finalPoints = append(finalPoints, point)
		}
	}

	return finalPoints
}

func isPresent(points []Point, toFind Point) bool {
	for _, point := range points {
		if point.IsEqual(toFind) {
			return true
		}
	}

	return false
}

func fromLeft(grid [][]int, point Point) []Point {
	value := grid[point.Y][point.X-1]
	if value >= point.Value && value != 9 {
		return []Point{
			{X: point.X - 1, Y: point.Y, Value: value},
		}
	}

	return []Point{}
}

func fromRight(grid [][]int, point Point) []Point {
	value := grid[point.Y][point.X+1]
	if value >= point.Value && value != 9 {
		return []Point{
			{X: point.X + 1, Y: point.Y, Value: value},
		}
	}

	return []Point{}
}

func fromTop(grid [][]int, point Point) []Point {
	value := grid[point.Y-1][point.X]
	if value > point.Value && value != 9 {
		return []Point{
			{X: point.X, Y: point.Y - 1, Value: value},
		}
	}

	return []Point{}
}

func fromDown(grid [][]int, point Point) []Point {
	value := grid[point.Y+1][point.X]
	if value >= point.Value && value != 9 {
		return []Point{
			{X: point.X, Y: point.Y + 1, Value: value},
		}
	}

	return []Point{}
}

func findLowPoints(grid [][]int) []Point {
	gridSize := len(grid)
	lowPoints := []Point{}

	for y, row := range grid {
		rowSize := len(row)
		for x, value := range row {
			switch {
			case isTopLeftCorner(x, y):
				if value < row[x+1] && value < grid[y+1][x] {
					lowPoints = append(lowPoints, Point{X: x, Y: y, Value: value})
				}
			case isTopRightCorner(x, y, rowSize):
				if value < row[x-1] && value < grid[y+1][x] {
					lowPoints = append(lowPoints, Point{X: x, Y: y, Value: value})
				}
			case isBottomLeftCorner(x, y, gridSize):
				if value < row[x+1] && value < grid[y-1][x] {
					lowPoints = append(lowPoints, Point{X: x, Y: y, Value: value})
				}
			case isBottomRightCorner(x, y, gridSize, rowSize):
				if value < row[x-1] && value < grid[y-1][x] {
					lowPoints = append(lowPoints, Point{X: x, Y: y, Value: value})
				}
			case isTopEdge(x, y, rowSize):
				if value < row[x-1] && value < row[x+1] && value < grid[y+1][x] {
					lowPoints = append(lowPoints, Point{X: x, Y: y, Value: value})
				}
			case isBottomEdge(x, y, gridSize, rowSize):
				if value < row[x-1] && value < row[x+1] && value < grid[y-1][x] {
					lowPoints = append(lowPoints, Point{X: x, Y: y, Value: value})
				}
			case isLeftEdge(x, y, gridSize):
				if value < row[x+1] && value < grid[y-1][x] && value < grid[y+1][x] {
					lowPoints = append(lowPoints, Point{X: x, Y: y, Value: value})
				}
			case isRightEdge(x, y, gridSize, rowSize):
				if value < row[x-1] && value < grid[y-1][x] && value < grid[y+1][x] {
					lowPoints = append(lowPoints, Point{X: x, Y: y, Value: value})
				}
			default:
				if value < row[x-1] && value < row[x+1] && value < grid[y-1][x] && value < grid[y+1][x] {
					lowPoints = append(lowPoints, Point{X: x, Y: y, Value: value})
				}
			}
		}
	}

	return lowPoints
}

func isTopLeftCorner(x, y int) bool {
	return x == 0 && y == 0
}

func isBottomLeftCorner(x, y, gridSize int) bool {
	return x == 0 && y == gridSize-1
}

func isTopRightCorner(x, y, rowSize int) bool {
	return x == rowSize-1 && y == 0
}

func isBottomRightCorner(x, y, gridSize, rowSize int) bool {
	return x == rowSize-1 && y == gridSize-1
}

func isTopEdge(x, y, rowSize int) bool {
	return y == 0 && !isTopLeftCorner(x, y) && !isTopRightCorner(x, y, rowSize)
}

func isBottomEdge(x, y, gridSize, rowSize int) bool {
	return y == gridSize-1 && !isBottomLeftCorner(x, y, gridSize) && !isBottomRightCorner(x, y, gridSize, rowSize)
}

func isLeftEdge(x, y, gridSize int) bool {
	return x == 0 && !isBottomLeftCorner(x, y, gridSize) && !isTopLeftCorner(x, y)
}

func isRightEdge(x, y, gridSize, rowSize int) bool {
	return x == rowSize-1 && !isBottomRightCorner(x, y, gridSize, rowSize) && !isTopRightCorner(x, y, gridSize)
}

func readGrid() [][]int {
	bytes, _ := ioutil.ReadFile("../input.txt")
	lines := strings.Split(string(bytes), "\n")

	grid := [][]int{}
	for _, line := range lines {
		row := stringToInt(strings.Split(line, ""))
		grid = append(grid, row)
	}

	return grid
}

func stringToInt(numbers []string) []int {
	newNumbers := []int{}
	for _, number := range numbers {
		value, _ := strconv.Atoi(number)
		newNumbers = append(newNumbers, value)
	}

	return newNumbers
}
