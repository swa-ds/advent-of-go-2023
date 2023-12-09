package day06

import (
	"fmt"
	"strings"
	"swads/aoc2023/aocutils"
)

var example = []string{
	"Time:        47     84     74     67",
	"Distance:   207   1394   1209   1014",
}

func Solve() {
	part1 := SolvePart1(example)
	fmt.Println("Day 06 - part1:", part1)
}

func SolvePart1(input []string) int {
	times := strings.Fields(input[0])[1:]
	distances := strings.Fields(input[1])[1:]
	result := 1

	for i := 0; i < len(times); i++ {
		wins := 0
		millis := aocutils.StrToInt(times[i])
		maxDist := aocutils.StrToInt(distances[i])
		for holdTime := 1; holdTime < millis; holdTime++ {
			rest := millis - holdTime
			dist := holdTime * rest
			if dist > maxDist {
				wins += 1
			}
		}
		result *= wins
	}

	return result
}
