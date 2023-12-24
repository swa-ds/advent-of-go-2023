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

func TestFindReflection(t *testing.T) {
	patterns := parsePatterns(input)
	hor := findReflection(patterns[0].pattern)

	assert.Equal(t, 4, hor)
}
