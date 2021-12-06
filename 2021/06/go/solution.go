package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("After 80 days: %d\n", numberOfFishAfter(80))
	fmt.Printf("After 256 days: %d\n", numberOfFishAfter(256))
}

func numberOfFishAfter(days int) int {
	countMap := map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
	}

	initialState := readInitialState()
	for _, fish := range initialState {
		countMap[fish] += 1
	}

	for i := 1; i <= days; i++ {
		zeroCount := countMap[0]
		countMap[0] = countMap[1]
		countMap[1] = countMap[2]
		countMap[2] = countMap[3]
		countMap[3] = countMap[4]
		countMap[4] = countMap[5]
		countMap[5] = countMap[6]
		countMap[6] = countMap[7] + zeroCount
		countMap[7] = countMap[8]
		countMap[8] = zeroCount
	}

	result := 0
	for _, v := range countMap {
		result += v
	}

	return result
}

func readInitialState() []int {
	bytes, _ := ioutil.ReadFile("../input.txt")
	numbers := strings.Split(string(bytes), ",")

	initialState := []int{}
	for _, numberAsString := range numbers {
		number, _ := strconv.Atoi(numberAsString)
		initialState = append(initialState, number)
	}

	return initialState
}
