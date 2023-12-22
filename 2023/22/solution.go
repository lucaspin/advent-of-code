package pkg202322

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
	Z int
}

type Brick struct {
	Start Point
	End   Point
}

type XY struct {
	X int
	Y int
}

func maxHeight(heights map[XY]int, xys []XY) int {
	max := 0
	for _, xy := range xys {
		if heights[xy] > max {
			max = heights[xy]
		}
	}

	return max
}

func drop(bricks *[]*Brick, update bool, ignore *Brick) int {
	dropped := 0
	heights := map[XY]int{}

	for _, brick := range *bricks {
		if ignore != nil && *brick == *ignore {
			continue
		}

		xys := []XY{}
		for x := brick.Start.X; x < brick.End.X+1; x++ {
			for y := brick.Start.Y; y < brick.End.Y+1; y++ {
				xys = append(xys, XY{X: x, Y: y})
			}
		}

		height := maxHeight(heights, xys)

		zEnd := brick.End.Z
		drop := brick.Start.Z - height - 1
		if drop > 0 {
			dropped += 1
			zEnd -= drop
			if update {
				brick.Start.Z -= drop
				brick.End.Z -= drop
			}
		}

		for _, xy := range xys {
			heights[xy] = zEnd
		}
	}

	return dropped
}

func count(list []int, v int) int {
	c := 0
	for _, i := range list {
		if i == v {
			c++
		}
	}

	return c
}

func A(input string) int {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	bricks := parse(f)
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].Start.Z < bricks[j].Start.Z
	})

	drop(&bricks, true, nil)

	dropCounts := []int{}
	for _, brick := range bricks {
		dropCounts = append(dropCounts, drop(&bricks, false, brick))
	}

	return count(dropCounts, 0)
}

func sum(list []int) int {
	s := 0
	for _, i := range list {
		s += i
	}

	return s
}

func B(input string) int {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	bricks := parse(f)
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].Start.Z < bricks[j].Start.Z
	})

	drop(&bricks, true, nil)

	dropCounts := []int{}
	for _, brick := range bricks {
		dropCounts = append(dropCounts, drop(&bricks, false, brick))
	}

	return sum(dropCounts)
}

func parse(f *os.File) []*Brick {
	bricks := []*Brick{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		bricks = append(bricks, parseBrick(scanner.Text()))
	}

	return bricks
}

func parseBrick(line string) *Brick {
	startAndEnd := strings.Split(line, "~")
	start := strings.Split(startAndEnd[0], ",")
	end := strings.Split(startAndEnd[1], ",")

	return &Brick{
		Start: Point{
			X: toI(start[0]),
			Y: toI(start[1]),
			Z: toI(start[2]),
		},
		End: Point{
			X: toI(end[0]),
			Y: toI(end[1]),
			Z: toI(end[2]),
		},
	}
}

func toI(i string) int {
	v, err := strconv.Atoi(i)
	if err != nil {
		panic(err)
	}

	return v
}
