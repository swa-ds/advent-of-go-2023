package day04

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
	"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
	"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
	"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
	"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
	"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
}

func TestSolve1(t *testing.T) {
	result := SolvePart1(input)

	assert.Equal(t, 13, result)
}

func TestFields(t *testing.T) {
	f := strings.Fields(input[0])

	assert.Equal(t, 16, len(f))
	assert.Equal(t, "Card", f[0])
	assert.Equal(t, "1:", f[1])
	assert.Equal(t, "|", f[7])
	assert.Equal(t, "53", f[15])

}

func TestXxx(t *testing.T) {

}
