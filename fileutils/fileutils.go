package fileutils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func ReadMatrix(path string) [][]rune {
	lines := ReadLines(path)
	return LinesToMatrix(lines)
}

func StringToMatrix(input string) [][]rune {
	return LinesToMatrix(strings.Split(input, "\n"))
}

func LinesToMatrix(lines []string) [][]rune {
	matrix := make([][]rune, len(lines))
	for i, line := range lines {
		matrix[i] = []rune(line)
	}
	return matrix
}

func StrToInt(s string) int {
	nr, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return nr
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// Deprecated: remove variables if they're really not needed.
// Call Unused with variables that are temporarily not needed, to avoid the compiler from complaining that they're unused.
func Unused(unused ...interface{}) {}
