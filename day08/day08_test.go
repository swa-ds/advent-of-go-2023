package day08

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input1a = []string{
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

var input1b = []string{
	"LLR",
	"",
	"AAA = (BBB, BBB)",
	"BBB = (AAA, ZZZ)",
	"ZZZ = (ZZZ, ZZZ)",
}

var input2 = []string{
	"LR",
	"",
	"11A = (11B, XXX)",
	"11B = (XXX, 11Z)",
	"11Z = (11B, XXX)",
	"22A = (22B, XXX)",
	"22B = (22C, 22C)",
	"22C = (22Z, 22Z)",
	"22Z = (22B, 22B)",
	"XXX = (XXX, XXX)",
}

func TestSolvePart1(t *testing.T) {
	result := SolvePart1(input1a)
	assert.Equal(t, 2, result)

	result = SolvePart1(input1b)
	assert.Equal(t, 6, result)
}

func TestSolvePart2(t *testing.T) {
	result := SolvePart2(input2)
	assert.Equal(t, 6, result)
}

func TestLastChar(t *testing.T) {
	s := "ABC"
	assert.Equal(t, "C", lastChar(s))
}

func TestEndsWith(t *testing.T) {
	assert.Equal(t, true, endsWith("ABZ", "Z"))
	assert.Equal(t, true, endsWith("ABZ", "BZ"))
	assert.Equal(t, true, endsWith("ABZ", "ABZ"))
	assert.Equal(t, false, endsWith("ABZ", "X"))
	assert.Equal(t, false, endsWith("A", "XYZ"))
}

func TestLcm(t *testing.T) {
	assert.Equal(t, 6, lcmTwo(2, 3))
}

func TestLcmMultiple(t *testing.T) {
	assert.Equal(t, 90, lcm(2, 5, 9))
}

func TestMap(t *testing.T) {
	m := make(map[string]int)
	assert.Equal(t, 0, m["x"])
}
