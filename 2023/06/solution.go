package pkg202306

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

var timeRegex = regexp.MustCompile(`Time:([\d\s]*)`)
var distanceRegex = regexp.MustCompile(`Distance:([\d\s]*)`)

type Race struct {
	Time     int
	Distance int
}

func (r *Race) WinningWays() int {
	v := 0
	for i := 0; i < r.Time; i++ {
		timeLeft := r.Time - i
		distance := i * timeLeft
		if distance > r.Distance {
			v += 1
		}
	}

	return v
}

func A(input string) int {
	data, _ := os.ReadFile(input)
	races := parseRaces(string(data))

	result := 1
	for _, race := range races {
		result *= race.WinningWays()
	}

	return result
}

func B(input string) int {
	data, _ := os.ReadFile(input)
	race := parseSingleRace(string(data))
	return race.WinningWays()
}

func parseSingleRace(input string) Race {
	input = strings.ReplaceAll(input, " ", "")
	timesMatch := timeRegex.FindStringSubmatch(input)
	times := toIntArray(timesMatch[1])

	distancesMatch := distanceRegex.FindStringSubmatch(input)
	distances := toIntArray(distancesMatch[1])

	if len(times) == 1 && len(distances) == 1 {
		return Race{Time: times[0], Distance: distances[0]}
	}

	panic("should not happen")
}

func parseRaces(input string) []Race {
	timesMatch := timeRegex.FindStringSubmatch(input)
	times := toIntArray(timesMatch[1])

	distancesMatch := distanceRegex.FindStringSubmatch(input)
	distances := toIntArray(distancesMatch[1])

	races := []Race{}
	for i, time := range times {
		races = append(races, Race{
			Time:     time,
			Distance: distances[i],
		})
	}

	return races
}

func toIntArray(list string) []int {
	a := []int{}
	for _, n := range strings.Split(list, " ") {
		x := strings.Trim(strings.Trim(n, " "), "\n")
		if x != "" {
			v, _ := strconv.Atoi(x)
			a = append(a, v)
		}
	}

	return a
}
