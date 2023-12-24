package pkg202324

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type stone struct {
	p Point
	v Velocity
}

func (l *stone) linegradient() float64 {
	other := Point{x: l.p.x + l.v.vx, y: l.p.y + l.v.vy}
	return float64(other.y-l.p.y) / float64(other.x-l.p.x)
}

func (l *stone) yintersect() float64 {
	// y = gradient*x + yintersect
	y := float64(l.p.y)
	gradient := l.linegradient()
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
func (l *stone) intersect(other stone, min, max float64) bool {
	a1 := l.linegradient()
	a2 := other.linegradient()
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

	stones := parseStones(string(d))

	intersections := 0
	for i := 0; i < len(stones); i++ {
		for j := i + 1; j < len(stones); j++ {
			if stones[i].intersect(stones[j], min, max) {
				intersections++
			}
		}
	}

	return intersections
}

// NOTE: this requires sagemath (https://www.sagemath.org) to be installed and available in the path.
func B(input string) int64 {
	d, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	stones := parseStones(string(d))

	// I'm guessing that if I find when my throw collapses with the first N blocks, I find it for everything.
	// I'm using 3 here, but any value >3 works.
	// I found 3 by experimenting, which didn't took long:
	// - 1 couldn't be used
	// - 2 didn't give a proper answer
	numStones := 3

	vars := []string{"x", "y", "z", "vx", "vy", "vz"}
	for i := 0; i < numStones; i++ {
		vars = append(vars, fmt.Sprintf("t%d", i+1))
	}

	equations := []string{}
	expressions := ""
	for i := 0; i < numStones; i++ {
		stone := stones[i]
		expressions += fmt.Sprintf("eq%d = x + (vx * t%d) == %d + (%d * t%d)\n", (i*numStones)+1, i+1, stone.p.x, stone.v.vx, i+1)
		expressions += fmt.Sprintf("eq%d = y + (vy * t%d) == %d + (%d * t%d)\n", (i*numStones)+2, i+1, stone.p.y, stone.v.vy, i+1)
		expressions += fmt.Sprintf("eq%d = z + (vz * t%d) == %d + (%d * t%d)\n", (i*numStones)+3, i+1, stone.p.z, stone.v.vz, i+1)
		equations = append(equations, []string{
			fmt.Sprintf("eq%d", (i*numStones)+1),
			fmt.Sprintf("eq%d", (i*numStones)+2),
			fmt.Sprintf("eq%d", (i*numStones)+3),
		}...)
	}

	script := ""
	script += "var ('" + strings.Join(vars, " ") + "')\n"
	script += expressions
	script += "print(solve([" + strings.Join(equations, ",") + "]," + strings.Join(vars, ",") + "))"

	f, err := os.CreateTemp("", "*.sage")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Expression: %s\n", script)

	_, err = f.Write([]byte(script))
	if err != nil {
		panic(err)
	}

	f.Close()
	defer os.Remove(f.Name())

	out, err := exec.Command("sage", f.Name()).Output()
	if err != nil {
		panic(fmt.Errorf("error executing sage: %v", err))
	}

	fmt.Printf("Result: %s", out)

	regex := regexp.MustCompile(`x == (\d+), y == (\d+), z == (\d+)*`)
	matches := regex.FindStringSubmatch(string(out))
	if len(matches) == 0 {
		panic("no matches")
	}

	x, err := strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		panic(fmt.Errorf("error parsing result: %v", err))
	}

	y, err := strconv.ParseInt(matches[2], 10, 64)
	if err != nil {
		panic(fmt.Errorf("error parsing result: %v", err))
	}

	z, err := strconv.ParseInt(matches[3], 10, 64)
	if err != nil {
		panic(fmt.Errorf("error parsing result: %v", err))
	}

	return x + y + z
}

func parseStones(data string) []stone {
	stones := []stone{}
	for _, line := range strings.Split(data, "\n") {
		pAndV := strings.Split(line, " @ ")
		ps := parseCommaSeparatedIntList(pAndV[0])
		vs := parseCommaSeparatedIntList(pAndV[1])
		stones = append(stones, stone{
			p: Point{x: ps[0], y: ps[1], z: ps[2]},
			v: Velocity{vx: vs[0], vy: vs[1], vz: vs[2]},
		})
	}

	return stones
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
