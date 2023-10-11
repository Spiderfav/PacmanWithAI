package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type MazeSquare struct {
	X_Value                                       float64
	LeftSquare, DownSquare, RightSquare, UpSquare *MazeSquare
}

func (square MazeSquare) DrawSquare(screen *ebiten.Image) {
	y := square.X_Value + 20

	ebitenutil.DrawLine(screen, square.X_Value, square.X_Value, square.X_Value, y, color.Black)
	ebitenutil.DrawLine(screen, square.X_Value, square.X_Value, y, square.X_Value, color.Black)
	ebitenutil.DrawLine(screen, y, y, square.X_Value, y, color.Black)
	ebitenutil.DrawLine(screen, y, y, y, square.X_Value, color.Black)
}

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

	var square1 = MazeSquare{20, nil, nil, nil, nil}

	square1.DrawSquare(screen)

	//var i float64

	// for i = 20; i < 200; i += 20 {
	// 	DrawSquareOld(screen, i)
	// 	//fmt.Println(sum)
	// }
}

func DrawSquareOld(screen *ebiten.Image, x float64) {
	ebitenutil.DrawLine(screen, 25, 25, 25, 45, color.Black)
	ebitenutil.DrawLine(screen, 25, 25, 45, 25, color.Black)
	ebitenutil.DrawLine(screen, 45, 45, 25, 45, color.Black)
	ebitenutil.DrawLine(screen, 45, 45, 45, 25, color.Black)

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
