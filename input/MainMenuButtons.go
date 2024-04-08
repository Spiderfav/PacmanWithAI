package input

import (
	"image/color"
)

// This function, takes in the current screen dimensions and returns the main menu buttons for the given screen size
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

	buttonSettings := &Button{
		Image:   buttonImage,
		X:       (screenWidth / 2) - 50, // Position of the button
		Y:       (screenHeight / 2) + 100,
		Width:   100,
		Height:  30,
		Message: "Settings",
		Enabled: true,
	}

	buttonExit := &Button{
		Image:   buttonImage,
		X:       (screenWidth / 2) - 50, // Position of the button
		Y:       (screenHeight / 2) + 150,
		Width:   100,
		Height:  30,
		Message: "Exit",
		Enabled: true,
	}

	var menuButtons = []*Button{button, buttonImport, buttonSettings, buttonExit}

	return menuButtons
}
