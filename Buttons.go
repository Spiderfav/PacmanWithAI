package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var buttonImage = ebiten.NewImage(100, 30) // Set the size of the button

func makeMainMenuButtons() []*Button {
	// Initialize the button
	buttonImage.Fill(color.RGBA{0, 255, 255, 250}) // Fill with a color

	button := &Button{
		image:   buttonImage,
		x:       (screenWidth / 2) - 50, // Position of the button
		y:       (screenHeight / 2),
		width:   100,
		height:  30,
		message: "Start Game",
		enabled: true,
	}

	buttonImport := &Button{
		image:   buttonImage,
		x:       (screenWidth / 2) - 50, // Position of the button
		y:       (screenHeight / 2) + 50,
		width:   100,
		height:  30,
		message: "Import Map",
		enabled: true,
	}

	var menuButtons = []*Button{button, buttonImport}

	return menuButtons
}

func changeStateButtons(arrButtons []*Button, state bool) {

	for i := 0; i < len(arrButtons); i++ {
		arrButtons[i].enabled = state
	}
}

func gameSizeButtons() []*Button {
	// Initialize the button
	buttonImage.Fill(color.RGBA{0, 255, 255, 250}) // Fill with a color

	buttonSmall := &Button{
		image:   buttonImage,
		x:       (screenWidth / 2) - 200, // Position of the button
		y:       (screenHeight / 2) + 50,
		width:   100,
		height:  30,
		message: "Small Map",
		enabled: true,
	}

	buttonMedium := &Button{
		image:   buttonImage,
		x:       (screenWidth / 2) - 200, // Position of the button
		y:       (screenHeight / 2) + 100,
		width:   100,
		height:  30,
		message: "Medium Map",
		enabled: true,
	}

	buttonLarge := &Button{
		image:   buttonImage,
		x:       (screenWidth / 2) - 200, // Position of the button
		y:       (screenHeight / 2) + 150,
		width:   100,
		height:  30,
		message: "Large Map",
		enabled: true,
	}

	buttonSaveMap := &Button{
		image:   buttonImage,
		x:       (screenWidth / 2) - 200, // Position of the button
		y:       (screenHeight / 2) + 200,
		width:   100,
		height:  30,
		message: "Save Map",
		enabled: true,
	}

	var sizeButtons = []*Button{buttonSmall, buttonMedium, buttonLarge, buttonSaveMap}

	return sizeButtons
}

func gameAlgoButtons() []*Button {
	// Initialize the button
	buttonImage.Fill(color.RGBA{0, 255, 255, 250}) // Fill with a color

	buttonAStar := &Button{
		image:   buttonImage,
		x:       (screenWidth / 2) - 20, // Position of the button
		y:       (screenHeight / 2) + 50,
		width:   100,
		height:  30,
		message: "A*",
		enabled: true,
	}

	buttonDij := &Button{
		image:   buttonImage,
		x:       (screenWidth / 2) - 20, // Position of the button
		y:       (screenHeight / 2) + 100,
		width:   100,
		height:  30,
		message: "Dijkstras",
		enabled: true,
	}

	buttonNodes := &Button{
		image:   buttonImage,
		x:       (screenWidth / 2) - 20, // Position of the button
		y:       (screenHeight / 2) + 150,
		width:   100,
		height:  30,
		message: "Nodes",
		enabled: true,
	}

	buttonPath := &Button{
		image:   buttonImage,
		x:       (screenWidth / 2) - 20, // Position of the button
		y:       (screenHeight / 2) + 200,
		width:   100,
		height:  30,
		message: "Path",
		enabled: true,
	}

	buttonNone := &Button{
		image:   buttonImage,
		x:       (screenWidth / 2) - 20, // Position of the button
		y:       (screenHeight / 2) + 250,
		width:   100,
		height:  30,
		message: "Maze Only",
		enabled: true,
	}

	var sizeButtons = []*Button{buttonAStar, buttonDij, buttonNodes, buttonPath, buttonNone}

	return sizeButtons
}
