package day04

import (
	"fmt"
	"math"
	"strings"
	util "swads/aoc2023/aocutils"

	mapset "github.com/deckarep/golang-set/v2"
)

func Solve() {
	lines := util.ReadLines("day04/input.txt")

	part1, _ := SolvePart1And2(lines)
	fmt.Printf("Day 04 - Part 1: %d\n", part1)
}

func SolvePart1And2(lines []string) (part1 int, part2 int) {
	winningNumbers := mapset.NewSet[string]()
	resultPart1 := 0
	for i, line := range lines {
		gameNr := i + 1
		util.Unused(gameNr)
		split := util.SplitAndTrim(line, ":")
		numbers := util.SplitAndTrim(split[1], "|")
		for _, winning := range strings.Fields(numbers[0]) {
			winningNumbers.Add(winning)
		}
		scratchcardNumbers := strings.Fields(numbers[1])
		winningNumbersCount := 0
		for _, cardNumber := range scratchcardNumbers {
			if winningNumbers.Contains(cardNumber) {
				winningNumbersCount += 1
			}
		}
		resultPart1 += scorePart1(winningNumbersCount)
		winningNumbers.Clear()
	}
	return resultPart1, 0
}

func scorePart1(count int) int {
	if count > 0 {
		return int(math.Pow(2, float64(count-1)))
	}
	return 0
}
