package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear the screen
	screen.Fill(color.White)
	DrawMaze(screen)

}

func DrawMaze(screen *ebiten.Image) {

	var i float64

	for i = 20; i < 200; i += 20 {
		DrawSquare(screen, i)
		//fmt.Println(sum)
	}
}

func DrawSquare(screen *ebiten.Image, x float64) {
	// ebitenutil.DrawLine(screen, 25, 25, 25, 45, color.Black)
	// ebitenutil.DrawLine(screen, 25, 25, 45, 25, color.Black)
	// ebitenutil.DrawLine(screen, 45, 45, 25, 45, color.Black)
	// ebitenutil.DrawLine(screen, 45, 45, 45, 25, color.Black)

	y := x + 20

	ebitenutil.DrawLine(screen, x, x, x, y, color.Black)
	ebitenutil.DrawLine(screen, x, x, y, x, color.Black)
	ebitenutil.DrawLine(screen, y, y, x, y, color.Black)
	ebitenutil.DrawLine(screen, y, y, y, x, color.Black)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
