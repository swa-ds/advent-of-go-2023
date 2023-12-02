package day01

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay01Part1(t *testing.T) {
	input := []string{"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet"}

	result := SolvePart1(input)

	assert.Equal(t, 142, result)
}

func TestDay02Part2(t *testing.T) {
	input := []string{"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen"}
	result := SolvePart2(input)

	assert.Equal(t, 281, result)
}

func TestFirstAndLastDigit(t *testing.T) {
	first, last := firstAndLastDigit("abc5sdf3sdf9s")

	assert.Equal(t, 5, first)
	assert.Equal(t, 9, last)
}

func TestFirstAndLastDigitWithWords(t *testing.T) {
	first, last := firstAndLastDigitWithWords("6twofive3two")
	assert.Equal(t, 6, first)
	assert.Equal(t, 2, last)
}

func TestDigitAt(t *testing.T) {
	runes := []rune("abc3deonefghij")

	sub := substr(runes, 2, 3)
	fmt.Printf("%d", len(sub))
	fmt.Println()

	assert.Equal(t, 3, digitAt(runes, 3))
	assert.Equal(t, 1, digitAt(runes, 6))
	assert.Equal(t, -1, digitAt(runes, 13))

}
