package pkg202323

import (
	"bufio"
	"os"
	"strings"
)

type Location struct {
	Row  int
	Col  int
	Type string
}

func findStart(locations [][]Location) Location {
	for _, t := range locations[0] {
		if t.Type == "." {
			return t
		}
	}

	panic("should not happen")
}

func next(locations [][]Location, location Location, part2 bool) []Location {
	next := []Location{}
	if location.Row > 0 {
		up := locations[location.Row-1][location.Col]
		if !part2 {
			if up.Type == "." || up.Type == "^" {
				next = append(next, up)
			}
		} else {
			if up.Type != "#" {
				next = append(next, up)
			}
		}
	}

	if location.Col > 0 {
		left := locations[location.Row][location.Col-1]
		if !part2 {
			if left.Type == "." || left.Type == "<" {
				next = append(next, left)
			}
		} else {
			if left.Type != "#" {
				next = append(next, left)
			}
		}
	}

	if location.Row < len(locations)-1 {
		down := locations[location.Row+1][location.Col]
		if !part2 {
			if down.Type == "." || down.Type == "v" {
				next = append(next, down)
			}
		} else {
			if down.Type != "#" {
				next = append(next, down)
			}
		}
	}

	if location.Col < len(locations[location.Row])-1 {
		right := locations[location.Row][location.Col+1]
		if !part2 {
			if right.Type == "." || right.Type == ">" {
				next = append(next, right)
			}
		} else {
			if right.Type != "#" {
				next = append(next, right)
			}
		}
	}

	return next
}

type queue struct {
	items []queueItem
}

func (q *queue) Push(item queueItem) {
	q.items = append(q.items, item)
}

func (q *queue) Pop() queueItem {
	head := q.items[0]
	q.items = q.items[1:]
	return head
}

func (q *queue) Len() int {
	return len(q.items)
}

type queueItem struct {
	Location Location
	seen     LocationSet
	path     []Location
}

type LocationSet map[Location]bool

func (n *LocationSet) Has(v Location) bool {
	_, ok := (*n)[v]
	return ok
}

func (n *LocationSet) Add(v Location) {
	(*n)[v] = true
}

func (n *LocationSet) Len() int {
	return len(*n)
}

func (n *LocationSet) Values() []Location {
	ls := []Location{}
	for k := range *n {
		ls = append(ls, k)
	}

	return ls
}

func (n *LocationSet) Copy() LocationSet {
	new := make(map[Location]bool, len(*n))
	for k, v := range *n {
		new[k] = v
	}

	return new
}

func _search(locations [][]Location, queue queue, part2 bool) [][]Location {
	paths := [][]Location{}
	for queue.Len() > 0 {
		current := queue.Pop()

		// Reached bottom row
		if current.Location.Row == len(locations)-1 {
			paths = append(paths, current.path)
			continue
		}

		next := next(locations, current.Location, part2)
		if len(next) == 0 {
			continue
		}

		for _, n := range next {
			if current.seen.Has(n) {
				continue
			}

			newSeen := current.seen.Copy()
			newSeen.Add(n)

			newPath := make([]Location, len(current.path))
			copy(newPath, current.path)
			newPath = append(newPath, n)

			queue.Push(queueItem{
				Location: n,
				seen:     newSeen,
				path:     newPath,
			})
		}
	}

	return paths
}

func longest(locations [][]Location, part2 bool) int {
	start := findStart(locations)
	seen := LocationSet{}
	seen.Add(start)

	queue := queue{}
	queue.Push(queueItem{
		Location: start,
		seen:     seen,
		path:     []Location{},
	})

	lengths := []int{}
	for _, p := range _search(locations, queue, part2) {
		lengths = append(lengths, len(p))
	}

	return max(lengths)
}

func max(list []int) int {
	max := 0
	for _, i := range list {
		if i > max {
			max = i
		}
	}

	return max
}

func A(input string) int {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	locations := parse(f)
	return longest(locations, false)
}

func B(input string) int {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	locations := parse(f)
	return longest(locations, true)
}

func parse(f *os.File) [][]Location {
	locations := [][]Location{}
	scanner := bufio.NewScanner(f)
	row := 0
	for scanner.Scan() {
		r := []Location{}
		for col, t := range strings.Split(scanner.Text(), "") {
			r = append(r, Location{Row: row, Col: col, Type: t})
		}
		locations = append(locations, r)
		row++
	}

	return locations
}
