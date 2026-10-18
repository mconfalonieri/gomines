package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/mconfalonieri/gomines/internal/pkg/glyphs"
	"github.com/mconfalonieri/gomines/internal/pkg/mybubbles/setup"
)

func main() {
	// Run setup
	s, err := setup.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	} else if s.HasQuit() {
		os.Exit(0)
	}

	// Main program
	if err := glyphs.SetDefault(s.GetGlyphSet()); err != nil {
		glyphs.SetDefault(glyphs.ASCII)
	}
	run := true
	for run {
		m, err := createModel(fromSetup(s))
		// m, err := createModel(options{playerName: "Test", width: 10, height: 10, mines: 10})
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		p := tea.NewProgram(m)
		fm, err := p.Run()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		run = fm.(model).restart
	}
}
