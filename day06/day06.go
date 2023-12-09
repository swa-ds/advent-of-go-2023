package day06

import (
	"fmt"
	"strings"
	"swads/aoc2023/aocutils"
)

var input1 = []string{
	"Time:        47     84     74     67",
	"Distance:   207   1394   1209   1014",
}

var input2 = []string{
	"Time:              47847467",
	"Distance:   207139412091014",
}

func Solve() {
	part1 := SolvePart1(input1)
	fmt.Println("Day 06 - part1:", part1)

	part2 := SolvePart1(input2)
	fmt.Println("Day 06 - part2:", part2)
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

func SolvePart2(input []string) int {
	time := strings.Fields(input[0])[1]
	distance := strings.Fields(input[1])[1]
	millis := aocutils.StrToInt(time)
	maxDist := aocutils.StrToInt(distance)

	minWinning := 0
	maxWinning := millis

	for holdTime := 1; holdTime < millis; holdTime++ {
		rest := millis - holdTime
		dist := holdTime * rest
		if dist > maxDist {
			minWinning = holdTime
			println("minWinning:", minWinning)
			break
		}
	}
	for holdTime := millis; holdTime > 0; holdTime-- {
		rest := millis - holdTime
		dist := holdTime * rest
		if dist > maxDist {
			maxWinning = holdTime
			println("maxWinning:", maxWinning)
			break
		}
	}
	return maxWinning - minWinning + 1
}
