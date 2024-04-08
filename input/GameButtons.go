package input

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var buttonImage = ebiten.NewImage(100, 30) // Set the size of the button

// This functions, given an array of buttons and a new state, will change those buttons to the given state
func ChangeStateButtons(arrButtons []*Button, state bool) {

	for i := 0; i < len(arrButtons); i++ {
		arrButtons[i].Enabled = state
	}
}

// This function, takes in the current screen dimensions and returns the game maze size buttons for the given screen size
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

// This function, takes in the current screen dimensions and returns the game change algorithm buttons for the given screen size
func GameAlgoButtons(screenWidth, screenHeight int) []*Button {
	// Initialize the button
	buttonImage.Fill(color.RGBA{0, 255, 255, 250}) // Fill with a color

	buttonAStar := &Button{
		Image:   buttonImage,
		X:       (screenWidth / 2) - 20, // Position of the button
		Y:       (screenHeight / 2) + 50,
		Width:   100,
		Height:  30,
		Message: "Dijkstras",
		Enabled: true,
	}

	buttonDij := &Button{
		Image:   buttonImage,
		X:       (screenWidth / 2) - 20, // Position of the button
		Y:       (screenHeight / 2) + 100,
		Width:   100,
		Height:  30,
		Message: "A*",
		Enabled: true,
	}

	buttonDFS := &Button{
		Image:   buttonImage,
		X:       (screenWidth / 2) - 20, // Position of the button
		Y:       (screenHeight / 2) + 150,
		Width:   100,
		Height:  30,
		Message: "BFS",
		Enabled: true,
	}

	buttonBFS := &Button{
		Image:   buttonImage,
		X:       (screenWidth / 2) - 20, // Position of the button
		Y:       (screenHeight / 2) + 200,
		Width:   100,
		Height:  30,
		Message: "DFS",
		Enabled: true,
	}

	buttonMiniMax := &Button{
		Image:   buttonImage,
		X:       (screenWidth / 2) - 20, // Position of the button
		Y:       (screenHeight / 2) + 250,
		Width:   100,
		Height:  30,
		Message: "MiniMax(Pruned)",
		Enabled: true,
	}

	buttonExpectimax := &Button{
		Image:   buttonImage,
		X:       (screenWidth / 2) - 20, // Position of the button
		Y:       (screenHeight / 2) + 300,
		Width:   100,
		Height:  30,
		Message: "Expectimax",
		Enabled: true,
	}

	var sizeButtons = []*Button{buttonAStar, buttonDij, buttonDFS, buttonBFS, buttonMiniMax, buttonExpectimax}

	return sizeButtons
}

// This function, takes in the current screen dimensions and returns the ghost's buttons to the game
func GameGhostButtons(screenWidth, screenHeight int) []*Button {
	// Initialize the button
	buttonImage.Fill(color.RGBA{0, 255, 255, 250}) // Fill with a color

	buttonAdd := &Button{
		Image:   buttonImage,
		X:       (screenWidth / 2) + 160, // Position of the button
		Y:       (screenHeight / 2) + 50,
		Width:   100,
		Height:  30,
		Message: "Add Ghost",
		Enabled: true,
	}

	buttonRemove := &Button{
		Image:   buttonImage,
		X:       (screenWidth / 2) + 160, // Position of the button
		Y:       (screenHeight / 2) + 100,
		Width:   100,
		Height:  30,
		Message: "Remove Ghost",
		Enabled: true,
	}

	var ghostButtons = []*Button{buttonAdd, buttonRemove}

	return ghostButtons
}
