package day18

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"R 6 (#70c710)",
	"D 5 (#0dc571)",
	"L 2 (#5713f0)",
	"D 2 (#d2c081)",
	"R 2 (#59c680)",
	"D 2 (#411b91)",
	"L 5 (#8ceee2)",
	"U 2 (#caa173)",
	"L 1 (#1b58a2)",
	"U 2 (#caa171)",
	"R 2 (#7807d2)",
	"U 3 (#a77fa3)",
	"L 2 (#015232)",
	"U 2 (#7a21e3)",
}

func TestSolvePart1(t *testing.T) {
	part1 := SolvePart1(input)

	assert.Equal(t, 62, part1)
}

func TestSolvePart2(t *testing.T) {
	part2 := SolvePart2(input)

	assert.Equal(t, 952408144115, part2)
}

func TestHexToInstruction(t *testing.T) {
	hexcode := "70c710"
	exp := Instruction{"R", 461937}

	inst := hexToInstruction(hexcode)

	assert.Equal(t, exp, inst)

	hexcode = "8ceee2"
	exp = Instruction{"L", 577262}

	inst = hexToInstruction(hexcode)

	assert.Equal(t, exp, inst)
}
