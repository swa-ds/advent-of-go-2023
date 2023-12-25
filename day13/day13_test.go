package day13

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"#.##..##.",
	"..#.##.#.",
	"##......#",
	"##......#",
	"..#.##.#.",
	"..##..##.",
	"#.#.##.#.",
	"",
	"#...##..#",
	"#....#..#",
	"..##..###",
	"#####.##.",
	"#####.##.",
	"..##..###",
	"#....#..#",
}

func TestSolvePart1(t *testing.T) {
	result := SolvePart1(input)

	assert.Equal(t, 405, result)
}

func TestParsePatterns(t *testing.T) {
	patterns := parsePatterns(input)
	assert.Equal(t, 2, len(patterns))
	fmt.Println(patterns)
}

func TestFind(t *testing.T) {
	patterns := parsePatterns(input)
	hor := find(patterns[0], 0)

	assert.Equal(t, 5, hor)

	vert := find(patterns[1], 0)

	assert.Equal(t, 400, vert)

}

func TestRowDiffs(t *testing.T) {
	diffs := diffsInRows("abcde", "abcde")
	assert.Equal(t, 0, diffs)

	diffs = diffsInRows("abcde", "abcdx")
	assert.Equal(t, 1, diffs)
}
