package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input1 = []string{
	".....",
	".S-7.",
	".|.|.",
	".L-J.",
	".....",
}

var input2 = []string{
	"7-F7-",
	".FJ|7",
	"SJLL7",
	"|F--J",
	"LJ.LJ",
}

var input1AsRunes = [][]rune{
	{'.', '.', '.', '.', '.'},
	{'.', 'S', '-', '7', '.'},
	{'.', '|', '.', '|', '.'},
	{'.', 'L', '-', 'J', '.'},
	{'.', '.', '.', '.', '.'},
}

func TestSolvePart1(t *testing.T) {
	result := SolvePart1(input1)

	assert.Equal(t, 4, result)

	result = SolvePart1(input2)
	assert.Equal(t, 8, result)
}

func TestParseField(t *testing.T) {
	field := parseField(input1)
	expected := Field{input1AsRunes, Position{1, 1}, Position{1, 1}, South}

	assert.Equal(t, expected, field)
	assert.Equal(t, 'S', field.symbolOnPos())
	assert.True(t, field.onStart())

}
