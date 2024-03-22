package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/input"
	"golang.org/x/image/font"
)

// This function draws all the buttons to the screen for the main menu
func mainMenu(screen *ebiten.Image, g *Game) {
	// Clear the screen to white
	screen.Fill(color.White)

	text.Draw(screen, "Pacman Game", g.fontFace, (screenWidth/2)-40, (screenHeight/2)-100, color.Black)

	for i := 0; i < len(g.buttonsMenu); i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(g.buttonsMenu[i].X), float64(g.buttonsMenu[i].Y))

		screen.DrawImage(g.buttonsMenu[i].Image, op)

		text.Draw(screen, g.buttonsMenu[i].Message, g.fontFace, g.buttonsMenu[i].X+10, g.buttonsMenu[i].Y+20, color.Black)
	}

}

// This functions, given an array of buttons, draws the buttons to the screen
func drawMenu(screen *ebiten.Image, arr []*input.Button, font font.Face) {

	for i := 0; i < len(arr); i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(arr[i].X), float64(arr[i].Y))

		screen.DrawImage(arr[i].Image, op)

		text.Draw(screen, arr[i].Message, font, arr[i].X+10, arr[i].Y+20, color.Black)
	}
}

// This function draws the game menu to the screen
func gameMenu(screen *ebiten.Image, g *Game) {
	// 	// Clear the screen to white
	screen.Fill(color.Black)
	// 	// Draw the maze to the screen
	drawMaze(screen, g)
	//OldMazeSystem(screen, g)
	backButton(screen, g)
	drawMenu(screen, g.buttonsSize, g.fontFace)
	drawMenu(screen, g.buttonsAlgo, g.fontFace)

}

// This function draws the back button to the screen
func backButton(screen *ebiten.Image, g *Game) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.buttonBack.X), float64(g.buttonBack.Y))
	screen.DrawImage(g.buttonBack.Image, op)

	text.Draw(screen, g.buttonBack.Message, g.fontFace, g.buttonBack.X+10, g.buttonBack.Y+20, color.Black)
}
