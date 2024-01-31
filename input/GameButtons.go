package input

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var buttonImage = ebiten.NewImage(100, 30) // Set the size of the button

func MakeMainMenuButtons(screenWidth, screenHeight int) []*Button {
	// Initialize the button
	buttonImage.Fill(color.RGBA{0, 255, 255, 250}) // Fill with a color

	button := &Button{
		Image:   buttonImage,
		X:       (screenWidth / 2) - 50, // Position of the button
		Y:       (screenHeight / 2),
		Width:   100,
		Height:  30,
		Message: "Start Game",
		Enabled: true,
	}

	buttonImport := &Button{
		Image:   buttonImage,
		X:       (screenWidth / 2) - 50, // Position of the button
		Y:       (screenHeight / 2) + 50,
		Width:   100,
		Height:  30,
		Message: "Import Map",
		Enabled: true,
	}

	var menuButtons = []*Button{button, buttonImport}

	return menuButtons
}

func ChangeStateButtons(arrButtons []*Button, state bool) {

	for i := 0; i < len(arrButtons); i++ {
		arrButtons[i].Enabled = state
	}
}

func GameSizeButtons(screenWidth, screenHeight int) []*Button {
	// Initialize the button
	buttonImage.Fill(color.RGBA{0, 255, 255, 250}) // Fill with a color

	buttonSmall := &Button{
		Image:   buttonImage,
		X:       (screenWidth / 2) - 200, // Position of the button
		Y:       (screenHeight / 2) + 50,
		Width:   100,
		Height:  30,
		Message: "Small Map",
		Enabled: true,
	}

	buttonMedium := &Button{
		Image:   buttonImage,
		X:       (screenWidth / 2) - 200, // Position of the button
		Y:       (screenHeight / 2) + 100,
		Width:   100,
		Height:  30,
		Message: "Medium Map",
		Enabled: true,
	}

	buttonLarge := &Button{
		Image:   buttonImage,
		X:       (screenWidth / 2) - 200, // Position of the button
		Y:       (screenHeight / 2) + 150,
		Width:   100,
		Height:  30,
		Message: "Large Map",
		Enabled: true,
	}

	buttonSaveMap := &Button{
		Image:   buttonImage,
		X:       (screenWidth / 2) - 200, // Position of the button
		Y:       (screenHeight / 2) + 200,
		Width:   100,
		Height:  30,
		Message: "Save Map",
		Enabled: true,
	}

	var sizeButtons = []*Button{buttonSmall, buttonMedium, buttonLarge, buttonSaveMap}

	return sizeButtons
}

func GameAlgoButtons(screenWidth, screenHeight int) []*Button {
	// Initialize the button
	buttonImage.Fill(color.RGBA{0, 255, 255, 250}) // Fill with a color

	buttonAStar := &Button{
		Image:   buttonImage,
		X:       (screenWidth / 2) - 20, // Position of the button
		Y:       (screenHeight / 2) + 50,
		Width:   100,
		Height:  30,
		Message: "A*",
		Enabled: true,
	}

	buttonDij := &Button{
		Image:   buttonImage,
		X:       (screenWidth / 2) - 20, // Position of the button
		Y:       (screenHeight / 2) + 100,
		Width:   100,
		Height:  30,
		Message: "Dijkstras",
		Enabled: true,
	}

	buttonNodes := &Button{
		Image:   buttonImage,
		X:       (screenWidth / 2) - 20, // Position of the button
		Y:       (screenHeight / 2) + 150,
		Width:   100,
		Height:  30,
		Message: "Nodes",
		Enabled: true,
	}

	buttonPath := &Button{
		Image:   buttonImage,
		X:       (screenWidth / 2) - 20, // Position of the button
		Y:       (screenHeight / 2) + 200,
		Width:   100,
		Height:  30,
		Message: "Path",
		Enabled: true,
	}

	buttonNone := &Button{
		Image:   buttonImage,
		X:       (screenWidth / 2) - 20, // Position of the button
		Y:       (screenHeight / 2) + 250,
		Width:   100,
		Height:  30,
		Message: "Maze Only",
		Enabled: true,
	}

	var sizeButtons = []*Button{buttonAStar, buttonDij, buttonNodes, buttonPath, buttonNone}

	return sizeButtons
}
