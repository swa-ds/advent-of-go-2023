package day09

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"0 3 6 9 12 15",
	"1 3 6 10 15 21",
	"10 13 16 21 30 45",
}

func TestSolvePart1(t *testing.T) {
	result := SolvePart1(input)

	assert.Equal(t, 114, result)
}

func TestPredictNextValue(t *testing.T) {
	line := ("0 3 6 9 12 15")
	nextValue := PredictNextValue(line)

	assert.Equal(t, 18, nextValue)
}

func TestGenerateNextSequence(t *testing.T) {
	sequence := []int{0, 3, 6, 9, 12, 15}
	expected := []int{3, 3, 3, 3, 3}
	next, allZeroes := generateNextSequence(sequence)

	assert.Equal(t, expected, next)
	assert.Equal(t, false, allZeroes)

	sequence = []int{3, 3, 3, 3, 3}
	expected = []int{0, 0, 0, 0}
	next, allZeroes = generateNextSequence(sequence)

	assert.Equal(t, expected, next)
	assert.Equal(t, true, allZeroes)
}

func TestAppendElementInMatrix(t *testing.T) {
	// Create a two-dimensional int slice
	twoD := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	expected := [][]int{
		{1, 2, 3},
		{4, 5, 6, 7},
	}

	// Append new value to the inner slice of the last entry
	idxLastElem := len(twoD) - 1
	twoD[idxLastElem] = append(twoD[idxLastElem], 7)

	assert.Equal(t, expected, twoD)
}
