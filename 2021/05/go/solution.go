package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

/**

0,9 -> 5,9 ===> [{0,9}, {1,9}, {2, 9}, {3, 9}, {4, 9}, {5, 9}]
8,0 -> 0,8 ===> discard or [{8,0}, {7, 1}, {6, 2}, {5, 3}, {4, 4}, {3, 5}, {2, 6}, {1, 7}, {0, 8}], slope -1
9,4 -> 3,4 ===> [{9,4}, {8,4}, {7,4}, {6,4}, {5,4}, {4,4}, {3,4}]
2,2 -> 2,1 ===> [{2,2}, {2,1}]
7,0 -> 7,4 ===> [{7,0}, {7,1}, {7,2}, {7,3}, {7,4}]
6,4 -> 2,0 ===> discard or [{6,4}, {5,3}, {4,2}, {3,1}, {2,0}], slope = 1
0,9 -> 2,9 ===> [{0,9}, {1,9}, {2,9}]
3,4 -> 1,4 ===> [{3,4}, {2,4}, {1,4}]
0,0 -> 8,8 ===> discard or [{0,0}, {1,1}, {2,2}, {3,3}, {4,4}, {5,5}, {6,6}, {7,7}, {8,8}], slope = 1
5,5 -> 8,2 ===> discard or [{5,5}, {6,4}, {7,3}, {8,2}], slope = -1

9 | 2 2 2 1 1 1 0 0 0 0
8 | 0 0 0 0 0 0 0 0 0 0
7 | 0 0 0 0 0 0 0 0 0 0
6 | 0 0 0 0 0 0 0 0 0 0
5 | 0 0 0 0 0 0 0 0 0 0
4 | 0 1 1 2 1 1 1 2 1 1
3 | 0 0 0 0 0 0 0 1 0 0
2 | 0 0 1 0 0 0 0 1 0 0
1 | 0 0 1 0 0 0 0 1 0 0
0 | 0 0 0 0 0 0 0 1 0 0
    - - - - - - - - - -
    0 1 2 3 4 5 6 7 8 9

counting the cells where n >= 2 ==> 5

Part 2: same thing, but consider diagonal lines with a 45 degree slope
*/

type Point struct {
	X int
	Y int
}

// coordinates is something like x,y
func NewPoint(coordinates string) Point {
	xAndY := strings.Split(coordinates, ",")
	x, _ := strconv.Atoi(xAndY[0])
	y, _ := strconv.Atoi(xAndY[1])
	return Point{X: x, Y: y}
}

func pointsInStraightLine(start, end int, fun func(int) Point) []Point {
	points := []Point{}
	if start < end {
		for i := start; i <= end; i++ {
			points = append(points, fun(i))
		}
	} else {
		for i := end; i <= start; i++ {
			points = append(points, fun(i))
		}
	}

	return points
}

func pointsInVerticalLine(start, end Point) []Point {
	return pointsInStraightLine(start.Y, end.Y, func(i int) Point {
		return Point{X: start.X, Y: i}
	})
}

func pointsInHorizontalLine(start, end Point) []Point {
	return pointsInStraightLine(start.X, end.X, func(i int) Point {
		return Point{X: i, Y: start.Y}
	})
}

func PointsInBetween(start, end Point, considerDiagonals bool) []Point {
	if start.X == end.X {
		return pointsInVerticalLine(start, end)
	}

	if start.Y == end.Y {
		return pointsInHorizontalLine(start, end)
	}

	if considerDiagonals && Has45DegreesSlope(start, end) {
		return pointsInBetweenDiagonal(start, end)
	}

	// diagonals not considered, ignore
	return []Point{}
}

func pointsInBetweenDiagonal(start, end Point) []Point {
	points := []Point{}
	currentX := start.X
	currentY := start.Y

	for currentX != end.X && currentY != end.Y {
		points = append(points, Point{X: currentX, Y: currentY})
		if currentX > end.X {
			currentX--
		} else {
			currentX++
		}

		if currentY > end.Y {
			currentY--
		} else {
			currentY++
		}
	}

	points = append(points, Point{X: end.X, Y: end.Y})
	return points
}

// slope = delta(y) / delta(x) = y2 - y1 / x2 - x1
// a slope of 1  = a slope of +45 degrees
// a slope of -1 = a slope of -45 degrees
func Has45DegreesSlope(start, end Point) bool {
	deltaY := end.Y - start.Y
	deltaX := end.X - start.X
	slope := deltaY / deltaX
	return slope == 1 || slope == -1
}

func main() {
	fmt.Println("Part one...")
	partOne()
	fmt.Println("Part two...")
	partTwo()
}

func partOne() {
	board := [1000][1000]int{}
	pointsToMark := findPointsToMark(false)
	for _, point := range pointsToMark {
		board[point.X][point.Y] += 1
	}

	count := 0
	for _, row := range board {
		for _, column := range row {
			if column >= 2 {
				count += 1
			}
		}
	}

	fmt.Println(count)
}

func partTwo() {
	board := [1000][1000]int{}
	pointsToMark := findPointsToMark(true)
	for _, point := range pointsToMark {
		board[point.X][point.Y] += 1
	}

	count := 0
	for _, row := range board {
		for _, column := range row {
			if column >= 2 {
				count += 1
			}
		}
	}

	fmt.Println(count)
}

func findPointsToMark(considerDiagonals bool) []Point {
	bytes, _ := ioutil.ReadFile("../input.txt")
	lines := strings.Split(string(bytes), "\n")
	points := []Point{}

	for _, line := range lines {
		startAndEnd := strings.Split(line, "->")
		startPoint := NewPoint(strings.Trim(startAndEnd[0], " "))
		endPoint := NewPoint(strings.Trim(startAndEnd[1], " "))
		pointsInBetween := PointsInBetween(startPoint, endPoint, considerDiagonals)
		points = append(points, pointsInBetween...)
	}

	return points
}
