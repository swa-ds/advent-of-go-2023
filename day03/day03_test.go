package day03

import (
	"swads/aoc2023/fileutils"
	"testing"

	"github.com/stretchr/testify/assert"
)

// slightly modified example from AoC, to have a valid number at the end of the row (755)
var input = []string{
	"467..114..",
	"...*......",
	"..35..633.",
	"......#...",
	"617*......",
	".....+.58.",
	"..592.....",
	".......755",
	"...$..*...",
	".664.598..",
}

var matrix = fileutils.LinesToMatrix(input)

func TestSolvePart1(t *testing.T) {
	result := SolvePart1(matrix)

	assert.Equal(t, 4361, result)
}

func TestSolvePart2(t *testing.T) {
	result := SolvePart2(matrix)

	assert.Equal(t, 467835, result)
}

func TestIsValidNumber(t *testing.T) {
	// 58 in row 6 is invalid
	assert.False(t, isValidNumber(matrix, 5, 7, 9))
	// 35 in row 3 is valid
	assert.True(t, isValidNumber(matrix, 2, 2, 4))
	// 617 in row 5 is valid
	assert.True(t, isValidNumber(matrix, 4, 0, 3))
	// 755 in row 8 is valid (literally an edge case)
	assert.True(t, isValidNumber(matrix, 7, 7, 9))
}
