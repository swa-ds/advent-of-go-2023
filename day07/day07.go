package day07

import (
	"fmt"
	"sort"
	"strings"
	"swads/aoc2023/aocutils"
)

const (
	part1 = iota
	part2
)

var cardOrderPart1 = "23456789TJQKA"
var cardOrderPart2 = "J23456789TQKA"

func cardValuePart1(r rune) int {
	return strings.Index(cardOrderPart1, string(r)) + 1
}

func Solve() {
	input := aocutils.ReadLines("day07/input.txt")
	part1 := SolvePart1(input)
	fmt.Println("Day 07 - Part 1:", part1)
}

func SolvePart1(input []string) int {
	hands := []Hand{}
	for _, hand := range input {
		hands = append(hands, NewHand(hand, part1))
	}
	sort.Sort(ByStrength(hands))
	//fmt.Println(hands)
	result := calculateResult(hands)
	return result
}

func calculateResult(hands []Hand) int {
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

func NewHand(input string, part int) Hand {
	f := strings.Fields(input)
	hand := f[0]
	bid := aocutils.StrToInt(f[1])
	var strength int
	if part == part1 {
		strength = calculateStrengthPart1(hand)
	} else if part == part2 {
		strength = calculateStrengthPart2(hand)
	}
	return Hand{
		hand,
		bid,
		strength,
	}
}

func calculateStrengthPart1(hand string) int {
	//fmt.Println("Calulating Hand:", hand, "frequencies are:", frequencies)
	strength := 0
	factor := 100_000_000
	handRunes := []rune(hand)
	for idx := 0; idx < len(hand); idx++ {
		cardValue := cardValuePart1(handRunes[idx])
		strength += cardValue * factor
		factor /= 100
	}
	frequencies := frequencies(hand)
	sort.Sort(ByFrequencyAndCardValue(frequencies))
	if frequencies[0].count == 5 {
		//fmt.Println("five of a kind")
		strength += 70_000_000_000
	} else if frequencies[0].count == 4 {
		//fmt.Println("four of a kind")
		strength += 60_000_000_000
	} else if frequencies[0].count == 3 && frequencies[1].count == 2 {
		//fmt.Println("full house")
		strength += 50_000_000_000
	} else if frequencies[0].count == 3 {
		//fmt.Println("three of a kind")
		strength += 40_000_000_000
	} else if frequencies[0].count == 2 && frequencies[1].count == 2 {
		//fmt.Println("two pair")
		strength += 30_000_000_000
	} else if frequencies[0].count == 2 {
		//fmt.Println("one pair")
		strength += 20_000_000_000
	} else {
		strength += 10_000_000_000
	}
	//fmt.Println("Strength is: ", strength)
	return strength
}

func calculateStrengthPart2(hand string) int {
	//fmt.Println("Calulating Hand:", hand, "frequencies are:", frequencies)
	strength := 0
	factor := 100_000_000
	handRunes := []rune(hand)
	for idx := 0; idx < len(hand); idx++ {
		cardValue := cardValuePart1(handRunes[idx])
		strength += cardValue * factor
		factor /= 100
	}
	frequencies := frequencies(hand)
	sort.Sort(ByFrequencyAndCardValue(frequencies))
	if frequencies[0].count == 5 {
		//fmt.Println("five of a kind")
		strength += 70_000_000_000
	} else if frequencies[0].count == 4 {
		//fmt.Println("four of a kind")
		strength += 60_000_000_000
	} else if frequencies[0].count == 3 && frequencies[1].count == 2 {
		//fmt.Println("full house")
		strength += 50_000_000_000
	} else if frequencies[0].count == 3 {
		//fmt.Println("three of a kind")
		strength += 40_000_000_000
	} else if frequencies[0].count == 2 && frequencies[1].count == 2 {
		//fmt.Println("two pair")
		strength += 30_000_000_000
	} else if frequencies[0].count == 2 {
		//fmt.Println("one pair")
		strength += 20_000_000_000
	} else {
		strength += 10_000_000_000
	}
	//fmt.Println("Strength is: ", strength)
	return strength
}

func frequencies(hand string) []Frequency {
	freqMap := make(map[rune]int)
	for _, char := range hand {
		freqMap[char] = freqMap[char] + 1
	}
	freq := []Frequency{}
	for char, count := range freqMap {
		f := Frequency{
			count: count,
			char:  char,
		}
		freq = append(freq, f)
	}
	return freq
}

type Frequency = struct {
	count int
	char  rune
}

type ByFrequencyAndCardValue []Frequency

func (a ByFrequencyAndCardValue) Len() int {
	return len(a)
}

func (a ByFrequencyAndCardValue) Less(i, j int) bool {
	if a[i].count == a[j].count {
		return cardValuePart1(a[i].char) < cardValuePart1(a[j].char)
	}
	return a[i].count > a[j].count
}

func (a ByFrequencyAndCardValue) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
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
