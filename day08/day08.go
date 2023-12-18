package day08

import (
	"fmt"
	"math"
	"regexp"
	"swads/aoc2023/aocutils"
)

func Solve() {
	input := aocutils.ReadLines("day08/input.txt")

	part1 := SolvePart1(input)
	fmt.Println("Day 08 - Part 1:", part1)
}

func SolvePart1(input []string) int {
	directions := []rune(input[0])
	forks := parseForks(input[2:])
	position := "AAA"
	steps := 0
	for i := 0; i < math.MaxInt; i++ {
		idx := i % len(directions)
		r := directions[idx]
		steps += 1
		next := forks[position]
		if r == 'L' {
			position = next.left
		} else if r == 'R' {
			position = next.right
		}
		if position == "ZZZ" {
			return steps
		}
	}
	panic("No end in sight...")
}

func parseForks(input []string) map[string]Fork {
	forks := make(map[string]Fork)
	re := regexp.MustCompile(`([A-Z]{3}) = \(([A-Z]{3}), ([A-Z]{3})\)`)
	for _, line := range input {
		matches := re.FindStringSubmatch(line)
		key := matches[1]
		fork := Fork{
			matches[2],
			matches[3],
		}
		forks[key] = fork
	}
	return forks
}

type Fork = struct {
	left  string
	right string
}
