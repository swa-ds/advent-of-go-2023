package day08

import (
	"fmt"
	"math"
	"regexp"
	util "swads/aoc2023/aocutils"
)

func Solve() {
	input := util.ReadLines("day08/input.txt")

	part1 := SolvePart1(input)
	fmt.Println("Day 08 - Part 1:", part1)

	part2 := SolvePart2(input)
	fmt.Println("Day 08 - Part 2:", part2)
}

func SolvePart1(input []string) int {
	directions := []rune(input[0])
	forks := parseForks(input[2:])
	position := "AAA"
	steps := 0
	found := false
	for i := 0; !found; i++ {
		idx := i % len(directions)
		dir := directions[idx]
		steps += 1
		next := forks[position]
		if dir == 'L' {
			position = next.left
		} else if dir == 'R' {
			position = next.right
		}
		if position == "ZZZ" {
			found = true
		}
	}
	return steps
}

func SolvePart2(input []string) int {
	directions := []rune(input[0])
	forks := parseForks(input[2:])
	positions := []string{}
	// store number of steps for each position until "..Z" is reached
	stepCounts := []int{}
	for key := range forks {
		if lastChar(key) == "A" {
			positions = append(positions, key)
		}
	}
	steps := 0
	// when we have a "stepsPerPosition" count for each position, we're done
	for i := 0; len(stepCounts) < len(positions); i++ {
		idx := i % len(directions)
		steps += 1
		dir := directions[idx]
		for i := 0; i < len(positions); i++ {
			next := forks[positions[i]]
			if dir == 'L' {
				positions[i] = next.left
			} else if dir == 'R' {
				positions[i] = next.right
			}
			if endsWith(positions[i], "Z") {
				stepCounts = append(stepCounts, steps)
			}
		}
	}
	// result is the least common multiple of all step counts
	return lcm(stepCounts...)
}

func endsWith(s string, suffix string) bool {
	if len(suffix) > len(s) {
		return false
	}
	testStr := s[len(s)-len(suffix):]
	return testStr == suffix
}

func lastChar(s string) string {
	return s[len(s)-1:]
}

// Calculates and returns the least common multiple
// (kleinses gemeinsames Vielfaches) of all given ints
func lcm(ints ...int) int {
	result := lcmTwo(ints[0], ints[1])
	for _, i := range ints[2:] {
		result = lcmTwo(result, i)
	}
	return result
}

// Calculates and returns the least common multiple
// (kleinses gemeinsames Vielfaches) of a and b
func lcmTwo(a, b int) int {
	return int(math.Abs(float64(a*b)) / float64(gcd(a, b)))
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func parseForks(input []string) map[string]Fork {
	forks := make(map[string]Fork)
	re := regexp.MustCompile(`(.{3}) = \((.{3}), (.{3})\)`)
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
