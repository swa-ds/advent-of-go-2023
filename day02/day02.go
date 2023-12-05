package day02

import (
	"fmt"
	"strings"
	"swads/aoc2023/aocutils"
)

func Solve() {
	lines := aocutils.ReadLines("day02/input.txt")

	part1 := SolvePart1(lines)
	fmt.Printf("Day 02 - Part 1: %d\n", part1)

	part2 := SolvePart2(lines)
	fmt.Printf("Day 02 - Part 2: %d\n", part2)
}

func SolvePart1(lines []string) int {
	cubesMax := map[string]int{"red": 12, "green": 13, "blue": 14}
	result := 0
	for _, line := range lines {
		gameNr, gameStr := parseGame(line)

		sets := aocutils.SplitAndTrim(gameStr, ";")
		isValid := true
		for _, set := range sets {
			draws := aocutils.SplitAndTrim(set, ",")
			for _, draw := range draws {
				draw := aocutils.SplitAndTrim(draw, " ")
				cubes := aocutils.StrToInt(draw[0])
				color := draw[1]
				if cubes > cubesMax[color] {
					isValid = false
					break
				}
			}
		}
		if isValid {
			result += gameNr
		}
	}
	return result
}

func SolvePart2(lines []string) int {
	result := 0

	for _, line := range lines {
		cubesMin := map[string]int{"red": 0, "green": 0, "blue": 0}
		_, gameStr := parseGame(line)

		sets := aocutils.SplitAndTrim(gameStr, ";")
		for _, set := range sets {
			draws := aocutils.SplitAndTrim(set, ",")
			for _, draw := range draws {
				draw := aocutils.SplitAndTrim(draw, " ")
				cubes := aocutils.StrToInt(draw[0])
				color := draw[1]
				if cubes > cubesMin[color] {
					cubesMin[color] = cubes
				}
			}
		}
		result += cubesMin["blue"] * cubesMin["red"] * cubesMin["green"]
	}
	return result
}

func parseGame(line string) (gameNr int, games string) {
	colonIdx := strings.Index(line, ":")
	nrStr := line[5:colonIdx]
	nr := aocutils.StrToInt(nrStr)
	return nr, line[colonIdx+1:]
}
