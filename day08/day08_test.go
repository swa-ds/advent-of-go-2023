package day08

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input1 = []string{
	"RL",
	"",
	"AAA = (BBB, CCC)",
	"BBB = (DDD, EEE)",
	"CCC = (ZZZ, GGG)",
	"DDD = (DDD, DDD)",
	"EEE = (EEE, EEE)",
	"GGG = (GGG, GGG)",
	"ZZZ = (ZZZ, ZZZ)",
}

var input2 = []string{
	"LLR",
	"",
	"AAA = (BBB, BBB)",
	"BBB = (AAA, ZZZ)",
	"ZZZ = (ZZZ, ZZZ)",
}

func TestSolvePart1(t *testing.T) {
	result := SolvePart1(input1)
	assert.Equal(t, 2, result)

	result = SolvePart1(input2)
	assert.Equal(t, 6, result)
}
