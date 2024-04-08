package input

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Add a Button struct
type Button struct {
	Image         *ebiten.Image
	X, Y          int
	Width, Height int
	Message       string
	Enabled       bool
}

// In returns true if mouse's (x, y) is in the button, and false otherwise.
func (b *Button) In(x, y int) bool {
	return x >= b.X && x < b.X+b.Width && y >= b.Y && y < b.Y+b.Height
}

func (b *Button) ChangeColour(newColour color.RGBA) {
	var buttonImage = ebiten.NewImage(100, 30) // Set the size of the button

	buttonImage.Fill(newColour)

	b.Image = buttonImage
}

func ResetColours(buttons []*Button) {
	var buttonImage = ebiten.NewImage(100, 30)     // Set the size of the button
	buttonImage.Fill(color.RGBA{0, 255, 255, 250}) // Fill with base color

	for i := range buttons {
		buttons[i].Image = buttonImage
	}

}
