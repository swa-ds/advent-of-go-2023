package day09

import (
	"fmt"
	util "swads/aoc2023/aocutils"
)

func Solve() {
	input := util.ReadLines("day09/input.txt")

	part1 := SolvePart1(input)
	fmt.Println("Day 09 - Part 1:", part1)
}

func SolvePart1(lines []string) int {
	result := 0
	for _, line := range lines {
		//fmt.Println("line=", line)
		next := PredictNextValue(line)
		//fmt.Println("next=", next)
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
		//fmt.Println("i=", i, "lastValOfCurrentRow=", lastValOfCurrentRow,
		//	"lastValOfLowerRow=", lastValOfCurrentRow, "valToAdd=", valToAdd)
		lines[i] = append(lines[i], valToAdd)
	}
	//fmt.Println(lines)
	return lines[0][len(lines[0])-1]
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
