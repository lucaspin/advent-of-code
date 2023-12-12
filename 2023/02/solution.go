package pkg202302

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var gameRegex, _ = regexp.Compile(`^Game (\d+): (.*)$`)

type Bag struct {
	Red   int
	Green int
	Blue  int
}

func (b *Bag) FillWithColor(draw string) {
	i := strings.Split(strings.Trim(draw, " "), " ")
	n, _ := strconv.Atoi(i[0])
	color := i[1]

	switch color {
	case "blue":
		b.Blue = n
	case "green":
		b.Green = n
	case "red":
		b.Red = n
	default:
		// nothing really
	}
}

func (b *Bag) IsPossible(available Bag) bool {
	fmt.Printf("Checking set: red=%d, green=%d, blue=%d - red=%d, green=%d, blue=%d\n", b.Red, b.Green, b.Blue, available.Red, available.Green, available.Blue)
	return available.Blue >= b.Blue && available.Green >= b.Green && available.Red >= b.Red
}

type Game struct {
	ID   int
	Sets []Bag
}

func (g *Game) HasImpossibleSet(available Bag) bool {
	for _, set := range g.Sets {
		if !set.IsPossible(available) {
			return false
		}
	}

	return true
}

func (g *Game) FindMinimumBag() Bag {
	blue := 0
	green := 0
	red := 0

	for _, set := range g.Sets {
		if set.Blue > blue {
			blue = set.Blue
		}

		if set.Green > green {
			green = set.Green
		}

		if set.Red > red {
			red = set.Red
		}
	}

	return Bag{
		Red:   red,
		Green: green,
		Blue:  blue,
	}
}

func A(available Bag, input string) int {
	f, _ := os.Open(input)
	defer f.Close()

	result := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		game := parseGame(scanner.Text())
		fmt.Printf("Checking game %d: %v\n", game.ID, game.Sets)
		if game.HasImpossibleSet(available) {
			result += game.ID
		}
	}

	return result
}

func B(available Bag, input string) int64 {
	f, _ := os.Open(input)
	defer f.Close()

	result := int64(0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		game := parseGame(scanner.Text())
		fmt.Printf("Checking game %d: %v\n", game.ID, game.Sets)
		bag := game.FindMinimumBag()
		power := bag.Green * bag.Red * bag.Blue
		result += int64(power)
	}

	return result
}

func parseGame(line string) Game {
	matches := gameRegex.FindAllStringSubmatch(line, 2)
	gameID, _ := strconv.Atoi(matches[0][1])
	game := Game{
		ID:   gameID,
		Sets: []Bag{},
	}

	for _, set := range strings.Split(matches[0][2], ";") {
		bag := Bag{}
		for _, draw := range strings.Split(set, ",") {
			bag.FillWithColor(draw)
		}

		game.Sets = append(game.Sets, bag)
	}

	return game
}
