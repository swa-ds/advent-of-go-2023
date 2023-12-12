package day07

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"32T3K 765",
	"T55J5 684",
	"KK677 28",
	"KTJJT 220",
	"QQQJA 483",
}

func TestSolvePart1(t *testing.T) {
	result := SolvePart1(input)

	assert.Equal(t, 6440, result)
}

func TestCalculateStrength(t *testing.T) {
	hand := "23456"
	assert.Equal(t, 10_102_030_405, calculateStrength(hand))

	hand = "AKQJT"
	assert.Equal(t, 11_312_111_009, calculateStrength(hand))
}

func TestIndex(t *testing.T) {
	assert.Equal(t, 9, strings.Index(string(cardOrder), "J"))
}

func TestReverseSort(t *testing.T) {
	ints := []int{3, 9, 0, 5, 1}

	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	assert.Equal(t, []int{9, 5, 3, 1, 0}, ints)
}
