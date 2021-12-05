package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part one...\n")
	partOne()
	fmt.Printf("Part two...\n")
	partTwo()
}

func partOne() {
	measurements := readMeasurements()
	increases := 0

	for i := 1; i < len(measurements); i++ {
		if measurements[i] > measurements[i-1] {
			increases += 1
		}
	}

	fmt.Printf("Number of increases: %d\n", increases)
}

func partTwo() {
	measurements := readMeasurements()
	increases := 0
	slidingWindow := 3
	for i := 0; i < len(measurements)-slidingWindow; i++ {
		if measurements[i+slidingWindow] > measurements[i] {
			increases += 1
		}
	}

	fmt.Printf("Number of increases: %d\n", increases)
}

func readMeasurements() []int64 {
	bytes, _ := ioutil.ReadFile("../input.txt")
	return stringToInteger(strings.Split(string(bytes), "\n"))
}

func stringToInteger(list []string) []int64 {
	newList := []int64{}
	for _, item := range list {
		value, _ := strconv.ParseInt(item, 10, 64)
		newList = append(newList, value)
	}

	return newList
}
