package day16

import (
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var raw = `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

var input = strings.Split(raw, "\n")

func TestSolvePart1(t *testing.T) {
	part1 := SolvePart1(input)

	assert.Equal(t, 46, part1)
}

func TestSolvePart2(t *testing.T) {
	part1 := SolvePart2(input)

	assert.Equal(t, 51, part1)
}

func TestMove(t *testing.T) {
	beam := NewBeam(Right, Position{1, 1})
	exp := Position{2, 1}

	beam.move()

	assert.Equal(t, exp, beam.pos)
}

func TestMax(t *testing.T) {
	ints := []int{2, 9, 17, 42, 15, 11}

	assert.Equal(t, 42, slices.Max(ints))
}
