package main

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/mconfalonieri/gomines/internal/pkg/gamelogic"
	"github.com/mconfalonieri/gomines/internal/pkg/glyphs"
	"github.com/mconfalonieri/gomines/internal/pkg/mybubbles/grid"
	"github.com/mconfalonieri/gomines/internal/pkg/styles"
)

// render renders a string according to its style.
func render(s string) string {
	gmap := map[string]string{
		gamelogic.CellEmpty:   " ",
		gamelogic.CellMarked:  styles.CheckMark.Render(glyphs.Get(glyphs.GCheck)),
		gamelogic.CellUnknown: styles.QMark.Render(glyphs.Get(glyphs.GQMark)),
		gamelogic.CellXMarked: styles.XMark.Render(glyphs.Get(glyphs.GXMark)),
		gamelogic.CellMine:    styles.Mine.Render(glyphs.Get(glyphs.GMine)),
		gamelogic.CellMineOK:  styles.MineOK.Render(glyphs.Get(glyphs.GCheck)),
		"1":                   styles.OneMine.Render(glyphs.Get(glyphs.GOne)),
		"2":                   styles.TwoMines.Render(glyphs.Get(glyphs.GTwo)),
		"3":                   styles.ThreeMines.Render(glyphs.Get(glyphs.GThree)),
		"4":                   styles.FourMines.Render(glyphs.Get(glyphs.GFour)),
		"5":                   styles.FiveMines.Render(glyphs.Get(glyphs.GFive)),
		"6":                   styles.SixMines.Render(glyphs.Get(glyphs.GSix)),
		"7":                   styles.SevenMines.Render(glyphs.Get(glyphs.GSeven)),
		"8":                   styles.EightMines.Render(glyphs.Get(glyphs.GEight)),
	}
	g, ok := gmap[s]
	if !ok {
		return s
	}
	return g
}

// cell is used by the grid component.
type cell string

// Render renders a single cell.
func (c cell) Render(width, height int, selected bool) string {
	s := string(c)
	if selected {
		return styles.Prompt.Render(">") + render(s) + styles.Prompt.Render("<")
	} else {
		return " " + render(s) + " "
	}
}

// asGridItems converts a string matrix to a cell matrix, returned as a
// grid.GridItem matrix.
func asGridItems(w, h int, f [][]string) [][]grid.GridItem {
	g := make([][]grid.GridItem, h)
	for y := range g {
		g[y] = make([]grid.GridItem, w)
		for x := range g[y] {
			g[y][x] = cell(f[x][y])
		}
	}
	return g
}

// model represents the ELM model used by BubbleTea.
type model struct {
	opts       options
	dead       bool
	won        bool
	grid       grid.Grid
	game       gamelogic.PlayField
	termWidth  int
	termHeight int
	restart    bool
}

// createModel creates a new model.
func createModel(o options) (model, error) {
	pf, err := gamelogic.New(o.width, o.height, o.mines)
	if err != nil {
		return model{}, err
	}
	return model{
		opts: o,
		grid: grid.NewGrid(o.width, o.height, styles.GridFocused, styles.GridNormal),
		game: pf,
	}, nil
}

// Init initializes BubbleTea.
func (m model) Init() tea.Cmd {
	return nil
}

// Update updates the view.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.termWidth = msg.Width
		m.termHeight = msg.Height
		return m, nil
	case tea.KeyPressMsg:
		if !m.dead && !m.won {
			switch msg.String() {
			case "esc":
				return m, tea.Quit
			case "space", "f":
				m.game.Play(m.grid.CursorX, m.grid.CursorY, true)
				if m.game.HiddenCells() == 0 && m.game.UnmarkedMines() == 0 {
					m.game.RevealCells()
					m.won = true
				}
				return m, nil
			case "enter":
				m.dead = m.game.Play(m.grid.CursorX, m.grid.CursorY, false)
				if !m.dead && m.game.HiddenCells() == 0 && m.game.UnmarkedMines() == 0 {
					m.game.RevealCells()
					m.won = true
				} else if m.dead {
					m.game.RevealCells()
				}
				return m, nil
			}
		} else {
			switch msg.String() {
			case "esc", "n":
				return m, tea.Quit
			case "enter", "y":
				m.restart = true
				return m, tea.Quit
			}
		}
	}
	if !m.dead && !m.won {
		g, cmd := m.grid.Update(msg)
		m.grid = g
		return m, cmd

	}
	return m, nil
}

func (m model) header() string {
	return styles.Title.
		Width(m.termWidth).
		Render("Gomines")
}

func (m model) footer() string {
	return fmt.Sprint(
		"Press ", styles.KeyChar.Render("Esc"), " to exit, ",
		styles.KeyChar.Render("Enter"), " to show a cell, ",
		styles.KeyChar.Render("Space"), ", ", styles.KeyChar.Render("F"),
		": to mark/unmark a mine.\nMove around using ",
		styles.KeyChar.Render(glyphs.GetKey(glyphs.KeyLeft)), ", ",
		styles.KeyChar.Render(glyphs.GetKey(glyphs.KeyDown)), ", ",
		styles.KeyChar.Render(glyphs.GetKey(glyphs.KeyUp)), ", ",
		styles.KeyChar.Render(glyphs.GetKey(glyphs.KeyRight)), " or ",
		styles.KeyChar.Render("h"), ", ",
		styles.KeyChar.Render("j"), ", ",
		styles.KeyChar.Render("k"), ", ",
		styles.KeyChar.Render("l"), ".",
	)
}

func (m model) Message() string {
	if m.dead {
		return styles.MessageLost.Render("You lost! Do you want to play again? ") +
			styles.KeyChar.Render("Y") + "/" + styles.KeyChar.Render("N")
	} else if m.won {
		return styles.MessageWon.Render("You won! Do you want to play again? ") +
			styles.KeyChar.Render("Y") + "/" + styles.KeyChar.Render("N")
	}
	return ""
}

func (m model) statusBar() string {
	ps := glyphs.Get(glyphs.SDoubt)
	if m.dead {
		ps = glyphs.Get(glyphs.SSad)
	} else if m.won {
		ps = glyphs.Get(glyphs.SHappy)
	}
	msg := fmt.Sprintf("%s: %s | Mines remaining: %d | Hidden cells: %d",
		m.opts.playerName, ps, m.game.UnmarkedMines(), m.game.HiddenCells())
	return styles.StatusBar.Width(m.termWidth).Render(msg)
}

// View renders the view.
func (m model) View() tea.View {
	m.grid.Items = asGridItems(m.opts.width, m.opts.height, m.game.VisibleField)
	// Create a view and set AltScreen to true for full screen mode
	m.grid.View()
	gridView := lipgloss.PlaceHorizontal(m.termWidth, lipgloss.Center, m.grid.View())
	msgView := lipgloss.PlaceHorizontal(m.termWidth, lipgloss.Center, m.Message())
	mainContent := lipgloss.PlaceVertical(m.termHeight-3, lipgloss.Top, joinRows(m.header(), gridView, msgView))
	v := tea.NewView(joinRows(mainContent, m.footer(), m.statusBar()))
	v.AltScreen = true
	m.grid.View()
	return v
}

func joinRows(a ...string) string {
	j := ""
	for i, s := range a {
		if i > 0 {
			j += "\n"
		}
		j += s
	}
	return j
}
