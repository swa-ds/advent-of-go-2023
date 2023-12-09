package day06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var example1 = []string{
	"Time:      7  15   30",
	"Distance:  9  40  200",
}

var example2 = []string{
	"Time:      71530",
	"Distance:  940200",
}

func TestSolvePart1(t *testing.T) {
	result := SolvePart1(example1)

	assert.Equal(t, 288, result)
}

func TestSolvePart2(t *testing.T) {
	result := SolvePart2(example2)

	assert.Equal(t, 71503, result)
}
