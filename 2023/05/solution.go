package pkg202305

import (
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type SmartMap struct {
	Entries []SmartEntry
}

func (m *SmartMap) Clear() {
	m.Entries = []SmartEntry{}
}

func (m *SmartMap) AddEntry(e SmartEntry) {
	m.Entries = append(m.Entries, e)
}

func (m *SmartMap) ReverseLookup(n int) int {
	for _, entry := range m.Entries {
		if entry.ReverseContains(n) {
			return entry.ReverseValue(n)
		}
	}

	return n
}

func (m *SmartMap) Lookup(n int) int {
	for _, entry := range m.Entries {
		if entry.Contains(n) {
			return entry.Value(n)
		}
	}

	return n
}

type SmartEntry struct {
	SrcRangeStart  int
	DestRangeStart int
	RangeLength    int
}

func (e *SmartEntry) Contains(n int) bool {
	if n < e.SrcRangeStart {
		return false
	}

	if n <= (e.SrcRangeStart + e.RangeLength - 1) {
		return true
	}

	return false
}

func (e *SmartEntry) ReverseContains(n int) bool {
	if n < e.DestRangeStart {
		return false
	}

	if n <= (e.DestRangeStart + e.RangeLength - 1) {
		return true
	}

	return false
}

func (e *SmartEntry) ReverseValue(n int) int {
	if !e.ReverseContains(n) {
		return n
	}

	diff := n - e.DestRangeStart
	return e.SrcRangeStart + diff
}

func (e *SmartEntry) Value(n int) int {
	if !e.Contains(n) {
		return n
	}

	diff := n - e.SrcRangeStart
	return e.DestRangeStart + diff
}

func NewSmartMap() *SmartMap {
	return &SmartMap{Entries: []SmartEntry{}}
}

var SeedToSoil = NewSmartMap()
var SoilToFertilizer = NewSmartMap()
var FertilizerToWater = NewSmartMap()
var WaterToLight = NewSmartMap()
var LightToTemperature = NewSmartMap()
var TemperatureToHumidity = NewSmartMap()
var HumidityToLocation = NewSmartMap()

var seedsRegex = regexp.MustCompile(`seeds: ([\d\s]*)`)
var seedToSoilRegex = regexp.MustCompile(`seed-to-soil map:\n([\d\s]*)`)
var soilToFertilizerRegex = regexp.MustCompile(`soil-to-fertilizer map:\n([\d\s]*)`)
var fertilizerToWaterRegex = regexp.MustCompile(`fertilizer-to-water map:\n([\d\s]*)`)
var waterToLightRegex = regexp.MustCompile(`water-to-light map:\n([\d\s]*)`)
var lightToTemperatureRegex = regexp.MustCompile(`light-to-temperature map:\n([\d\s]*)`)
var TemperatureToHumidityRegex = regexp.MustCompile(`temperature-to-humidity map:\n([\d\s]*)`)
var HumidityToLocationRegex = regexp.MustCompile(`humidity-to-location map:\n([\d\s]*)`)

func A(input string) int {
	clearMaps()
	d, _ := os.ReadFile(input)
	parse(string(d))

	lowestLocation := math.MaxInt
	for _, seed := range findSeeds(string(d)) {
		location := findLocation(seed)
		if location < lowestLocation {
			lowestLocation = location
		}
	}

	return lowestLocation
}

func B(input string) int {
	clearMaps()
	d, _ := os.ReadFile(input)
	parse(string(d))

	seedRanges := findSeedRanges(string(d))

	for min := 0; ; min++ {
		if hasSeed(min, seedRanges) {
			return min
		}
	}
}

func hasSeed(location int, seedRanges [][]int) bool {
	seed := findSeed(location)

	for _, seedRange := range seedRanges {
		if seedRange[0] <= seed && seed < seedRange[1] {
			return true
		}
	}

	return false
}

func findSeed(location int) int {
	humidity := HumidityToLocation.ReverseLookup(location)
	temperature := TemperatureToHumidity.ReverseLookup(humidity)
	light := LightToTemperature.ReverseLookup(temperature)
	water := WaterToLight.ReverseLookup(light)
	fertilizer := FertilizerToWater.ReverseLookup(water)
	soil := SoilToFertilizer.ReverseLookup(fertilizer)
	seed := SeedToSoil.ReverseLookup(soil)
	return seed
}

func findLocation(seed int) int {
	soil := SeedToSoil.Lookup(seed)
	fertilizer := SoilToFertilizer.Lookup(soil)
	water := FertilizerToWater.Lookup(fertilizer)
	light := WaterToLight.Lookup(water)
	temperature := LightToTemperature.Lookup(light)
	humidity := TemperatureToHumidity.Lookup(temperature)
	location := HumidityToLocation.Lookup(humidity)
	return location
}

func findOrSame(n int, values map[int]int) int {
	if v, ok := values[n]; ok {
		return v
	}

	return n
}

func findSeeds(input string) []int {
	match := seedsRegex.FindStringSubmatch(input)
	ns := []int{}

	for _, n := range strings.Split(match[1], " ") {
		v, _ := strconv.Atoi(strings.Trim(n, "\n"))
		ns = append(ns, v)
	}

	return ns
}

func findSeedRanges(input string) [][]int {
	seedRanges := [][]int{}
	seeds := findSeeds(input)
	for i := 0; i+1 < len(seeds); {
		start := seeds[i]
		length := seeds[i+1]
		seedRanges = append(seedRanges, []int{start, start + length})
		i += 2
	}

	return seedRanges
}

func parse(data string) {
	applyRegexes(
		data,
		[]*regexp.Regexp{
			seedToSoilRegex,
			soilToFertilizerRegex,
			fertilizerToWaterRegex,
			waterToLightRegex,
			lightToTemperatureRegex,
			TemperatureToHumidityRegex,
			HumidityToLocationRegex,
		},
		[]*SmartMap{
			SeedToSoil,
			SoilToFertilizer,
			FertilizerToWater,
			WaterToLight,
			LightToTemperature,
			TemperatureToHumidity,
			HumidityToLocation,
		},
	)
}

func applyRegexes(data string, regexes []*regexp.Regexp, maps []*SmartMap) {
	for i, regex := range regexes {
		for _, ns := range applyRegex(data, regex) {
			maps[i].AddEntry(SmartEntry{
				SrcRangeStart:  ns[1],
				DestRangeStart: ns[0],
				RangeLength:    ns[2],
			})
		}
	}
}

func applyRegex(data string, r *regexp.Regexp) [][]int {
	match := r.FindStringSubmatch(data)
	g1 := match[1]
	a := [][]int{}

	for _, line := range strings.Split(g1, "\n") {
		if strings.Trim(line, " ") != "" {
			ns := []int{}

			for _, n := range strings.Split(line, " ") {
				v, _ := strconv.Atoi(strings.Trim(n, " "))
				ns = append(ns, v)
			}

			a = append(a, ns)
		}
	}

	return a
}

func clearMaps() {
	SeedToSoil.Clear()
	SoilToFertilizer.Clear()
	FertilizerToWater.Clear()
	WaterToLight.Clear()
	LightToTemperature.Clear()
	TemperatureToHumidity.Clear()
	HumidityToLocation.Clear()
}
