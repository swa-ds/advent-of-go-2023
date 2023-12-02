package day02

import (
	"fmt"
	"strconv"
	"strings"
	"swads/aoc2023/fileutils"
)

var cubesMax = map[string]int{"red": 12, "green": 13, "blue": 14}

func Solve() {
	lines := fileutils.ReadLines("day02/input.txt")

	part1 := SolvePart1(lines)
	fmt.Printf("Day 02 - Part 1: %d\n", part1)
}

func SolvePart1(lines []string) int {
	result := 0
	for _, line := range lines {
		gameNr, gameStr := parseGame(line)

		sets := splitAndTrim(gameStr, ";")
		isValid := true
		for _, set := range sets {
			draws := splitAndTrim(set, ",")
			for _, draw := range draws {
				draw := splitAndTrim(draw, " ")
				cubes := strToInt(draw[0])
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

func parseGame(line string) (gameNr int, games string) {
	colonIdx := strings.Index(line, ":")
	nrStr := line[5:colonIdx]
	nr := strToInt(nrStr)
	return nr, line[colonIdx+1:]
}

func splitAndTrim(s string, delim string) []string {
	parts := strings.Split(s, delim)
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts

}

func strToInt(s string) int {
	nr, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return nr
}
