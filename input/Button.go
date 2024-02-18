package input

import "github.com/hajimehoshi/ebiten/v2"

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
