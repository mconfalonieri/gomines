package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/mconfalonieri/gomines/internal/pkg/glyphs"
	"github.com/mconfalonieri/gomines/internal/pkg/mybubbles/setup"
	"github.com/mconfalonieri/gomines/internal/pkg/styles"
)

var (
	// Version compiled by goreleaser.
	Version = "dev"
	// Gitsha (commit SHA1) compiled by goreleaser.
	Gitsha = "none"
)

// checkError checks if an error occurred. If err is not nil, it prints a
// message and quits.
func checkError(err error) {
	if err != nil {
		msg := fmt.Sprintf("ERROR: %s", err.Error())
		fmt.Println(styles.Error.Render(msg))
		os.Exit(1)
	}
}

// runSetup runs the setup routine.
func runSetup() setup.Options {
	s, err := setup.Run()
	checkError(err)
	if s.HasQuit() {
		os.Exit(0)
	}
	return s
}

// runProgram runs the game.
func runProgram(o options) {
	run := true
	for run {
		m, err := createModel(o)
		checkError(err)
		p := tea.NewProgram(m)
		fm, err := p.Run()
		checkError(err)
		run = fm.(model).restart
	}
}

func main() {
	if len(os.Args) == 2 && os.Args[1] == "-V" {
		fmt.Printf("Gomines %s (SHA:%s)\n\n", Version, Gitsha)
		os.Exit(0)
	}

	s := runSetup()

	// Set the "font"
	if err := glyphs.SetDefault(s.GetGlyphSet()); err != nil {
		glyphs.SetDefault(glyphs.ASCII)
	}

	// Run the main program
	o := fromSetup(s)
	runProgram(o)
}
