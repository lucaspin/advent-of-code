package pkg202318

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Vertice struct {
	X int64
	Y int64
}

type Polygon struct {
	Vertices []Vertice
}

func (p *Polygon) Perimeter() float64 {
	perimeter := float64(0)
	for i := 0; i < len(p.Vertices); i++ {
		var p1, p2 Vertice
		p1 = p.Vertices[i]
		if i == len(p.Vertices)-1 {
			p2 = p.Vertices[0]
		} else {
			p2 = p.Vertices[i+1]
		}

		dist := math.Sqrt(math.Pow(float64(p2.X-p1.X), 2) + math.Pow(float64(p2.Y-p1.Y), 2))
		perimeter += dist
	}

	return perimeter
}

func (p *Polygon) Area() float64 {
	sum := float64(0)
	for i := len(p.Vertices) - 1; i >= 0; i-- {
		r := i
		l := i - 1
		if i == 0 {
			l = len(p.Vertices) - 1
		}

		sum += float64((p.Vertices[r].X * p.Vertices[l].Y) - (p.Vertices[l].X * p.Vertices[r].Y))
	}

	area := sum / 2
	return area + (p.Perimeter() / 2) + 1
}

func (p *Polygon) Last() Vertice {
	return p.Vertices[len(p.Vertices)-1]
}

func (p *Polygon) AddVertice(v Vertice) {
	p.Vertices = append(p.Vertices, v)
}

func A(input string) float64 {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	p := parsePolygon(f, false)
	return p.Area()
}

func B(input string) float64 {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	p := parsePolygon(f, true)
	return p.Area()
}

func parsePolygon(f *os.File, part2 bool) *Polygon {
	if part2 {
		return parsePolygon2(f)
	}

	p := &Polygon{Vertices: []Vertice{{X: 0, Y: 0}}}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		if len(parts) != 3 {
			panic("each line should have 3 parts")
		}

		direction := parts[0]
		count, _ := strconv.ParseInt(parts[1], 10, 64)
		fill(p, direction, count)
	}

	p.Vertices = p.Vertices[0 : len(p.Vertices)-1]

	return p
}

func parsePolygon2(f *os.File) *Polygon {
	p := &Polygon{Vertices: []Vertice{{X: 0, Y: 0}}}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		if len(parts) != 3 {
			panic("each line should have 3 parts")
		}

		hex := strings.ReplaceAll(parts[2], "(", "")
		hex = strings.ReplaceAll(hex, ")", "")
		direction := string(hex[len(hex)-1])
		count, err := strconv.ParseInt(hex[1:len(hex)-1], 16, 32)
		if err != nil {
			panic(err)
		}

		fill(p, direction, count)
	}

	p.Vertices = p.Vertices[0 : len(p.Vertices)-1]

	return p
}

func fill(p *Polygon, direction string, count int64) {
	current := p.Last()
	switch direction {
	case "R", "0":
		p.Vertices = append(p.Vertices, Vertice{
			X: current.X + count,
			Y: current.Y,
		})
	case "L", "2":
		p.Vertices = append(p.Vertices, Vertice{
			X: current.X - count,
			Y: current.Y,
		})
	case "D", "1":
		p.Vertices = append(p.Vertices, Vertice{
			X: current.X,
			Y: current.Y - count,
		})
	case "U", "3":
		p.Vertices = append(p.Vertices, Vertice{
			X: current.X,
			Y: current.Y + count,
		})
	default:
		panic(fmt.Errorf("unknown direction %s", direction))
	}
}
