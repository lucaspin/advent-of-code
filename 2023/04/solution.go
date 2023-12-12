package pkg202304

import (
	"bufio"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var cardRegex = regexp.MustCompile(`^Card\s*(\d*): (.*) \| (.*)$`)

type Deck struct {
	Cards map[int]Card
}

func (d *Deck) Process() {
	for _, k := range d.Keys() {
		c := d.Cards[k]
		for i := 1; i <= c.MatchingNumbers(); i++ {
			d.AddCopies(c.ID+i, c.Copies)
		}
	}
}

func (d *Deck) Keys() []int {
	keys := []int{}
	for k := range d.Cards {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	return keys
}

func (d *Deck) TotalCards() int64 {
	total := int64(0)
	for _, c := range d.Cards {
		total += int64(c.Copies)
	}

	return total
}

func (d *Deck) AddCard(card Card) {
	card.Copies = 1
	d.Cards[card.ID] = card
}

func (d *Deck) AddCopies(cardID int, copies int) {
	if c, ok := d.Cards[cardID]; ok {
		c.Copies += copies
		d.Cards[cardID] = c
	}
}

type Card struct {
	ID      int
	Copies  int
	Winning []int
	Holding []int
}

func (c *Card) Score() int64 {
	score := int64(0)

	for _, n := range c.Holding {
		if c.IsWinning(n) {
			if score == 0 {
				score = 1
				continue
			}

			score *= 2
		}
	}

	return score
}

func (c *Card) MatchingNumbers() int {
	matching := 0
	for _, n := range c.Holding {
		if c.IsWinning(n) {
			matching += 1
		}
	}

	return matching
}

func (c *Card) IsWinning(n int) bool {
	for _, m := range c.Winning {

		if n == m {
			return true
		}
	}

	return false
}

func A(input string) int64 {
	f, err := os.Open(input)
	if err != nil {
		return 0
	}

	total := int64(0)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		card := parseCard(scanner.Text())
		total += card.Score()
	}

	return total
}

func B(input string) int64 {
	f, err := os.Open(input)
	if err != nil {
		return 0
	}

	deck := Deck{
		Cards: map[int]Card{},
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		card := parseCard(scanner.Text())

		deck.AddCard(card)
	}

	deck.Process()

	return deck.TotalCards()
}

func parseCard(line string) Card {
	matches := cardRegex.FindAllStringSubmatch(line, 2)
	ID, _ := strconv.Atoi(matches[0][1])

	card := Card{
		ID:      ID,
		Winning: toList(matches[0][2]),
		Holding: toList(matches[0][3]),
	}

	return card
}

func toList(list string) []int {
	ns := []int{}
	for _, n := range removeEmpty(strings.Split(list, " ")) {
		v, _ := strconv.Atoi(strings.Trim(n, " "))
		ns = append(ns, v)
	}

	return ns
}

func removeEmpty(list []string) []string {
	new := []string{}
	for _, s := range list {
		if s != "" {
			new = append(new, s)
		}
	}

	return new
}
