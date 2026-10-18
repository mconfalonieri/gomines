package setup

import (
	"errors"
	"fmt"
	"strconv"

	"charm.land/bubbles/v2/textinput"
	"github.com/mconfalonieri/gomines/internal/pkg/styles"
)

const (
	prompt = "> " // textinput prompt

	// TextInput field limits
	nameFieldLength  = 100
	nameFieldSize    = 25
	sizeFieldLength  = 3
	sizeFieldSize    = 3
	sizeMinValue     = 5
	sizeMaxValue     = 100
	minesFieldLength = 4
	minesFieldSize   = 4
	minesMinValue    = 1

	// Default values
	defaultName  = "Player One"
	defaultSize  = "10"
	defaultMines = "10"
)

// createPlayerNameTI creates a player name textinput.
func createPlayerNameTI() textinput.Model {
	n := textinput.New()
	n.Placeholder = defaultName
	n.SetVirtualCursor(false)
	n.Focus()
	n.CharLimit = nameFieldLength
	n.SetWidth(nameFieldSize)
	n.Prompt = styles.Prompt.Render(prompt)
	return n
}

// createSizeTI creates a size textinput.
func createSizeTI() textinput.Model {
	n := textinput.New()
	n.Placeholder = defaultSize
	n.SetVirtualCursor(false)
	n.Focus()
	n.CharLimit = sizeFieldLength
	n.SetWidth(sizeFieldSize)
	n.Prompt = styles.Prompt.Render(prompt)
	return n
}

// createMineNumberTI creates a mine number textinput.
func createMineNumberTI() textinput.Model {
	n := textinput.New()
	n.Placeholder = defaultMines
	n.SetVirtualCursor(false)
	n.Focus()
	n.CharLimit = minesFieldLength
	n.SetWidth(minesFieldSize)
	n.Prompt = styles.Prompt.Render(prompt)
	return n
}

func checkSize(_ Options, v string) error {
	n, err := strconv.Atoi(v)
	if err != nil {
		return errors.New("this value is not an integer number.")
	} else if n < sizeMinValue || sizeMaxValue > 100 {
		return fmt.Errorf("this value must be beetween %d and %d.",
			sizeMinValue, sizeMaxValue)
	}
	return nil
}

func checkMines(o Options, v string) error {
	n, err := strconv.Atoi(v)
	sw := o.getValue(iWidth)
	sh := o.getValue(iHeight)
	w, _ := strconv.Atoi(sw)
	h, _ := strconv.Atoi(sh)
	max := w*h - 1
	if err != nil {
		return errors.New("this value is not an integer number.")
	} else if n < minesMinValue || n > max {
		return fmt.Errorf("this value must be beetween %d and %d", minesMinValue, max)
	}
	return nil
}
