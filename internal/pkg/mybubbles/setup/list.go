package setup

import (
	"fmt"
	"strconv"
	"strings"

	"charm.land/bubbles/v2/list"
	"github.com/mconfalonieri/gomines/internal/pkg/glyphs"
)

type ListItem interface {
	Value() string
}

func compileDescription(gSet [glyphs.TotGlyphs]string) string {
	var s strings.Builder
	s.WriteString("|")
	for i, g := range gSet {
		fmt.Fprintf(&s, "% 5s|", g)
		if i%5 == 4 {
			fmt.Fprintf(&s, "\n|")
		}
	}
	return s.String()
}

type glyphItem struct {
	gSet [glyphs.TotGlyphs]string
	name string
	idx  int
	desc string
}

func (gi glyphItem) FilterValue() string {
	return gi.name
}

func (gi glyphItem) Title() string {
	return gi.name
}

func (gi glyphItem) Description() string {
	return gi.desc
}

func (gi glyphItem) Value() string {
	return strconv.Itoa(gi.idx)
}

func createGlyphSetList() list.Model {
	items := []list.Item{
		glyphItem{
			gSet: glyphs.GetSet(glyphs.UNICODE),
			name: "UNICODE",
			idx:  glyphs.UNICODE,
			desc: compileDescription(glyphs.GetSet(glyphs.UNICODE)),
		},
		glyphItem{
			gSet: glyphs.GetSet(glyphs.ASCII),
			name: "ASCII",
			idx:  glyphs.ASCII,
			desc: compileDescription(glyphs.GetSet(glyphs.ASCII)),
		},
	}
	delegate := list.NewDefaultDelegate()
	delegate.SetHeight(4)
	m := list.New(items, delegate, 40, 20)
	m.SetShowHelp(false)
	m.SetShowTitle(false)
	m.SetShowStatusBar(false)
	return m
}
