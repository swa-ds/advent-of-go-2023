package day15

import (
	"fmt"
	"strings"
	util "swads/aoc2023/aocutils"
)

func Solve() {
	input := util.ReadFile("day15/input.txt")

	part1 := SolvePart1(input)
	fmt.Println("Day 15 - Part 1:", part1)

	part2 := SolvePart2(input)
	fmt.Println("Day 15 - Part 2:", part2)

}

func SolvePart1(input string) int {
	steps := strings.Split(input, ",")
	result := 0
	for _, step := range steps {
		result += hash(step)
	}
	return result
}

func SolvePart2(input string) int {
	result := 0
	steps := strings.Split(input, ",")
	boxes := make([][]Box, 256)
	for _, step := range steps {
		var label string
		if strings.Contains(step, "-") {
			label = step[:len(step)-1]
			remove(label, boxes)
		} else {
			split := strings.Split(step, "=")
			label = split[0]
			focLen := util.StrToInt(split[1])
			put(Box{label, focLen}, boxes)
		}
	}

	for i, lenses := range boxes {
		if len(lenses) > 0 {
			//fmt.Println("Box", i, lenses)
			for j, lens := range lenses {
				result += (i + 1) * (j + 1) * lens.focalLen
			}
		}
	}

	return result
}

func put(box Box, boxes [][]Box) {
	hash := hash(box.label)
	lenses := boxes[hash]
	for i, lens := range lenses {
		if lens.label == box.label {
			lenses[i] = box
			boxes[hash] = lenses
			return
		}
	}
	boxes[hash] = append(boxes[hash], box)
}

func remove(label string, boxes [][]Box) {
	hash := hash(label)
	lenses := boxes[hash]
	for i, box := range lenses {
		if box.label == label {
			lenses = append(lenses[:i], lenses[i+1:]...)
			boxes[hash] = lenses
			return
		}
	}
}

func hash(s string) int {
	hash := 0
	for _, r := range []rune(s) {
		hash += int(r)
		hash *= 17
		hash = hash % 256
	}
	return hash
}

type Box struct {
	label    string
	focalLen int
}
