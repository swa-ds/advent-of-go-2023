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
		result += find(p)
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

func find(p Pattern) int {
	hor := findReflection(p.pattern) * 100
	rotated := rotateMatrix(p.pattern)
	vert := findReflection(rotated)
	return hor + vert
}

func findReflection(input []string) int {
	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			//fmt.Println("Reflection candidate at row", i)
			if isReflectionAt(input, i) {
				return i + 1
			}
		}
	}
	return 0
}

func isReflectionAt(input []string, i int) bool {
	for a, b := i-1, i+2; a >= 0 && b < len(input); a, b = a-1, b+1 {
		if input[a] != input[b] {
			return false
		}
	}
	return true
}

type Pattern struct {
	pattern []string
}
