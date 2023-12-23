package day11

import (
	"fmt"
	"math"
	"strings"
	util "swads/aoc2023/aocutils"
)

func Solve() {
	input := util.ReadLines("day11/input.txt")

	part1 := SolvePart1(input)
	fmt.Println("Day 11 - Part 1:", part1)
}

func SolvePart1(input []string) int {
	space := expandSpace(input)
	galaxies := []Galaxy{}
	for row := 0; row < len(space); row++ {
		for col := 0; col < len(space[row]); col++ {
			if space[row][col] == '#' {
				galaxies = append(galaxies, Galaxy{row, col})
			}
		}
	}
	pairs := findAllPairs(galaxies)
	//fmt.Println("Galaxies:", galaxies)
	//fmt.Println("Pairs:", pairs)

	result := 0
	for _, pair := range pairs {
		result += shortestDistance(pair[0], pair[1])
	}
	return result
}

func expandSpace(input []string) [][]rune {
	rowsExpanded := expandRows(input)
	space := expandCols(rowsExpanded)
	return space
}

func expandCols(rows []string) (space [][]rune) {
	for _, row := range rows {
		space = append(space, []rune(row))
	}
	rowLength := len(space[0])
	rowCount := len(space)
	for col := 0; col < rowLength; col++ {
		hasNoGalaxy := true
		for row := 0; row < rowCount; row++ {
			if space[row][col] == '#' {
				hasNoGalaxy = false
			}
		}
		if hasNoGalaxy {
			//fmt.Println("extending column", col)
			for row := 0; row < rowCount; row++ {
				r := space[row]
				r = append(r[:col], append([]rune{'.'}, r[col:]...)...)
				space[row] = r
			}
			//skip inserted column
			col++
			// length of row has been increased by the inserted column
			rowLength++
		}
	}
	return
}

func expandRows(rows []string) []string {
	for i := 0; i < len(rows); i++ {
		if !strings.Contains(rows[i], "#") {
			// insert additional row without galaxy at current row index
			rows = append(rows[:i], append([]string{rows[i]}, rows[i:]...)...)
			// skip inserted row
			i++
		}
	}
	return rows
}

func findAllPairs(galaxies []Galaxy) [][]Galaxy {
	var pairs [][]Galaxy

	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			pairs = append(pairs, []Galaxy{galaxies[i], galaxies[j]})
		}
	}
	return pairs
}

func shortestDistance(a, b Galaxy) int {
	return int(math.Abs(float64(a.col-b.col)) + math.Abs(float64(a.row-b.row)))
}

type Galaxy struct {
	row int
	col int
}
