package fileutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadMatrix(t *testing.T) {
	expectedMatrix := [][]rune{
		{'a', 'b', 'c'},
		{'d', '3', 'f'},
		{'g', 'h', '1'},
	}

	matrix := ReadMatrix("../test/fileutil/matrix.txt")

	for y := 0; y < len(expectedMatrix); y++ {
		for x := 0; x < len(expectedMatrix[0]); x++ {
			if expectedMatrix[y][x] != matrix[y][x] {
				t.Errorf("Error at y=%d, x=%d: want %c, but got %c", y, x, expectedMatrix[y][x], matrix[y][x])
			}
		}
	}
}

func TestMin(t *testing.T) {
	assert.Equal(t, 3, Min(3, 8))
	assert.Equal(t, 3, Min(8, 3))
}

func TestMax(t *testing.T) {
	assert.Equal(t, 8, Max(3, 8))
	assert.Equal(t, 8, Max(8, 3))
}
