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

	part1, part2 := SolvePart1And2(lines)
	fmt.Printf("Day 04 - Part 1: %d\n", part1)
	fmt.Printf("Day 04 - Part 2: %d\n", part2)

}

func SolvePart1And2(lines []string) (part1 int, part2 int) {
	winningNumbers := mapset.NewSet[string]()
	cardInstances := map[int]int{}
	allCards := len(lines)
	for i := 1; i <= allCards; i++ {
		cardInstances[i] = 1
	}
	resultPart1 := 0
	for i, line := range lines {
		gameNr := i + 1
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
		for i := gameNr + 1; i <= util.Min(len(lines), gameNr+winningNumbersCount); i++ {
			cardInstances[i] += cardInstances[gameNr]
		}
		resultPart1 += scorePart1(winningNumbersCount)
		winningNumbers.Clear()
	}
	resultPart2 := 0
	for _, instaceCount := range cardInstances {
		resultPart2 += instaceCount
	}
	return resultPart1, resultPart2
}

func scorePart1(count int) int {
	if count > 0 {
		return int(math.Pow(2, float64(count-1)))
	}
	return 0
}
