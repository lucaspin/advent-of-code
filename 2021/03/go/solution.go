package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Submarine struct{}

func (s *Submarine) PowerConsumption(values []string) int {
	sample := values[0]
	sample_size := len(sample)
	gammaRateBinary := []string{}
	epsilonRateBinary := []string{}

	for i := 0; i < sample_size; i++ {
		mostCommonBit := mostCommonBit(values, i)
		gammaRateBinary = append(gammaRateBinary, mostCommonBit)
		leastCommonBit := leastCommonBit(mostCommonBit)
		epsilonRateBinary = append(epsilonRateBinary, leastCommonBit)
	}

	gammaRate, _ := strconv.ParseInt(strings.Join(gammaRateBinary, ""), 2, 16)
	epsilonRate, _ := strconv.ParseInt(strings.Join(epsilonRateBinary, ""), 2, 16)
	return int(gammaRate) * int(epsilonRate)
}

func (s *Submarine) LifeSupportRating(values []string) int {
	sample := values[0]
	sample_size := len(sample)

	o2Values := values
	for i := 0; i < sample_size; i++ {
		mostCommon := mostCommonBit(o2Values, i)
		o2Values = filterByBit(o2Values, i, mostCommon)
		if len(o2Values) == 1 {
			break
		}
	}

	co2Values := values
	for i := 0; i < sample_size; i++ {
		mostCommon := mostCommonBit(co2Values, i)
		co2Values = filterByBit(co2Values, i, leastCommonBit(mostCommon))
		if len(co2Values) == 1 {
			break
		}
	}

	o2Binary := o2Values[0]
	co2Binary := co2Values[0]
	o2Value, _ := strconv.ParseInt(o2Binary, 2, 16)
	co2Value, _ := strconv.ParseInt(co2Binary, 2, 16)
	return int(o2Value) * int(co2Value)
}

func mostCommonBit(values []string, index int) string {
	zeroCount := 0
	oneCount := 0
	for _, number := range values {
		bit := number[index]
		if string(bit) == "0" {
			zeroCount += 1
		} else {
			oneCount += 1
		}
	}

	if zeroCount > oneCount {
		return "0"
	} else {
		return "1"
	}
}

func leastCommonBit(mostCommonBit string) string {
	if mostCommonBit == "0" {
		return "1"
	} else {
		return "0"
	}
}

func filterByBit(values []string, index int, bit string) []string {
	newValues := []string{}
	for _, value := range values {
		if string(value[index]) == bit {
			newValues = append(newValues, value)
		}
	}

	return newValues
}

func main() {
	submarine := Submarine{}
	fmt.Printf("Power consumption: %d\n", submarine.PowerConsumption(readNumbers()))
	fmt.Printf("Life support rating: %d\n", submarine.LifeSupportRating(readNumbers()))
}

func readNumbers() []string {
	bytes, _ := ioutil.ReadFile("../input.txt")
	return strings.Split(string(bytes), "\n")
}
