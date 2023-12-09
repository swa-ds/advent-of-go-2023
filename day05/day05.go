package day05

import (
	"fmt"
	"math"
	"strings"
	util "swads/aoc2023/aocutils"
)

func Solve() {
	input := util.ReadFile("day05/input.txt")
	part1 := SolvePart1(input)
	fmt.Println("Day 05 - Part 1:", part1)
	part2 := SolvePart2(input)
	fmt.Println("Day 05 - Part 2:", part2)
}

func SolvePart1(input string) int {
	parts := util.SplitAndTrim(input, "\n\n")
	seeds := parseSeeds1(parts[0])
	seedToSoil := parseMap(parts[1])
	soilToFertilizer := parseMap(parts[2])
	fertilizerToWater := parseMap(parts[3])
	waterToLight := parseMap(parts[4])
	lightToTemp := parseMap(parts[5])
	tempToHumidity := parseMap(parts[6])
	humidityToLocation := parseMap(parts[7])

	min := math.MaxInt
	for _, seed := range seeds {
		conv := seedToSoil.mappedValue(seed)
		conv = soilToFertilizer.mappedValue(conv)
		conv = fertilizerToWater.mappedValue(conv)
		conv = waterToLight.mappedValue(conv)
		conv = lightToTemp.mappedValue(conv)
		conv = tempToHumidity.mappedValue(conv)
		conv = humidityToLocation.mappedValue(conv)
		min = util.Min(min, conv)
	}
	return min
}

func SolvePart2(input string) int {
	parts := util.SplitAndTrim(input, "\n\n")
	seeds := parseSeeds2(parts[0])
	seedToSoil := parseMap(parts[1])
	soilToFertilizer := parseMap(parts[2])
	fertilizerToWater := parseMap(parts[3])
	waterToLight := parseMap(parts[4])
	lightToTemp := parseMap(parts[5])
	tempToHumidity := parseMap(parts[6])
	humidityToLocation := parseMap(parts[7])

	min := math.MaxInt
	for _, seed := range seeds {
		for i := seed.Start; i < seed.End; i++ {
			conv := seedToSoil.mappedValue(i)
			conv = soilToFertilizer.mappedValue(conv)
			conv = fertilizerToWater.mappedValue(conv)
			conv = waterToLight.mappedValue(conv)
			conv = lightToTemp.mappedValue(conv)
			conv = tempToHumidity.mappedValue(conv)
			conv = humidityToLocation.mappedValue(conv)
			min = util.Min(min, conv)
		}
	}
	return min
}

func parseSeeds1(line string) []int {
	seeds := []int{}
	for _, seed := range strings.Fields(line)[1:] {
		seeds = append(seeds, util.StrToInt(seed))
	}
	return seeds
}

func parseSeeds2(line string) []SeedRange {
	ranges := []SeedRange{}
	fields := strings.Fields(line)[1:]
	for i := 0; i < len(fields); i += 2 {
		s := util.StrToInt(fields[i])
		e := s + util.StrToInt(fields[i+1])
		ranges = append(ranges, SeedRange{s, e})
	}
	return ranges
}

type SeedRange struct {
	Start int
	End   int
}

func parseMap(mapInput string) Map {
	m := NewMap()
	lines := util.SplitAndTrim(mapInput, "\n")
	for _, line := range lines[1:] {
		fields := strings.Fields(line)
		r := NewRange(fields[1], fields[0], fields[2])
		m.addRange(r)
	}
	return m
}

type Map struct {
	Ranges []Range
}

func NewMap() Map {
	return Map{
		Ranges: []Range{},
	}
}

func (m *Map) addRange(r Range) {
	m.Ranges = append(m.Ranges, r)
}

func (m Map) mappedValue(value int) int {
	for _, rng := range m.Ranges {
		if rng.inRange(value) {
			return value + rng.Diff
		}
	}
	return value
}

type Range struct {
	Start int
	End   int
	Diff  int
}

func (rng Range) inRange(value int) bool {
	return value >= rng.Start && value <= rng.End
}

func NewRange(sourceStartStr string, destStartStr string, rangeLenStr string) Range {
	sourceStart := util.StrToInt(sourceStartStr)
	destStart := util.StrToInt(destStartStr)
	rangeLen := util.StrToInt(rangeLenStr)

	end := sourceStart + rangeLen - 1
	diff := destStart - sourceStart

	return Range{
		sourceStart,
		end,
		diff,
	}

}
