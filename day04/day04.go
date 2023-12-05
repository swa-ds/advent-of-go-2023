package day04

import (
	"fmt"
	"strings"
	util "swads/aoc2023/aocutils"

	mapset "github.com/deckarep/golang-set/v2"
)

func Solve() {
	lines := util.ReadLines("day04/input.txt")

	part1 := SolvePart1(lines)
	fmt.Printf("Day 04 - Part 1: %d\n", part1)
}

func SolvePart1(lines []string) int {
	winningNumbers := mapset.NewSet[string]()
	result := 0
	for _, line := range lines {
		split := util.SplitAndTrim(line, ":")
		numbers := util.SplitAndTrim(split[1], "|")
		for _, winning := range strings.Fields(numbers[0]) {
			winningNumbers.Add(winning)
		}
		scratchcardNumbers := strings.Fields(numbers[1])
		points := 0
		for _, cardNumber := range scratchcardNumbers {
			if winningNumbers.Contains(cardNumber) {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
		result += points
		winningNumbers.Clear()
	}
	return result
}
