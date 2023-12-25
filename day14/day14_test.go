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

func TestSolvePart1(t *testing.T) {
	result := SolvePart1(input)

	assert.Equal(t, 136, result)
}

func TestTiltNorth(t *testing.T) {
	platform := util.LinesToMatrix(input)
	exp := util.LinesToMatrix(expTilted)
	tiltNorth(platform)

	assert.Equal(t, exp, platform)

}
