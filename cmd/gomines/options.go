package main

import "github.com/mconfalonieri/gomines/internal/pkg/mybubbles/setup"

// options selected during setup phase.
type options struct {
	playerName string
	width      int
	height     int
	mines      int
}

// fromSetup initializes an options object from a setup.Options object.
func fromSetup(o setup.Options) options {
	return options{
		playerName: o.GetPlayerName(),
		width:      o.GetWidth(),
		height:     o.GetHeight(),
		mines:      o.GetMines(),
	}
}
