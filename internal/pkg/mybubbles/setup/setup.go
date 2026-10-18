package setup

import (
	tea "charm.land/bubbletea/v2"
)

const (
	// Prompts
	promptPlayerName = "What is your glamourous name?"
	promptWidth      = "What is the width of your minefield?"
	promptHeight     = "What is the height of your minefield?"
	promptMineNumber = "How many mines do you want to sweep?"
	promptGlyphSet   = "Which glyph set looks better on your terminal?"
)

// initialModel prepares all the fields for the model.
func initialModel(o *Options) model {

	return model{
		opt: o,
		fields: []field{
			{
				Header:         promptPlayerName,
				TextInputModel: createPlayerNameTI(),
				Type:           fieldTextInput,
			},
			{
				Header:         promptWidth,
				TextInputModel: createSizeTI(),
				Check:          checkSize,
				Type:           fieldTextInput,
			},
			{
				Header:         promptHeight,
				TextInputModel: createSizeTI(),
				Check:          checkSize,
				Type:           fieldTextInput,
			},
			{
				Header:         promptMineNumber,
				TextInputModel: createMineNumberTI(),
				Check:          checkMines,
				Type:           fieldTextInput,
			},
			{
				Header:    promptGlyphSet,
				ListModel: createGlyphSetList(),
				Type:      fieldList,
			},
		},
	}
}

// Run runs the setup process.
func Run() (Options, error) {
	o := newOptions()
	m := initialModel(o)
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		return Options{}, err
	}
	return *o, nil
}
