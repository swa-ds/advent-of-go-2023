package day06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"Time:      7  15   30",
	"Distance:  9  40  200",
}

func TestSolvePart1(t *testing.T) {
	result := SolvePart1(input)

	assert.Equal(t, 288, result)
}
