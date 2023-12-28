package day16

import (
	"fmt"
	"slices"
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

func (self *Beam) move() {
	if self.direction == Up {
		self.pos.y--
	} else if self.direction == Down {
		self.pos.y++
	} else if self.direction == Left {
		self.pos.x--
	} else if self.direction == Right {
		self.pos.x++
	}
}

const (
	Up = iota
	Down
	Left
	Right
)

func Solve() {
	input := util.ReadLines("day16/input.txt")

	part1 := SolvePart1(input)
	fmt.Println("Day 16 - Part 1:", part1)

	part2 := SolvePart2(input)
	fmt.Println("Day 16 - Part 2:", part2)
}

func SolvePart1(input []string) int {
	grid := util.LinesToMatrix(input)

	beam := NewBeam(Right, Position{0, 0})

	return energize(beam, grid)
}

func SolvePart2(input []string) int {
	grid := util.LinesToMatrix(input)
	beams := []Beam{}

	rows := len(grid)
	cols := len(grid[0])

	for row := 0; row < rows; row++ {
		beams = append(beams, Beam{Right, Position{0, row}})
		beams = append(beams, Beam{Left, Position{cols - 1, row}})
	}
	for col := 0; col < cols; col++ {
		beams = append(beams, Beam{Down, Position{col, 0}})
		beams = append(beams, Beam{Up, Position{col, rows - 1}})
	}
	//fmt.Println(beams)

	energized := []int{}

	for _, beam := range beams {
		energized = append(energized, energize(beam, grid))
	}

	return slices.Max(energized)
}

func energize(beam Beam, grid [][]rune) int {
	energized := mapset.NewSet[Position]()
	// save seen beams to prevent endless loops
	seen := mapset.NewSet[Beam]()

	run(beam, grid, energized, seen)

	//printEnergized(grid, energized)

	return energized.Cardinality()
}

func run(beam Beam, grid [][]rune, energized mapset.Set[Position], seen mapset.Set[Beam]) {
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
				run(newBeam, grid, energized, seen)
			}
		} else if beam.direction == Left {
			if grid[beam.pos.y][beam.pos.x] == '/' {
				beam.direction = Down
			} else if grid[beam.pos.y][beam.pos.x] == '\\' {
				beam.direction = Up
			} else if grid[beam.pos.y][beam.pos.x] == '|' {
				beam.direction = Down
				newBeam := NewBeam(Up, Position{beam.pos.x, beam.pos.y - 1})
				run(newBeam, grid, energized, seen)
			}
		} else if beam.direction == Up {
			if grid[beam.pos.y][beam.pos.x] == '/' {
				beam.direction = Right
			} else if grid[beam.pos.y][beam.pos.x] == '\\' {
				beam.direction = Left
			} else if grid[beam.pos.y][beam.pos.x] == '-' {
				beam.direction = Left
				newBeam := NewBeam(Right, Position{beam.pos.x + 1, beam.pos.y})
				run(newBeam, grid, energized, seen)
			}
		} else if beam.direction == Down {
			if grid[beam.pos.y][beam.pos.x] == '/' {
				beam.direction = Left
			} else if grid[beam.pos.y][beam.pos.x] == '\\' {
				beam.direction = Right
			} else if grid[beam.pos.y][beam.pos.x] == '-' {
				beam.direction = Left
				newBeam := NewBeam(Right, Position{beam.pos.x + 1, beam.pos.y})
				run(newBeam, grid, energized, seen)
			}
		}
		beam.move()
		if seen.Contains(beam) {
			break
		}
		seen.Add(beam)
	}
}

func outsideBounds(beam Beam, grid [][]rune) bool {
	return beam.pos.y >= len(grid) || beam.pos.y < 0 || beam.pos.x < 0 || beam.pos.x >= len(grid[0])
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
