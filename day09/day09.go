package day09

import (
	"fmt"
	util "swads/aoc2023/aocutils"
)

func Solve() {
	input := util.ReadLines("day09/input.txt")

	part1 := SolvePart1(input)
	fmt.Println("Day 09 - Part 1:", part1)

	part2 := SolvePart2(input)
	fmt.Println("Day 09 - Part 2:", part2)
}

func SolvePart1(lines []string) int {
	result := 0
	for _, line := range lines {
		next := PredictNextValue(line)
		result += next
	}
	return result
}

func SolvePart2(lines []string) int {
	result := 0
	for _, line := range lines {
		next := PredictPreviousValue(line)
		result += next
	}
	return result
}

func PredictNextValue(line string) int {
	sequence := util.StringsToNumbers(line)
	lines := [][]int{sequence}
	allZeroes := false
	for i := 0; !allZeroes; i++ {
		sequence, allZeroes = generateNextSequence(sequence)
		lines = append(lines, sequence)
	}
	// index of last line in matrix
	idxLastElem := len(lines) - 1
	lines[idxLastElem] = append(lines[idxLastElem], 0)
	for i := idxLastElem - 1; i >= 0; i-- {
		lastValOfCurrentRow := lines[i][len(lines[i])-1]
		lastValOfLowerRow := lines[i+1][len(lines[i])-1]
		valToAdd := lastValOfCurrentRow + lastValOfLowerRow
		lines[i] = append(lines[i], valToAdd)
	}
	//fmt.Println(lines)
	return lines[0][len(lines[0])-1]
}

func PredictPreviousValue(line string) int {
	sequence := util.StringsToNumbers(line)
	lines := [][]int{sequence}
	allZeroes := false
	for i := 0; !allZeroes; i++ {
		sequence, allZeroes = generateNextSequence(sequence)
		lines = append(lines, sequence)
	}
	// index of last line in matrix
	idxLastElem := len(lines) - 1
	lines[idxLastElem] = append([]int{0}, lines[idxLastElem]...)
	for i := idxLastElem - 1; i >= 0; i-- {
		firstValOfCurrentRow := lines[i][0]
		firstValOfLowerRow := lines[i+1][0]
		valToAdd := firstValOfCurrentRow - firstValOfLowerRow
		lines[i] = append([]int{valToAdd}, lines[i]...)
	}
	//fmt.Println("lines", lines)
	return lines[0][0]
}

func generateNextSequence(sequence []int) (nextSequence []int, allZeroes bool) {
	allZeroes = true
	for i := 0; i < len(sequence)-1; i++ {
		val := sequence[i+1] - sequence[i]
		nextSequence = append(nextSequence, val)
		if val != 0 {
			allZeroes = false
		}
	}
	return
}
