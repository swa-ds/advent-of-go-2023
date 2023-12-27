package day14

import (
	"testing"

	util "swads/aoc2023/aocutils"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"O....#....",
	"O.OO#....#",
	".....##...",
	"OO.#O....O",
	".O.....O#.",
	"O.#..O.#.#",
	"..O..#O..O",
	".......O..",
	"#....###..",
	"#OO..#....",
}

var expTilted = []string{
	"OOOO.#.O..",
	"OO..#....#",
	"OO..O##..O",
	"O..#.OO...",
	"........#.",
	"..#....#.#",
	"..O..#.O.O",
	"..O.......",
	"#....###..",
	"#....#....",
}

var expOneCycle = []string{
	".....#....",
	"....#...O#",
	"...OO##...",
	".OO#......",
	".....OOO#.",
	".O#...O#.#",
	"....O#....",
	"......OOOO",
	"#...O###..",
	"#..OO#....",
}

var expThreeCycles = []string{
	".....#....",
	"....#...O#",
	".....##...",
	"..O#......",
	".....OOO#.",
	".O#...O#.#",
	"....O#...O",
	".......OOO",
	"#...O###.O",
	"#.OOO#...O",
}

func TestSolvePart1(t *testing.T) {
	result := SolvePart1(input)

	assert.Equal(t, 136, result)
}

func TestSolvePart2(t *testing.T) {
	result := SolvePart2(input)

	assert.Equal(t, 64, result)
}

func TestSpinCycle(t *testing.T) {
	platform := util.LinesToMatrix(input)
	result := spinCycle(platform)
	exp := util.LinesToMatrix(expOneCycle)

	assert.Equal(t, exp, result)

	result = spinCycle(result)
	result = spinCycle(result)
	exp = util.LinesToMatrix(expThreeCycles)

	assert.Equal(t, exp, result)
}

func TestTiltNorth(t *testing.T) {
	platform := util.LinesToMatrix(input)
	exp := util.LinesToMatrix(expTilted)
	tilted := tiltNorth(platform)

	assert.Equal(t, exp, tilted)
	assert.NotEqual(t, exp, platform)
}

func TestRotateMatrix(t *testing.T) {
	original := [][]rune{
		{'1', '2', '3', '4'},
		{'5', '6', '7', '8'},
	}
	exp := [][]rune{
		{'5', '1'},
		{'6', '2'},
		{'7', '3'},
		{'8', '4'},
	}

	// rotate 1st time
	rotated := rotateMatrix(original)

	assert.Equal(t, exp, rotated)

	// rotate 2nd time
	exp = [][]rune{
		{'8', '7', '6', '5'},
		{'4', '3', '2', '1'},
	}
	rotated = rotateMatrix(rotated)

	assert.Equal(t, exp, rotated)

	// rotate 3rd time
	exp = [][]rune{
		{'4', '8'},
		{'3', '7'},
		{'2', '6'},
		{'1', '5'},
	}
	rotated = rotateMatrix(rotated)

	assert.Equal(t, exp, rotated)

	// rotate 4th time
	rotated = rotateMatrix(rotated)

	assert.Equal(t, original, rotated)
}

func TestEqual(t *testing.T) {
	a := util.LinesToMatrix(input)
	b := util.LinesToMatrix(input)

	assert.True(t, equal(a, b))

	b = util.LinesToMatrix(expTilted)

	assert.False(t, equal(a, b))
}

func TestIndexOf(t *testing.T) {
	encountered := []Encountered{}
	encountered = append(encountered, Encountered{util.LinesToMatrix(input)})
	encountered = append(encountered, Encountered{util.LinesToMatrix(expThreeCycles)})

	exp3 := util.LinesToMatrix(expThreeCycles)
	assert.Equal(t, 1, indexOf(encountered, exp3))

	tilted := util.LinesToMatrix(expTilted)
	assert.Equal(t, -1, indexOf(encountered, tilted))
}
