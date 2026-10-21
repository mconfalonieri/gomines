package setup

import (
	"charm.land/bubbles/v2/list"
	"charm.land/bubbles/v2/textinput"
)

const (
	// Field types
	fieldTextInput = 0 // textinput
	fieldList      = 1 // list
)

// field represents an input field.
type field struct {
	Type           int
	TextInputModel textinput.Model
	ListModel      list.Model
	Header         string
	Check          func(Options, string) error
}

func (f *field) Focus() {
	if f.Type == fieldTextInput {
		f.TextInputModel.Focus()
	}
}

func (f *field) Blur() {
	if f.Type == fieldTextInput {
		f.TextInputModel.Blur()
	}
}
