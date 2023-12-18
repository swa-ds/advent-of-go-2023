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
	for i := 0; i < math.MaxInt; i++ {
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
			return steps
		}
	}
	panic("No end in sight...")
}

func SolvePart2(input []string) int {
	directions := []rune(input[0])
	forks := parseForks(input[2:])
	positions := []string{}
	for key, _ := range forks {
		if lastChar(key) == "A" {
			positions = append(positions, key)
		}
	}
	steps := 0
	for i := 0; i < math.MaxInt; i++ {
		dirIdx := i % len(directions)
		dir := directions[dirIdx]
		steps += 1
		allZ := true
		//fmt.Println(positions)
		for posIdx, pos := range positions {
			//fmt.Println("pos", pos, "posIdx", posIdx, "dir", string(dir))
			next := forks[pos]
			//fmt.Println("next:", next)
			if dir == 'L' {
				//fmt.Println("going left")
				positions[posIdx] = next.left
			} else if dir == 'R' {
				//fmt.Println("going right")
				positions[posIdx] = next.right
			}
			if lastChar(positions[posIdx]) != "Z" {
				allZ = false
			}
		}
		if allZ {
			//fmt.Println(positions)
			return steps
		}
	}
	panic("No end in sight...")
}

func lastChar(s string) string {
	return s[len(s)-1:]
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
