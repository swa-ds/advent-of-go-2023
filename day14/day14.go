package day14

import (
	"fmt"
	"reflect"
	util "swads/aoc2023/aocutils"
)

func Solve() {
	input := util.ReadLines("day14/input.txt")

	part1 := SolvePart1(input)
	fmt.Println("Day 14 - Part 1:", part1)

	part2 := SolvePart2(input)
	fmt.Println("Day 14 - Part 2:", part2)
}

func SolvePart1(input []string) int {
	platform := util.LinesToMatrix(input)
	platform = tiltNorth(platform)

	return weightOnNorthBeam(platform)
}

func SolvePart2(input []string) int {
	encountered := []Encountered{}
	platform := util.LinesToMatrix(input)
	var loopLen int
	var initial int

	for i := 1; true; i++ {
		//fmt.Println("Cycle", i)
		platform = spinCycle(platform)
		//printMatrix(platform)
		idxEnc := indexOf(encountered, platform)
		if idxEnc >= 0 {
			//fmt.Printf("Detected loop after %d spin cycles, "+
			//	"pattern first encountered at position %d\n", i, idxEnc)
			loopLen = len(encountered) - idxEnc
			initial = i - loopLen
			//fmt.Println("Loop length:", loopLen, "initial:", initial)
			break
		}
		encountered = append(encountered, Encountered{platform})
	}

	//fmt.Println("saved patterns: ", len(encountered))

	loops := (1_000_000_000 - initial) / loopLen
	rest := 1_000_000_000 - loopLen*loops - initial

	//fmt.Println("initial:", initial, "loops:", loops, "rest:", rest)
	//for i, enc := range encountered {
	//	fmt.Println("i:", i, "weight:", weightOnNorthBeam(enc.pattern))
	//}

	platform = encountered[initial+rest-1].pattern

	return weightOnNorthBeam(platform)
}

func indexOf(encountered []Encountered, platform [][]rune) int {
	for i, enc := range encountered {
		if equal(platform, enc.pattern) {
			return i
		}
	}
	return -1
}

func weightOnNorthBeam(original [][]rune) int {
	result := 0
	weight := len(original)
	for row := 0; row < len(original); row++ {
		for col := 0; col < len(original[row]); col++ {
			if (original[row][col]) == 'O' {
				result += weight
			}
		}
		weight--
	}
	return result
}

func spinCycle(original [][]rune) [][]rune {
	result := original
	for i := 0; i < 4; i++ {
		result = tiltNorth(result)
		result = rotateMatrix(result)
	}

	return result
}

func tiltNorth(original [][]rune) [][]rune {
	rows := len(original)
	cols := len(original[0])

	tilted := clone(original)
	for col := 0; col < cols; col++ {
		moved := true
		for i := 0; moved; i++ {
			moved = false
			for row := 0; row < rows-1; row++ {
				if tilted[row][col] == '.' && tilted[row+1][col] == 'O' {
					tilted[row][col] = 'O'
					tilted[row+1][col] = '.'
					moved = true
				}
			}
		}
	}
	return tilted
}

func clone(original [][]rune) [][]rune {
	rows := len(original)
	cols := len(original[0])

	copy := make([][]rune, rows)
	for i := 0; i < len(copy); i++ {
		copy[i] = make([]rune, cols)
		for j := 0; j < len(copy[i]); j++ {
			copy[i][j] = original[i][j]
		}
	}
	return copy
}

// function kindly provided by ChatGPT ;-)
func rotateMatrix(original [][]rune) [][]rune {
	rows := len(original)
	cols := len(original[0])
	// Create a new matrix with swapped dimensions
	rotated := make([][]rune, cols)
	for i := range rotated {
		rotated[i] = make([]rune, rows)
	}

	// Transpose the original matrix and reverse the order of rows
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			rotated[col][rows-row-1] = original[row][col]
		}
	}

	return rotated
}

func printMatrix(afterCycle [][]rune) {
	for _, row := range afterCycle {
		fmt.Println(string(row))
	}
	fmt.Println("")
}

func equal(a [][]rune, b [][]rune) bool {
	return reflect.DeepEqual(a, b)
}

type Encountered struct {
	pattern [][]rune
}
