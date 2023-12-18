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
var cardOrder = ""

func cardValue(r rune) int {
	//fmt.Println("card order", cardOrder)
	return strings.Index(cardOrder, string(r)) + 1
}

func Solve() {
	input := aocutils.ReadLines("day07/input.txt")
	resultPart1 := SolvePart(input, part1)
	fmt.Println("Day 07 - Part 1:", resultPart1)
	resultPart2 := SolvePart(input, part2)
	fmt.Println("Day 07 - Part 2:", resultPart2)
}

func SolvePart(input []string, part int) int {
	hands := []Hand{}
	for _, hand := range input {
		hands = append(hands, NewHand(hand, part))
	}
	result := calculateResult(hands)
	return result
}

func calculateResult(hands []Hand) int {
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
	cardOrder = cardOrderPart1
	//fmt.Println("card order:", cardOrder)
	//fmt.Println("Calulating Hand:", hand, "frequencies are:", frequencies)
	strength := 0
	factor := 100_000_000
	handRunes := []rune(hand)
	for idx := 0; idx < len(hand); idx++ {
		cardValue := cardValue(handRunes[idx])
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
	cardOrder = cardOrderPart2
	//fmt.Println("card order:", cardOrder)
	//fmt.Println("Calulating Hand:", hand, "frequencies are:", frequencies)
	strength := 0
	factor := 100_000_000
	handRunes := []rune(hand)
	for idx := 0; idx < len(hand); idx++ {
		cardValue := cardValue(handRunes[idx])
		strength += cardValue * factor
		factor /= 100
	}
	frequencies := frequencies(hand)

	jokerIdx := -1
	for i, freq := range frequencies {
		if freq.char == 'J' {
			jokerIdx = i
			break
		}
	}
	jokerCount := 0
	if jokerIdx >= 0 {
		jokerCount = frequencies[jokerIdx].count
		// remove joker from frequencies
		frequencies = append(frequencies[:jokerIdx], frequencies[jokerIdx+1:]...)
	}
	sort.Sort(ByFrequencyAndCardValue(frequencies))
	// for special case 'JJJJJ' we get an empty slice because frequency for 'J' has been filtered out before
	if len(frequencies) == 0 {
		frequencies = make([]Frequency, 1)
	}
	frequencies[0].count += jokerCount
	//fmt.Println(frequencies)
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
		return cardValue(a[i].char) < cardValue(a[j].char)
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
