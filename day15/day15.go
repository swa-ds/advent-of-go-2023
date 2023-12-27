package day15

import (
	"fmt"
	"strings"
	"swads/aoc2023/aocutils"
)

func Solve() {
	input := aocutils.ReadFile("day15/input.txt")

	part1 := SolvePart1(input)
	fmt.Println("Day 15 - Part 1:", part1)

}

func SolvePart1(input string) int {
	steps := strings.Split(input, ",")
	result := 0
	for _, step := range steps {
		result += hash(step)
	}
	return result
}

func hash(s string) int {
	hash := 0
	for _, r := range []rune(s) {
		hash += int(r)
		hash *= 17
		hash = hash % 256
	}
	return hash
}
