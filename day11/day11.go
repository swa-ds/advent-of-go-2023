package day11

import (
	"fmt"
	"math"
	"strings"
	util "swads/aoc2023/aocutils"
)

func Solve() {
	input := util.ReadLines("day11/input.txt")

	part1 := SolvePart2(input, 2)
	fmt.Println("Day 11 - Part 1:", part1)

	part2 := SolvePart2(input, 1_000_000)
	fmt.Println("Day 11 - Part 2:", part2)

}

func SolvePart1(input []string) int {
	// Part 1: actually expand the space grid
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

func SolvePart2(input []string, expansionRate int) int {
	space := [][]rune{}
	expandedRows := []int{}
	expandedCols := []int{}
	for rowNum, row := range input {
		space = append(space, []rune(row))
		if !strings.Contains(row, "#") {
			expandedRows = append(expandedRows, rowNum)
		}
	}
	for col := 0; col < len(space[0]); col++ {
		containsNoGalaxy := true
		for row := 0; row < len(space); row++ {
			if space[row][col] == '#' {
				containsNoGalaxy = false
			}
		}
		if containsNoGalaxy {
			expandedCols = append(expandedCols, col)
		}
	}
	// due to the large expansion rate for part 2 (1 mio)
	// an actual expansion as done in part1 would not be feasible.
	// Therefore instead, we just calculate the position of each galaxy
	// after the expansion, according to the expanded rows & columns
	galaxies := []Galaxy{}
	for row := 0; row < len(space); row++ {
		for col := 0; col < len(space[row]); col++ {
			if space[row][col] == '#' {
				galaxyRow := expandCoord(row, expansionRate, expandedRows)
				galaxyCol := expandCoord(col, expansionRate, expandedCols)
				galaxies = append(galaxies, Galaxy{galaxyRow, galaxyCol})
			}
		}
	}
	//fmt.Println("expanded rows:", expandedRows)
	//fmt.Println("expanded cols:", expandedCols)

	pairs := findAllPairs(galaxies)
	//fmt.Println("Galaxies:", galaxies)
	//fmt.Println("Pairs:", pairs)
	result := 0
	for _, pair := range pairs {
		result += shortestDistance(pair[0], pair[1])
	}
	return result
}

// coord can be row or column
func expandCoord(coord int, expansionRate int, expandedCoords []int) int {
	expands := 0
	for _, exp := range expandedCoords {
		if coord > exp {
			expands++
		}
	}
	return coord + (expansionRate-1)*expands
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
