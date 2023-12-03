package day01

import (
	"fmt"
	"swads/aoc2023/aocutils"
	"unicode"
)

func Solve() {
	lines := aocutils.ReadLines("day01/input.txt")

	part1 := SolvePart1(lines)
	fmt.Printf("Day 01 - Part 1: %d\n", part1)

	part2 := SolvePart2(lines)
	fmt.Printf("Day 01 - Part 2: %d\n", part2)
}

func SolvePart1(input []string) int {
	result := 0
	for _, line := range input {
		first, last := firstAndLastDigit(line)
		result += first*10 + last
	}
	return result
}

func SolvePart2(input []string) int {
	result := 0
	for _, line := range input {
		first, last := firstAndLastDigitWithWords(line)
		result += first*10 + last
	}
	return result
}

func firstAndLastDigit(line string) (firstDigit int, lastDigit int) {
	runes := []rune(line)
	first := -1
	last := -1
	for _, rune := range runes {
		if unicode.IsDigit(rune) {
			num := int(rune - '0')
			if first < 0 {
				first = num
			}
			last = num
		}
	}
	return first, last
}

func firstAndLastDigitWithWords(line string) (firstDigit int, lastDigit int) {
	runes := []rune(line)
	first := -1
	last := -1
	for i := 0; i < len(runes); i++ {
		num := digitAt(runes, i)
		if num > -1 {
			if first < 0 {
				first = num
			}
			last = num
		}
	}
	return first, last
}

func digitAt(r []rune, index int) int {
	if unicode.IsDigit(r[index]) {
		return int(r[index] - '0')
	}
	if digitWordAt(r, index, "one") > -1 {
		return 1
	}
	if digitWordAt(r, index, "two") > -1 {
		return 2
	}
	if digitWordAt(r, index, "three") > -1 {
		return 3
	}
	if digitWordAt(r, index, "four") > -1 {
		return 4
	}
	if digitWordAt(r, index, "five") > -1 {
		return 5
	}
	if digitWordAt(r, index, "six") > -1 {
		return 6
	}
	if digitWordAt(r, index, "seven") > -1 {
		return 7
	}
	if digitWordAt(r, index, "eight") > -1 {
		return 8
	}
	if digitWordAt(r, index, "nine") > -1 {
		return 9
	}

	return -1
}

func digitWordAt(r []rune, index int, word string) int {
	if isSubstrAt(r, index, word) {
		return 1
	}

	return -1
}

func isSubstrAt(r []rune, begin int, str string) bool {
	if begin+len(str) > len(r) {
		return false
	}
	sub := string(r[begin : begin+len(str)])
	return sub == str
}

func substr(r []rune, begin int, len int) string {
	return string(r[begin : begin+len])
}
