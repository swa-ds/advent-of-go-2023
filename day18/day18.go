package day18

import (
	"fmt"
	"regexp"
	"swads/aoc2023/aocutils"
)

type Position struct {
	x int
	y int
}

type Instruction struct {
	dir     string
	dist    int
	hexcode string
}

func Solve() {
	input := aocutils.ReadLines("day18/input.txt")

	part1 := SolvePart1(input)
	fmt.Println("Day 18 - Part 1:", part1)
}

// taken over solution from https://github.com/bsadia/aoc_goLang/blob/53ed198e644324d559a366a17c712e1b8c6bb4fe/day18/main.go
func SolvePart1(input []string) int {
	re := regexp.MustCompile(`([UDLR]) (\d+) \((#[a-z0-9]{6})\)`)
	instructions := []Instruction{}

	for _, line := range input {
		matches := re.FindStringSubmatch(line)
		in := Instruction{
			matches[1],
			aocutils.StrToInt(matches[2]),
			matches[3],
		}
		instructions = append(instructions, in)
		//fmt.Println(in)
	}

	pos := Position{0, 0}
	area := 0

	for _, inst := range instructions {
		x := pos.x
		y := pos.y
		if inst.dir == "U" {
			y -= inst.dist
		} else if inst.dir == "D" {
			y += inst.dist
		} else if inst.dir == "L" {
			x -= inst.dist
		} else if inst.dir == "R" {
			x += inst.dist
		}
		next := Position{x, y}

		area += (pos.x*next.y - pos.y*next.x) + inst.dist // shoelace
		pos = next
	}

	return area/2 + 1
}
