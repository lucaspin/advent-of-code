package pkg202307

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	Cards []int
	Bid   int
	Type  int
}

func (h *Hand) CardsAsString() string {
	r := ""
	for _, c := range h.Cards {
		r += reverseCardValue(c)
	}

	return r
}

func (h *Hand) Compare(other Hand) bool {
	if h.Type == other.Type {
		for i := 0; i < len(h.Cards); i++ {
			if h.Cards[i] == other.Cards[i] {
				continue
			}

			return h.Cards[i] < other.Cards[i]
		}
	}

	return h.Type < other.Type
}

func highestValueInMap(freqs map[int]int) int {
	highest := 0
	for _, v := range freqs {
		if v > highest {
			highest = v
		}
	}

	return highest
}

func hasFreq(freqs map[int]int, x int) bool {
	for _, v := range freqs {
		if v == x {
			return true
		}
	}

	return false
}

func countFreq(freqs map[int]int, x int) int {
	count := 0
	for _, v := range freqs {
		if v == x {
			count += 1
		}
	}

	return count
}

func hasJoker(cards []int) bool {
	for _, c := range cards {
		if c == 1 {
			return true
		}
	}

	return false
}

func findJoker(cards []int) int {
	for i, c := range cards {
		if c == 1 {
			return i
		}
	}

	return -1
}

func ComputeTypeWithJoker(cards []int, maxAcc int) int {
	if hasJoker(cards) {
		replacements := []int{14, 13, 12, 10, 9, 8, 7, 6, 5, 4, 3, 2}
		index := findJoker(cards)
		c := make([]int, len(cards))
		copy(c, cards)
		curMax := 0

		for _, r := range replacements {
			c[index] = r
			v := ComputeTypeWithJoker(c, maxAcc)
			if v > curMax {
				curMax = v
			}
		}

		return curMax
	}

	return ComputeType(cards)
}

func ComputeType(cards []int) int {
	freqs := findFrequencies(cards)
	highestFreq := highestValueInMap(freqs)

	if highestFreq == 5 {
		return 7
	}

	if highestFreq == 4 {
		return 6
	}

	if highestFreq == 3 {
		if hasFreq(freqs, 2) {
			return 5
		} else {
			return 4
		}
	}

	if highestFreq == 2 {
		if countFreq(freqs, 2) == 2 {
			return 3
		} else {
			return 2
		}
	}

	return 1
}

func findFrequencies(cards []int) map[int]int {
	freqs := map[int]int{}

	for _, n := range cards {
		if _, ok := freqs[n]; ok {
			freqs[n] += 1
		} else {
			freqs[n] = 1
		}
	}

	return freqs
}

func A(input string) int64 {
	f, _ := os.Open(input)
	scanner := bufio.NewScanner(f)
	hands := []Hand{}
	for scanner.Scan() {
		hand := parseHand(scanner.Text())
		hands = append(hands, hand)
	}

	sort.SliceStable(hands, func(i, j int) bool {
		return hands[i].Compare(hands[j])
	})

	result := int64(0)
	for i := 0; i < len(hands); i++ {
		hand := hands[i]
		rank := i + 1
		fmt.Printf("Hand: %s rank: %d bid: %d\n", hand.CardsAsString(), rank, hand.Bid)
		result += int64(hand.Bid * rank)
	}

	return result
}

func B(input string) int64 {
	f, _ := os.Open(input)
	scanner := bufio.NewScanner(f)
	hands := []Hand{}
	for scanner.Scan() {
		hand := parseHandWithJoker(scanner.Text())
		hands = append(hands, hand)
	}

	sort.SliceStable(hands, func(i, j int) bool {
		return hands[i].Compare(hands[j])
	})

	result := int64(0)
	for i := 0; i < len(hands); i++ {
		hand := hands[i]
		rank := i + 1
		fmt.Printf("Hand: %s rank: %d bid: %d\n", hand.CardsAsString(), rank, hand.Bid)
		result += int64(hand.Bid * rank)
	}

	return result
}

func parseHand(line string) Hand {
	parts := strings.Split(line, " ")
	hand := Hand{Cards: []int{}}
	for _, n := range strings.Split(parts[0], "") {
		hand.Cards = append(hand.Cards, parseCardValue(n))
	}

	hand.Type = ComputeType(hand.Cards)
	bid, _ := strconv.Atoi(parts[1])
	hand.Bid = bid
	return hand
}

func parseHandWithJoker(line string) Hand {
	parts := strings.Split(line, " ")
	hand := Hand{Cards: []int{}}
	for _, n := range strings.Split(parts[0], "") {
		v := parseCardValueWithJoker(n)
		fmt.Printf("Card %s became %d\n", n, v)
		hand.Cards = append(hand.Cards, v)
	}

	fmt.Printf("Cards before compute: %v\n", hand.Cards)
	hand.Type = ComputeTypeWithJoker(hand.Cards, 0)
	fmt.Printf("Cards after compute: %v\n", hand.Cards)
	bid, _ := strconv.Atoi(parts[1])
	hand.Bid = bid
	return hand
}

func parseCardValue(card string) int {
	v, err := strconv.Atoi(card)

	// card is a number, just use its value
	if err == nil {
		return v
	}

	switch card {
	case "T":
		return 10
	case "J":
		return 11
	case "Q":
		return 12
	case "K":
		return 13
	case "A":
		return 14
	default:
		fmt.Printf("This should not happen\n")
		return 0
	}
}

func parseCardValueWithJoker(card string) int {
	v, err := strconv.Atoi(card)

	// card is a number, just use its value
	if err == nil {
		return v
	}

	if card == "J" {
		return 1
	}

	switch card {
	case "T":
		return 10
	case "Q":
		return 12
	case "K":
		return 13
	case "A":
		return 14
	default:
		fmt.Printf("This should not happen\n")
		return 0
	}
}

func reverseCardValue(card int) string {
	if card == 1 {
		return "J"
	}

	if card <= 9 {
		return fmt.Sprintf("%d", card)
	}

	if card == 10 {
		return "T"
	}

	if card == 11 {
		return "J"
	}

	if card == 12 {
		return "Q"
	}

	if card == 13 {
		return "K"
	}

	if card == 14 {
		return "A"
	}

	return "?"
}
