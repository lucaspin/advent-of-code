package pkg202315

import (
	"os"
	"strconv"
	"strings"
)

type Box struct {
	Slots []Slot
	Code  int
}

func (b *Box) Remove(label string) {
	new := []Slot{}
	for _, s := range b.Slots {
		if s.Label != label {
			new = append(new, s)
		}
	}

	b.Slots = new
}

func (b *Box) Update(label string, value int) {
	new := []Slot{}
	exists := false
	for _, s := range b.Slots {
		if s.Label == label {
			new = append(new, Slot{Strength: value, Label: label})
			exists = true
		} else {
			new = append(new, s)
		}
	}

	if !exists {
		new = append(new, Slot{Label: label, Strength: value})
	}

	b.Slots = new
}

type Slot struct {
	Label    string
	Strength int
}

func Hash(s string) int {
	current := 0
	for i := 0; i < len(s); i++ {
		current += int(s[i])
		current *= 17
		current = current % 256
	}

	return current
}

func A(input string) int {
	data, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	total := 0
	for _, sequence := range strings.Split(string(data), ",") {
		s := strings.Trim(sequence, "\n")
		if s == "" {
			continue
		}

		total += Hash(s)
	}

	return total
}

func B(input string) int {
	data, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	boxes := map[int]*Box{}
	for i := 0; i < 256; i++ {
		boxes[i] = &Box{Code: i, Slots: []Slot{}}
	}

	for _, sequence := range strings.Split(string(data), ",") {
		s := strings.Trim(sequence, "\n")
		if s == "" {
			continue
		}

		if strings.HasSuffix(s, "-") {
			s = strings.Replace(s, "-", "", 1)
			boxes[Hash(s)].Remove(s)
			continue
		}

		parts := strings.Split(s, "=")
		label := parts[0]
		value, _ := strconv.Atoi(parts[1])
		boxes[Hash(label)].Update(label, value)
	}

	total := 0
	for key, box := range boxes {
		for slotIndex, slot := range box.Slots {
			total += (1 + key) * (slotIndex + 1) * slot.Strength
		}

	}

	return total
}
