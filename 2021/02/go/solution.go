package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Submarine struct {
	HPosition int
	Depth     int
	Aim       int
}

func (s *Submarine) Navigate(line string) {
	commandAndValue := strings.Split(line, " ")
	command := commandAndValue[0]
	value, _ := strconv.Atoi(commandAndValue[1])

	switch command {
	case "forward":
		s.HPosition += value
	case "down":
		s.Depth += value
	case "up":
		s.Depth -= value
	}
}

func (s *Submarine) Navigate2(line string) {
	commandAndValue := strings.Split(line, " ")
	command := commandAndValue[0]
	value, _ := strconv.Atoi(commandAndValue[1])

	switch command {
	case "forward":
		s.HPosition += value
		s.Depth += s.Aim * value
	case "down":
		s.Aim += value
	case "up":
		s.Aim -= value
	}
}

func (s *Submarine) Result() int {
	return s.HPosition * s.Depth
}

func main() {
	fmt.Printf("Part one...\n")
	partOne()
	fmt.Printf("Part one...\n")
	partTwo()
}

func partOne() {
	submarine := Submarine{}
	for _, command := range readCommands() {
		submarine.Navigate(command)
	}

	fmt.Printf("Result: %d\n", submarine.Result())
}

func partTwo() {
	submarine := Submarine{}
	for _, command := range readCommands() {
		submarine.Navigate2(command)
	}

	fmt.Printf("Result: %d\n", submarine.Result())
}

func readCommands() []string {
	bytes, _ := ioutil.ReadFile("../input.txt")
	return strings.Split(string(bytes), "\n")
}
