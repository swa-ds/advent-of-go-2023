package day15

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"

func TestSolvePart1(t *testing.T) {
	part1 := SolvePart1(input)

	assert.Equal(t, 1320, part1)
}

func TestSolvePart2(t *testing.T) {
	part2 := SolvePart2(input)

	assert.Equal(t, 145, part2)
}

func TestHash(t *testing.T) {
	h := hash("HASH")
	assert.Equal(t, 52, h)

	h = hash("rn=1")
	assert.Equal(t, 30, h)
}

func TestXyz(t *testing.T) {
	fmt.Println(hash("rn"))
	fmt.Println(hash("qp"))
	fmt.Println(hash("pc"))
	fmt.Println(hash("ot"))
}
