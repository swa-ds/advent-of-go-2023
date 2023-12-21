package day10

import (
	"fmt"
	"swads/aoc2023/aocutils"
)

const (
	North = iota
	East
	South
	West
)

func Solve() {
	input := aocutils.ReadLines("day10/input.txt")

	part1 := SolvePart1(input)
	fmt.Println("Day 10 - Part 1:", part1)
}

func SolvePart1(input []string) int {
	field := parseField(input)
	field.move()
	//fmt.Println(string(field.symbolOnPos()))
	var steps int
	for steps = 1; !field.onStart() || steps > 1_000_000; steps++ {
		field.move()
		//fmt.Println(string(field.symbolOnPos()))
	}
	//fmt.Println("steps:", steps)
	return steps / 2
}

func parseField(input []string) Field {
	field := [][]rune{}
	start := Position{0, 0}
	for y, line := range input {
		rline := []rune(line)
		field = append(field, rline)
		for x, r := range rline {
			if r == 'S' {
				start = Position{x, y}
			}
		}
	}
	return Field{field, start, start, South}
}

type Field struct {
	field     [][]rune
	start     Position
	pos       Position
	direction int
}

type Position struct {
	x int
	y int
}

func (f Field) onStart() bool {
	return f.start == f.pos
}

func (f Field) symbolOnPos() rune {
	return f.field[f.pos.y][f.pos.x]
}

func (f *Field) move() {
	if f.symbolOnPos() == 'S' {
		f.goSouth()
	} else if f.symbolOnPos() == 'F' {
		f.moveF()
	} else if f.symbolOnPos() == '7' {
		f.move7()
	} else if f.symbolOnPos() == 'L' {
		f.moveL()
	} else if f.symbolOnPos() == 'J' {
		f.moveJ()
	} else if f.symbolOnPos() == '|' {
		f.movePipe()
	} else if f.symbolOnPos() == '-' {
		f.moveMinus()
	}
}

func (f *Field) moveF() {
	if f.direction == North {
		f.goEast()
	} else if f.direction == West {
		f.goSouth()
	}
}

func (f *Field) move7() {
	if f.direction == East {
		f.goSouth()
	} else if f.direction == North {
		f.goWest()
	}
}

func (f *Field) moveL() {
	if f.direction == South {
		f.goEast()
	} else if f.direction == West {
		f.goNorth()
	}
}

func (f *Field) moveJ() {
	if f.direction == South {
		f.goWest()
	} else if f.direction == East {
		f.goNorth()
	}
}

func (f *Field) movePipe() {
	if f.direction == South {
		f.goSouth()
	} else if f.direction == North {
		f.goNorth()
	}
}

func (f *Field) moveMinus() {
	if f.direction == East {
		f.goEast()
	} else if f.direction == West {
		f.goWest()
	}
}

func (f *Field) goNorth() {
	f.direction = North
	f.pos = Position{f.pos.x, f.pos.y - 1}
}

func (f *Field) goEast() {
	f.direction = East
	f.pos = Position{f.pos.x + 1, f.pos.y}
}

func (f *Field) goSouth() {
	f.direction = South
	f.pos = Position{f.pos.x, f.pos.y + 1}
}

func (f *Field) goWest() {
	f.direction = West
	f.pos = Position{f.pos.x - 1, f.pos.y}
}
