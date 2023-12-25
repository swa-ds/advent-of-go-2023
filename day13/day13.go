package day13

import (
	"fmt"
	"strings"
	util "swads/aoc2023/aocutils"
)

func Solve() {
	input := util.ReadLines("day13/input.txt")

	part1 := SolvePart1(input)
	fmt.Println("Day 13 - Part 1:", part1)

	//part2 := SolvePart2(input)
	//fmt.Println("Day 13 - Part 2:", part2)
}
func SolvePart1(input []string) int {
	patterns := parsePatterns(input)
	result := 0
	for _, p := range patterns {
		result += find(p, 0)
	}
	return result
}

func rotateMatrix(input []string) (rotated []string) {
	for col := 0; col < len(input[0]); col++ {
		var newRow strings.Builder
		for row := 0; row < len(input); row++ {
			newRow.WriteString(input[row][col : col+1])
		}
		rotated = append(rotated, newRow.String())
	}
	return
}

func parsePatterns(input []string) (patterns []Pattern) {
	pattern := []string{}
	for _, line := range input {
		if line == "" {
			patterns = append(patterns, Pattern{pattern})
			pattern = []string{}
		} else {
			pattern = append(pattern, line)
		}
	}
	patterns = append(patterns, Pattern{pattern})
	return
}

func find(p Pattern, maxDiffs int) int {
	hor := findReflection(p.pattern, maxDiffs) * 100
	rotated := rotateMatrix(p.pattern)
	vert := findReflection(rotated, maxDiffs)
	return hor + vert
}

func findReflection(input []string, maxDiffs int) int {
	for i := 0; i < len(input)-1; i++ {
		if diffsInRows(input[i], input[i+1]) <= maxDiffs {
			//fmt.Println("Reflection candidate at row", i)
			if isReflectionAt(input, i, maxDiffs) {
				return i + 1
			}
		}
	}
	return 0
}

func diffsInRows(s1, s2 string) int {
	diffs := 0
	for i := 0; i < len(s1); i++ {
		if s1[i:i+1] != s2[i:i+1] {
			diffs++
		}
	}
	return diffs
}

func isReflectionAt(input []string, row int, requiredDiffs int) bool {
	diffs := 0
	for a, b := row, row+1; a >= 0 && b < len(input); a, b = a-1, b+1 {
		diffs += diffsInRows(input[a], input[b])
	}
	return diffs == requiredDiffs
}

type Pattern struct {
	pattern []string
}
