package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kbinani/screenshot"
)

// These variables set the game size to the screen the user is running the game on
var (
	screenWidth  = screenshot.GetDisplayBounds(0).Dx()
	screenHeight = screenshot.GetDisplayBounds(0).Dy()
)

func main() {

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Pacman")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
