package day16

import (
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

func TestMoveRight(t *testing.T) {
	beam := NewBeam(Right, Position{1, 1})
	exp := Position{2, 1}

	moveRight(&beam)

	assert.Equal(t, exp, beam.pos)
}
