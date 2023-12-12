package day07

import (
	"fmt"
	"sort"
	"strings"
	"swads/aoc2023/aocutils"
)

var cardOrder = "23456789TJQKA"

func Solve() {
	input := aocutils.ReadLines("day07/input.txt")
	part1 := SolvePart1(input)
	fmt.Println("Day 07 - Part 1:", part1)
}

func SolvePart1(input []string) int {
	hands := []Hand{}
	for _, hand := range input {
		hands = append(hands, NewHand(hand))
	}
	sort.Sort(ByStrength(hands))
	//fmt.Println(hands)
	result := 0
	for i, hand := range hands {
		result += (i + 1) * hand.bid
	}
	return result
}

type Hand struct {
	hand     string
	bid      int
	strength int
}

func NewHand(input string) Hand {
	f := strings.Fields(input)
	hand := f[0]
	bid := aocutils.StrToInt(f[1])
	strength := calculateStrength(hand)
	return Hand{
		hand,
		bid,
		strength,
	}
}

func calculateStrength(hand string) int {
	//fmt.Println("Calulating Hand:", hand, "frequencies are:", frequencies)
	strength := 0
	factor := 100_000_000
	for idx := 0; idx < len(hand); idx++ {
		cardValue := strings.Index(cardOrder, hand[idx:idx+1]) + 1
		strength += cardValue * factor
		factor /= 100
	}
	frequencies := frequencies(hand)
	sort.Sort(sort.Reverse(sort.IntSlice(frequencies)))
	if frequencies[0] == 5 {
		//fmt.Println("five of a kind")
		strength += 70_000_000_000
	} else if frequencies[0] == 4 {
		//fmt.Println("four of a kind")
		strength += 60_000_000_000
	} else if frequencies[0] == 3 && frequencies[1] == 2 {
		//fmt.Println("full house")
		strength += 50_000_000_000
	} else if frequencies[0] == 3 {
		//fmt.Println("three of a kind")
		strength += 40_000_000_000
	} else if frequencies[0] == 2 && frequencies[1] == 2 {
		//fmt.Println("two pair")
		strength += 30_000_000_000
	} else if frequencies[0] == 2 {
		//fmt.Println("one pair")
		strength += 20_000_000_000
	} else {
		strength += 10_000_000_000
	}
	//fmt.Println("Strength is: ", strength)
	return strength
}

func frequencies(hand string) []int {
	freqMap := make(map[rune]int)
	for _, char := range hand {
		freqMap[char] = freqMap[char] + 1
	}
	freq := []int{}
	for _, count := range freqMap {
		freq = append(freq, count)
	}
	return freq
}

type ByStrength []Hand

func (a ByStrength) Len() int {
	return len(a)
}

func (a ByStrength) Less(i, j int) bool {
	return a[i].strength < a[j].strength
}

func (a ByStrength) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
