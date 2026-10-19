package setup

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/mconfalonieri/gomines/internal/pkg/styles"
)

// model is the BubbleTea model.
type model struct {
	opt      *Options
	focused  int
	fields   []field
	errorMsg string
}

// Next switches to the next field.
func (s model) Next() (model, tea.Cmd) {
	focus := s.focused + 1
	if focus == len(s.fields) {
		return s, tea.Quit
	}
	s.focused = focus
	return s, nil
}

// Select selects a specific field.
func (s model) Select(i int) model {
	n := s
	return n
}

// FocusedField returns the focused field.
func (s model) FocusedField() field {
	return s.fields[s.focused]
}

// Init initializes the model.
func (s model) Init() tea.Cmd {
	return nil
}

// SetError sets the error message.
func (s *model) SetError(e error) {
	s.errorMsg = "ERROR: " + e.Error()
}

// ClearError clears the error message.
func (s *model) ClearError() {
	s.errorMsg = ""
}

// Update updates the state.
func (s model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	s.ClearError()
	switch s.FocusedField().Type {
	case fieldTextInput:
		return updateTextInput(s, msg)
	case fieldList:
		return updateList(s, msg)
	default:
		e := fmt.Errorf("unknown field type %d", s.FocusedField().Type)
		s.SetError(e)
		return s, tea.Quit
	}
}

// header prints the header.
func (s model) header() string { return s.FocusedField().Header }

// footer prints the footer.
func (s model) footer() string {
	return styles.Error.Render(s.errorMsg) +
		"\n(Press " + styles.KeyChar.Render("Enter") + " to accept the value and " +
		styles.KeyChar.Render("Esc") + " to quit)"
}

func componentView(f field) string {
	switch f.Type {
	case fieldTextInput:
		return f.TextInputModel.View()
	case fieldList:
		return f.ListModel.View()
	}
	return ""
}

// View renders the view.
func (s model) View() tea.View {
	t := "\n" + styles.Title.Width(len(s.header())).Render("Gomines setup") + "\n"

	t += s.header() + "\n\n"
	t += componentView(s.FocusedField()) + "\n\n"
	t += s.footer()

	// Create a view
	v := tea.NewView(t)
	return v
}

// updateTextInput updates a textinput field.
func updateTextInput(s model, msg tea.Msg) (tea.Model, tea.Cmd) {
	idx := s.focused
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "enter":
			value := s.FocusedField().TextInputModel.Value()
			if value == "" {
				value = s.FocusedField().TextInputModel.Placeholder
			}
			check := s.FocusedField().Check
			if check != nil {
				if err := check(*s.opt, value); err != nil {
					s.SetError(err)
					return s, nil
				}
			}
			s.opt.setValue(idx, value)
			return s.Next()
		case "esc":
			s.opt.quit = true
			return s, tea.Quit
		}
	}
	m := s.FocusedField().TextInputModel
	m, cmd := m.Update(msg)
	s.fields[idx].TextInputModel = m
	return s, cmd
}

// updateList updates a list field.
func updateList(s model, msg tea.Msg) (tea.Model, tea.Cmd) {
	idx := s.focused
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "esc":
			s.opt.quit = true
			return s, tea.Quit
		case "enter":
			m := s.FocusedField().ListModel
			s.opt.setValue(idx, m.SelectedItem().(ListItem).Value())
			return s.Next()
		}
	}
	m := s.FocusedField().ListModel
	m, cmd := m.Update(msg)
	s.fields[idx].ListModel = m
	return s, cmd
}
