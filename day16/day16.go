package day16

import (
	"fmt"
	util "swads/aoc2023/aocutils"

	mapset "github.com/deckarep/golang-set/v2"
)

type Position struct {
	x int
	y int
}

type Beam struct {
	direction int
	pos       Position
}

func NewBeam(direction int, position Position) Beam {
	return Beam{direction, position}
}

const (
	Up = iota
	Down
	Left
	Right
)

// save seen beams to prevent infinite loops
var seen = mapset.NewSet[Beam]()

func Solve() {
	input := util.ReadLines("day16/input.txt")

	part1 := SolvePart1(input)
	fmt.Println("Day 16 - Part 1:", part1)
}

func SolvePart1(input []string) int {
	grid := util.LinesToMatrix(input)

	beam := NewBeam(Right, Position{0, 0})

	energized := mapset.NewSet[Position]()

	run(beam, grid, energized)

	//printEnergized(grid, energized)

	return energized.Cardinality()
}

func SolvePart2(input []string) int {
	return 0
}

func run(beam Beam, grid [][]rune, energized mapset.Set[Position]) {
	seen.Add(beam)
	//fmt.Println("starting new run", beam)
	//defer fmt.Println("ending run", beam)

	for i := 0; true; i++ {
		if outsideBounds(beam, grid) {
			return
		}
		energized.Add(beam.pos)
		//fmt.Println("beam:", beam, "on", string(grid[beam.pos.y][beam.pos.x]))
		if beam.direction == Right {
			if grid[beam.pos.y][beam.pos.x] == '/' {
				beam.direction = Up
			} else if grid[beam.pos.y][beam.pos.x] == '\\' {
				beam.direction = Down
			} else if grid[beam.pos.y][beam.pos.x] == '|' {
				beam.direction = Down
				newBeam := NewBeam(Up, Position{beam.pos.x, beam.pos.y - 1})
				run(newBeam, grid, energized)
			}
		} else if beam.direction == Left {
			if grid[beam.pos.y][beam.pos.x] == '/' {
				beam.direction = Down
			} else if grid[beam.pos.y][beam.pos.x] == '\\' {
				beam.direction = Up
			} else if grid[beam.pos.y][beam.pos.x] == '|' {
				beam.direction = Down
				newBeam := NewBeam(Up, Position{beam.pos.x, beam.pos.y - 1})
				run(newBeam, grid, energized)
			}
		} else if beam.direction == Up {
			if grid[beam.pos.y][beam.pos.x] == '/' {
				beam.direction = Right
			} else if grid[beam.pos.y][beam.pos.x] == '\\' {
				beam.direction = Left
			} else if grid[beam.pos.y][beam.pos.x] == '-' {
				beam.direction = Left
				newBeam := NewBeam(Right, Position{beam.pos.x + 1, beam.pos.y})
				run(newBeam, grid, energized)
			}
		} else if beam.direction == Down {
			if grid[beam.pos.y][beam.pos.x] == '/' {
				beam.direction = Left
			} else if grid[beam.pos.y][beam.pos.x] == '\\' {
				beam.direction = Right
			} else if grid[beam.pos.y][beam.pos.x] == '-' {
				beam.direction = Left
				newBeam := NewBeam(Right, Position{beam.pos.x + 1, beam.pos.y})
				run(newBeam, grid, energized)
			}
		}
		if beam.direction == Up {
			moveUp(&beam)
		} else if beam.direction == Down {
			moveDown(&beam)
		} else if beam.direction == Left {
			moveLeft(&beam)
		} else if beam.direction == Right {
			moveRight(&beam)
		}
		if seen.Contains(beam) {
			break
		}
		seen.Add(beam)
	}
}

func outsideBounds(beam Beam, grid [][]rune) bool {
	return beam.pos.y >= len(grid) || beam.pos.y < 0 || beam.pos.x < 0 || beam.pos.x >= len(grid[0])
}

func moveUp(beam *Beam) {
	beam.pos = Position{beam.pos.x, beam.pos.y - 1}
}

func moveDown(beam *Beam) {
	beam.pos = Position{beam.pos.x, beam.pos.y + 1}
}

func moveRight(beam *Beam) {
	beam.pos = Position{beam.pos.x + 1, beam.pos.y}
}

func moveLeft(beam *Beam) {
	beam.pos = Position{beam.pos.x - 1, beam.pos.y}
}

func printEnergized(grid [][]rune, energized mapset.Set[Position]) {
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if energized.Contains(Position{col, row}) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}