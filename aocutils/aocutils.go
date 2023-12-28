package aocutils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func ReadFile(path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(content)
}

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

func StringsToNumbers(numStrings string) []int {
	strFields := strings.Fields(numStrings)
	numbers := []int{}
	for _, numStr := range strFields {
		numbers = append(numbers, StrToInt(numStr))
	}
	return numbers
}

func StrToUint(s string) uint {
	nr, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return uint(nr)
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

func SplitAndTrim(s string, delim string) []string {
	parts := strings.Split(s, delim)
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}

func PrintMatrix(matrix [][]rune) {
	for _, row := range matrix {
		fmt.Println(string(row))
	}
	fmt.Println("")
}

// Deprecated: remove variables if they're really not needed.
// Call Unused with variables that are temporarily not needed, to avoid the compiler from complaining that they're unused.
func Unused(unused ...interface{}) {}
