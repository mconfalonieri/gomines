package grid

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/mconfalonieri/gomines/internal/pkg/styles"
)

// GridItem represents an individual cell within the grid.
type GridItem interface {
	Render(width, height int, selected bool) string
}

// Grid Model holds the layout state and items.
type Grid struct {
	Columns int
	Rows    int
	Width   int // Total grid rendering width
	Height  int // Total grid rendering height
	CursorX int
	CursorY int
	Items   [][]GridItem

	// Styles
	FocusedStyle lipgloss.Style
	NormalStyle  lipgloss.Style
}

func NewGrid(cols, rows int, focusedStyle, normalStyle lipgloss.Style) Grid {
	items := make([][]GridItem, rows)
	for i := range items {
		items[i] = make([]GridItem, cols)
	}

	return Grid{
		Columns:      cols,
		Rows:         rows,
		CursorX:      0,
		CursorY:      0,
		Items:        items,
		FocusedStyle: focusedStyle,
		NormalStyle:  normalStyle,
	}
}

func (g *Grid) SetSize(width, height int) {
	g.Width = width
	g.Height = height
}

func (g Grid) Init() tea.Cmd {
	return nil
}

func (g Grid) Update(msg tea.Msg) (Grid, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "left", "h":
			if g.CursorX > 0 {
				g.CursorX--
			}
		case "right", "l":
			if g.CursorX < g.Columns-1 {
				g.CursorX++
			}
		case "up", "k":
			if g.CursorY > 0 {
				g.CursorY--
			}
		case "down", "j":
			if g.CursorY < g.Rows-1 {
				g.CursorY++
			}
		}
	}
	return g, nil
}

func (g Grid) View() string {
	if g.Columns == 0 || g.Rows == 0 {
		return ""
	}

	// Calculate cell dimensions minus outer border padding
	cellWidth := (g.Width / g.Columns) - 2
	cellHeight := (g.Height / g.Rows) - 2

	if cellWidth < 1 {
		cellWidth = 1
	}
	if cellHeight < 1 {
		cellHeight = 1
	}

	var rowViews []string

	for y := 0; y < g.Rows; y++ {
		var colViews []string
		for x := 0; x < g.Columns; x++ {
			isSelected := (x == g.CursorX && y == g.CursorY)

			var cellContent string
			if g.Items[y][x] != nil {
				cellContent = g.Items[y][x].Render(cellWidth, cellHeight, isSelected)
			} else if isSelected {
				cellContent = styles.Prompt.Render(">") + "?" + styles.Prompt.Render("<")
			} else {
				cellContent = " ? "
			}

			// Apply standard border box
			style := g.NormalStyle
			if isSelected {
				style = g.FocusedStyle
			}

			cellBox := style.
				Width(cellWidth).
				Height(cellHeight).
				Render(cellContent)

			colViews = append(colViews, cellBox)
		}
		// Join cells horizontally into a single row
		rowViews = append(rowViews, lipgloss.JoinHorizontal(lipgloss.Top, colViews...))
	}

	// Join all rows vertically to form the grid
	return lipgloss.JoinVertical(lipgloss.Left, rowViews...)
}
