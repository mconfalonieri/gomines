package minefield

import (
	"fmt"
	"math/rand/v2"
	"slices"
)

// MineField represents the mine field.
type MineField struct {
	width  int
	height int
	mines  int
	Field  [][]bool
}

func createField(w, h, m int) [][]bool {
	// Create an empty field
	f := make([][]bool, w)
	for i := range f {
		f[i] = make([]bool, h)
	}

	// Create an array of indices
	n := w * h
	indexes := make([]int, m)
	for i := range indexes {
		indexes[i] = -1
	}
	mines := 0
	for mines < m {
		sel := rand.IntN(n)
		if slices.Contains(indexes, sel) {
			continue
		}
		indexes[mines] = sel
		x := sel / h
		y := sel % h

		// Position the mine
		f[x][y] = true
		mines++
	}
	return f
}

// New creates a new mine field.
func New(w, h, m int) (*MineField, error) {
	n := w * h
	if m < 1 || m >= n {
		max := n - 1
		return nil, fmt.Errorf("the number of mines must be between 1 and %d", max)
	}
	f := createField(w, h, m)
	return &MineField{
		width:  w,
		height: h,
		mines:  m,
		Field:  f,
	}, nil
}

func (m MineField) Size() (int, int) {
	return m.width, m.height
}

func (m MineField) Mines() int {
	return m.mines
}
