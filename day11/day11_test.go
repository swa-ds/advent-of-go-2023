package day11

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"...#......",
	".......#..",
	"#.........",
	"..........",
	"......#...",
	".#........",
	".........#",
	"..........",
	".......#..",
	"#...#.....",
}

var expanded = []string{
	"....#........",
	".........#...",
	"#............",
	".............",
	".............",
	"........#....",
	".#...........",
	"............#",
	".............",
	".............",
	".........#...",
	"#....#.......",
}

func TestSolvePart1(t *testing.T) {
	part1 := SolvePart1(input)

	assert.Equal(t, 374, part1)
}

func TestSolvePart1WithPart2Method(t *testing.T) {
	part1 := SolvePart2(input, 2)

	assert.Equal(t, 374, part1)
}

func TestSolvePart2(t *testing.T) {
	part2 := SolvePart2(input, 10)

	assert.Equal(t, 1030, part2)

	part2 = SolvePart2(input, 100)
	assert.Equal(t, 8410, part2)
}

func TestExpandCoord(t *testing.T) {
	expansionRate := 2
	expandedCoords := []int{1, 5}
	galaxyCoord := 2

	expandedCoord := expandCoord(galaxyCoord, expansionRate, expandedCoords)

	assert.Equal(t, 3, expandedCoord)

	expansionRate = 10
	galaxyCoord = 7

	expandedCoord = expandCoord(galaxyCoord, expansionRate, expandedCoords)

	assert.Equal(t, 25, expandedCoord)
}

func TestExpandSpace(t *testing.T) {
	space := expandSpace(input)
	for i, row := range space {
		assert.Equal(t, expanded[i], string(row))
	}
}

func TestExpandRows(t *testing.T) {
	//space - the final frontier... (Der Weltraum: unendliche Weiten...)
	space := expandRows(input)
	// input has 10 lines; two without galaxies; therefore lines should be 10 + 2
	assert.Equal(t, 12, len(space))
}

func TestExpandCols(t *testing.T) {
	space := expandCols(input)
	for _, s := range space {
		fmt.Println(string(s))
	}
	assert.Equal(t, 13, len(space[0]))
}

func TestFindAllPairs(t *testing.T) {
	galaxies := []Galaxy{}
	for i := 1; i <= 9; i++ {
		galaxies = append(galaxies, Galaxy{i, i})
	}
	pairs := findAllPairs(galaxies)
	assert.Equal(t, 36, len(pairs))
}

func TestShortestDistance(t *testing.T) {
	a := Galaxy{6, 1}
	b := Galaxy{11, 5}
	dist := shortestDistance(a, b)

	assert.Equal(t, 9, dist)
}

func TestString(t *testing.T) {
	s := "...#."
	assert.Equal(t, '#', []rune(s)[3])

	assert.True(t, strings.Contains(s, "#"))
}
