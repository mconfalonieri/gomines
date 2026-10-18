package setup

import (
	"fmt"
	"strconv"
)

const (
	// Value indexes
	iName     = 0 // name value index
	iWidth    = 1 // field width value index
	iHeight   = 2 // field height value index
	iMines    = 3 // number of mines index
	iGlyphSet = 4 // chosen glyph set

	// Number of values
	totValues = 5
)

// Options resulting from the setup process.
type Options struct {
	quit   bool     // true if quit was called during the setup process
	values []string // all the values stored as strings
}

// newOptions creates a new Options object.
func newOptions() *Options {
	return &Options{
		values: make([]string, totValues),
	}
}

// getValue returns a specific value.
func (o Options) getValue(i int) string {
	return o.values[i]
}

// getIntValue returns a value as an integer. It panics if it cannot be
// converted.
func (o Options) getIntValue(i int) int {
	v := o.values[i]
	n, err := strconv.Atoi(v)
	if err != nil {
		panic(fmt.Sprintf("Unexpected error while converting \"%s\" to int: %s", v, err.Error()))
	}
	return n
}

// setValue sets a value at the specific index.
func (o *Options) setValue(i int, v string) {
	o.values[i] = v
}

// GetPlayerName returns the player name.
func (o Options) GetPlayerName() string {
	return o.values[iName]
}

// GetWidth returns the mine field width.
func (o Options) GetWidth() int {
	return o.getIntValue(iWidth)
}

// GetHeight returns the mine field height.
func (o Options) GetHeight() int {
	return o.getIntValue(iHeight)
}

// GetMines returns the mine number.
func (o Options) GetMines() int {
	return o.getIntValue(iMines)
}

// GetMines returns the mine number.
func (o Options) GetGlyphSet() int {
	return o.getIntValue(iGlyphSet)
}

// HasQuit returns the value of the quit flag.
func (o Options) HasQuit() bool {
	return o.quit
}
