package day15

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	input := "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"
	part1 := SolvePart1(input)

	assert.Equal(t, 1320, part1)
}

func TestHash(t *testing.T) {
	h := hash("HASH")
	assert.Equal(t, 52, h)

	h = hash("rn=1")
	assert.Equal(t, 30, h)
}
