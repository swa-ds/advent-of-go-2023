package day03

import (
	"fmt"
	"swads/aoc2023/fileutils"
	"unicode"
)

type star struct {
	row, col int
}

type MultiMap map[star][]int

func (m MultiMap) add(key star, value int) {
	m[key] = append(m[key], value)
}

func (m MultiMap) get(key star) []int {
	return m[key]
}

func (m MultiMap) keys() []star {
	keys := []star{}
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

var noStar = star{-1, -1}

func Solve() {
	matrix := fileutils.ReadMatrix("day03/input.txt")

	part1 := SolvePart1(matrix)
	fmt.Printf("Day 03 - Part 1: %d\n", part1)

	part2 := SolvePart2(matrix)
	fmt.Printf("Day 03 - Part 2: %d\n", part2)
}

func SolvePart1(matrix [][]rune) int {
	result := 0
	numbers := []string{}
	for row := 0; row < len(matrix); row++ {
		withinNumber := false
		numberStart := -1
		numberEnd := -1
		for col := 0; col < len(matrix[0]); col++ {
			if unicode.IsDigit(matrix[row][col]) {
				if !withinNumber {
					withinNumber = true
					numberStart = col
				} else if col == len(matrix[row])-1 {
					// within number and reached end of row
					numberEnd = col + 1
					numberStr := string(matrix[row])[numberStart:numberEnd]
					if isValidNumber(matrix, row, numberStart, numberEnd) {
						number := fileutils.StrToInt(numberStr)
						result += number
						numbers = append(numbers, numberStr+"v")
					} else {
						numbers = append(numbers, numberStr+"i")
					}
				}
			} else {
				if withinNumber {
					withinNumber = false
					numberEnd = col
					numberStr := string(matrix[row])[numberStart:numberEnd]
					if isValidNumber(matrix, row, numberStart, numberEnd) {
						number := fileutils.StrToInt(numberStr)
						result += number
						numbers = append(numbers, numberStr+"v")
					} else {
						numbers = append(numbers, numberStr+"i")
					}
				}
			}
		}
		numbers = append(numbers, "\n")
	}
	//fmt.Println("Numbers: ", numbers)
	return result
}

func isValidNumber(matrix [][]rune, row int, numStart int, numEnd int) bool {
	//fmt.Printf("row=%d, numStart=%d, numEnd=%d\n", row, numStart, numEnd)
	// check symbol before numStart
	if numStart > 0 && matrix[row][numStart-1] != '.' {
		return true
	}
	// check symbol after numEnd
	if numEnd < len(matrix[row])-1 && matrix[row][numEnd] != '.' {
		return true
	}
	checkUpperLine := row > 0
	checkLowerLine := row < len(matrix)-1
	for col := fileutils.Max(0, numStart-1); col < fileutils.Min(numEnd+1, len(matrix[0])); col++ {
		if checkUpperLine && matrix[row-1][col] != '.' {
			return true
		}
		if checkLowerLine && matrix[row+1][col] != '.' {
			return true
		}
	}
	return false
}

func SolvePart2(matrix [][]rune) int {
	starredNumbers := make(MultiMap)
	for row := 0; row < len(matrix); row++ {
		withinNumber := false
		numberStart := -1
		numberEnd := -1
		for col := 0; col < len(matrix[0]); col++ {
			if unicode.IsDigit(matrix[row][col]) {
				if !withinNumber {
					withinNumber = true
					numberStart = col
				} else if col == len(matrix[row])-1 {
					// within number and reached end of row
					numberEnd = col + 1
					numberStr := string(matrix[row])[numberStart:numberEnd]
					adjStar := findAdjacentStar(matrix, row, numberStart, numberEnd)
					if adjStar != noStar {
						number := fileutils.StrToInt(numberStr)
						starredNumbers.add(adjStar, number)
					}
				}
			} else {
				if withinNumber {
					withinNumber = false
					numberEnd = col
					numberStr := string(matrix[row])[numberStart:numberEnd]
					adjStar := findAdjacentStar(matrix, row, numberStart, numberEnd)
					if adjStar != noStar {
						number := fileutils.StrToInt(numberStr)
						starredNumbers.add(adjStar, number)
					}
				}
			}
		}
	}
	result := 0
	for _, key := range starredNumbers.keys() {
		numbers := starredNumbers.get(key)
		if len(numbers) == 2 {
			result += numbers[0] * numbers[1]
		}
	}

	return result
}

func findAdjacentStar(matrix [][]rune, row int, numStart int, numEnd int) star {
	//fmt.Printf("row=%d, numStart=%d, numEnd=%d\n", row, numStart, numEnd)
	// check symbol before numStart
	if numStart > 0 && matrix[row][numStart-1] == '*' {
		return star{row, numStart - 1}
	}
	// check symbol after numEnd
	if numEnd < len(matrix[row])-1 && matrix[row][numEnd] == '*' {
		return star{row, numEnd}
	}
	checkUpperLine := row > 0
	checkLowerLine := row < len(matrix)-1
	for col := fileutils.Max(0, numStart-1); col < fileutils.Min(numEnd+1, len(matrix[0])); col++ {
		if checkUpperLine && matrix[row-1][col] == '*' {
			return star{row - 1, col}
		}
		if checkLowerLine && matrix[row+1][col] == '*' {
			return star{row + 1, col}
		}
	}
	return noStar
}
