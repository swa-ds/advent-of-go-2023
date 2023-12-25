package day14

import (
	"fmt"
	util "swads/aoc2023/aocutils"
)

func Solve() {
	input := util.ReadLines("day14/input.txt")

	part1 := SolvePart1(input)
	fmt.Println("Day 14 - Part 1:", part1)

	//part2 := SolvePart2(input)
	//fmt.Println("Day 14 - Part 2:", part2)
}

func SolvePart1(input []string) int {
	platform := util.LinesToMatrix(input)
	tiltNorth(platform)
	weight := len(platform)
	result := 0
	for row := 0; row < len(platform); row++ {
		for col := 0; col < len(platform[row]); col++ {
			if (platform[row][col]) == 'O' {
				result += weight
			}
		}
		weight--
	}
	return result
}

func tiltNorth(platform [][]rune) {
	for col := 0; col < len(platform[0]); col++ {
		moved := true
		for i := 0; moved; i++ {
			moved = false
			for row := 0; row < len(platform)-1; row++ {
				if platform[row][col] == '.' && platform[row+1][col] == 'O' {
					platform[row][col] = 'O'
					platform[row+1][col] = '.'
					moved = true
				}
			}
		}
	}
}
