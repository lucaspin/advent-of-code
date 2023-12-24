package pkg202324

import (
	"os"
	"strconv"
	"strings"
)

type Line struct {
	p Point
	v Velocity
}

func (l *Line) gradient() float64 {
	other := Point{x: l.p.x + l.v.vx, y: l.p.y + l.v.vy}
	return float64(other.y-l.p.y) / float64(other.x-l.p.x)
}

func (l *Line) yintersect() float64 {
	// y = gradient*x + yintersect
	y := float64(l.p.y)
	gradient := l.gradient()
	x := float64(l.p.x)

	return y - (gradient * x)
}

// if we have two lines:
// a1*x + b1*y + c1 = 0
// a2*x + b2*y + c2 = 0
//
// Then, the intersection will bappen at:
// x = (b1*c2 - b2*c1) / (a1*b2 - a2*b1)
// y = (c1*a2 - c2*a1) / (a1*b2 - a2*b1)
func (l *Line) intersect(other Line, min, max float64) bool {
	a1 := l.gradient()
	a2 := other.gradient()
	b1 := float64(-1)
	b2 := float64(-1)
	c1 := l.yintersect()
	c2 := other.yintersect()

	ix := (b1*c2 - b2*c1) / (a1*b2 - a2*b1)
	iy := (c1*a2 - c2*a1) / (a1*b2 - a2*b1)

	// in the past for l
	if l.v.vx > 0 && ix < float64(l.p.x) {
		return false
	} else if l.v.vx < 0 && ix > float64(l.p.x) {
		return false
	}

	// in the past for other
	if other.v.vx > 0 && ix < float64(other.p.x) {
		return false
	} else if other.v.vx < 0 && ix > float64(other.p.x) {
		return false
	}

	if ix <= max && ix >= min && iy <= max && iy >= min {
		return true
	}

	return false
}

type Point struct {
	x, y, z int
}

type Velocity struct {
	vx, vy, vz int
}

func A(input string, min, max float64) int {
	d, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	lines := parseLines(string(d))

	intersections := 0
	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			if lines[i].intersect(lines[j], min, max) {
				intersections++
			}
		}
	}

	return intersections
}

func parseLines(data string) []Line {
	lines := []Line{}
	for _, line := range strings.Split(data, "\n") {
		pAndV := strings.Split(line, " @ ")
		ps := parseCommaSeparatedIntList(pAndV[0])
		vs := parseCommaSeparatedIntList(pAndV[1])
		lines = append(lines, Line{
			p: Point{x: ps[0], y: ps[1], z: ps[2]},
			v: Velocity{vx: vs[0], vy: vs[1], vz: vs[2]},
		})
	}

	return lines
}

func parseCommaSeparatedIntList(s string) []int {
	list := []int{}

	for _, s := range strings.Split(s, ",") {
		v := strings.Trim(s, " ")
		if v != "" {
			i, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}

			list = append(list, i)
		}
	}

	return list
}
