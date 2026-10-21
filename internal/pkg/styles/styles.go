package styles

import (
	"os"

	"charm.land/bubbles/v2/textinput"
	"charm.land/lipgloss/v2"
	"charm.land/tone"
)

var (
	TextInputStyles = textinput.Styles{}

	Title       = lipgloss.Style{}
	KeyChar     = lipgloss.Style{}
	Prompt      = lipgloss.Style{}
	Error       = lipgloss.Style{}
	GridNormal  = lipgloss.Style{}
	GridFocused = lipgloss.Style{}
	Mine        = lipgloss.Style{}
	MineOK      = lipgloss.Style{}
	OneMine     = lipgloss.Style{}
	TwoMines    = lipgloss.Style{}
	ThreeMines  = lipgloss.Style{}
	FourMines   = lipgloss.Style{}
	FiveMines   = lipgloss.Style{}
	SixMines    = lipgloss.Style{}
	SevenMines  = lipgloss.Style{}
	EightMines  = lipgloss.Style{}
	CheckMark   = lipgloss.Style{}
	XMark       = lipgloss.Style{}
	QMark       = lipgloss.Style{}
	StatusBar   = lipgloss.Style{}
	MessageWon  = lipgloss.Style{}
	MessageLost = lipgloss.Style{}
)

func init() {
	if lipgloss.HasDarkBackground(os.Stdin, os.Stdout) {
		TextInputStyles = textinput.DefaultDarkStyles()
		Title = lipgloss.NewStyle().
			Foreground(tone.Citron).
			Background(tone.Charple).
			Bold(true).AlignHorizontal(lipgloss.Center)
		KeyChar = lipgloss.NewStyle().
			Foreground(tone.Julep).
			Background(tone.BBQ).
			Bold(true)
		Prompt = lipgloss.NewStyle().
			Foreground(tone.Orchid).
			Bold(true)
		Error = lipgloss.NewStyle().
			Foreground(tone.Mustard).
			Background(lipgloss.Red)
		GridNormal = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(tone.Salt)
		GridFocused = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(tone.Malibu)
		Mine = lipgloss.NewStyle().
			Foreground(tone.Citron)
		MineOK = lipgloss.NewStyle().
			Foreground(tone.Julep)
		OneMine = lipgloss.NewStyle().
			Foreground(tone.Malibu)
		TwoMines = lipgloss.NewStyle().
			Foreground(tone.Julep)
		ThreeMines = lipgloss.NewStyle().
			Foreground(tone.Cheeky)
		FourMines = lipgloss.NewStyle().
			Foreground(tone.Guppy)
		FiveMines = lipgloss.NewStyle().
			Foreground(tone.Tuna)
		SixMines = lipgloss.NewStyle().
			Foreground(tone.Sardine)
		SevenMines = lipgloss.NewStyle().
			Foreground(tone.Butter).Bold(true)
		EightMines = lipgloss.NewStyle().
			Foreground(tone.Smoke).Bold(true)
		CheckMark = lipgloss.NewStyle().
			Foreground(tone.Blush)
		XMark = lipgloss.NewStyle().
			Foreground(tone.Coral)
		QMark = lipgloss.NewStyle().
			Foreground(tone.Citron)
		StatusBar = lipgloss.NewStyle().
			Background(tone.Ox).
			Foreground(tone.Bok)
		MessageWon = lipgloss.NewStyle().
			Foreground(tone.Lichen)
		MessageLost = lipgloss.NewStyle().
			Foreground(tone.Salmon)
	} else {
		TextInputStyles = textinput.DefaultLightStyles()
		Title = lipgloss.NewStyle().
			Foreground(tone.Cumin).
			Background(tone.Hazy).
			Bold(true)
		KeyChar = lipgloss.NewStyle().
			Foreground(tone.Guac).
			Background(tone.Steep).
			Bold(true)
		Prompt = lipgloss.NewStyle().
			Foreground(tone.Grape).
			Bold(true)
		GridNormal = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(tone.Iron)
		GridFocused = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(tone.Malibu)
		Mine = lipgloss.NewStyle().
			Foreground(tone.Cumin)
		MineOK = lipgloss.NewStyle().
			Foreground(tone.Guac)
		OneMine = lipgloss.NewStyle().
			Foreground(tone.Damson)
		TwoMines = lipgloss.NewStyle().
			Foreground(tone.Guac)
		ThreeMines = lipgloss.NewStyle().
			Foreground(tone.Macaron)
		FourMines = lipgloss.NewStyle().
			Foreground(tone.Sapphire)
		FiveMines = lipgloss.NewStyle().
			Foreground(tone.Chili)
		SixMines = lipgloss.NewStyle().
			Foreground(tone.Zinc)
		SevenMines = lipgloss.NewStyle().
			Foreground(tone.BBQ).Bold(true)
		EightMines = lipgloss.NewStyle().
			Foreground(tone.Oyster).Bold(true)
		CheckMark = lipgloss.NewStyle().
			Foreground(tone.Flamingo)
		XMark = lipgloss.NewStyle().
			Foreground(tone.Coral)
		QMark = lipgloss.NewStyle().
			Foreground(tone.Cumin)
		MessageWon = lipgloss.NewStyle().
			Foreground(tone.Zinc)
		MessageLost = lipgloss.NewStyle().
			Foreground(tone.Sriracha)
	}
}
