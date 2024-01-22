package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/algorithms"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/generation"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)

const (
	screenWidth  = 1920
	screenHeight = 1080
)

// Add a Button struct
type Button struct {
	image         *ebiten.Image
	x, y          int
	width, height int
	message       string
	enabled       bool
}

// In returns true if mouse's (x, y) is in the button, and false otherwise.
func (b *Button) In(x, y int) bool {
	return x >= b.x && x < b.x+b.width && y >= b.y && y < b.y+b.height
}

type Game struct {
	buttonsMenu []*Button
	buttonsSize []*Button
	buttonsAlgo []*Button
	buttonBack  *Button
	fontFace    font.Face
}

// MouseStrokeSource is a StrokeSource implementation of mouse.
type MouseStrokeSource struct{}

func (m *MouseStrokeSource) Position() (int, int) {
	return ebiten.CursorPosition()
}

func (m *MouseStrokeSource) IsJustReleased() bool {
	return inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft)
}

var mazeSizeOriginal = 8
var oldGameGridDFS [][]mazegrid.MazeSquare = algorithms.DFS(mazeSizeOriginal, nil)
var gameGridDFS [][]mazegrid.MazeSquare = algorithms.DFS(mazeSizeOriginal, oldGameGridDFS)
var mazeSize = len(gameGridDFS[0])

var dijkstrasPath = algorithms.Dijkstras(gameGridDFS, 20, 20, 20*mazeSizeOriginal, 20*mazeSizeOriginal)
var absolutePathDijkstras, weightDijkstras = algorithms.AbsolutePath(dijkstrasPath)

var aStarPath = algorithms.AStar(gameGridDFS, 20, 20, 20*mazeSizeOriginal, 20*mazeSizeOriginal)
var absolutePathAStar, weigthAStar = algorithms.AbsolutePath(aStarPath)

var graph = generation.MazeToGraph(gameGridDFS, 20, 20, float32(20*mazeSizeOriginal), float32(20*mazeSizeOriginal))
var graphPaths = generation.AllPaths(gameGridDFS, graph)

var whichPath = 3

var typeOfMaze = 0

func (g *Game) Update() error {
	// Check if the button is clicked
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if g.buttonsMenu[0].In(x, y) && g.buttonsMenu[0].enabled {
			typeOfMaze = 1
			g.buttonBack.enabled = true
			changeStateButtons(g.buttonsSize[:], true)
			changeStateButtons(g.buttonsAlgo[:], true)
			changeStateButtons(g.buttonsMenu[:], false)
			return nil

		} else if g.buttonBack.In(x, y) && g.buttonBack.enabled {
			typeOfMaze = 0
			g.buttonBack.enabled = false
			changeStateButtons(g.buttonsSize[:], false)
			changeStateButtons(g.buttonsAlgo[:], false)
			changeStateButtons(g.buttonsMenu[:], true)
			return nil

		} else if g.buttonsSize[0].In(x, y) && g.buttonsSize[0].enabled {
			changeMazeSize(mazeSizeOriginal)

		} else if g.buttonsSize[1].In(x, y) && g.buttonsSize[1].enabled {
			changeMazeSize(mazeSizeOriginal * 2)

		} else if g.buttonsSize[2].In(x, y) && g.buttonsSize[2].enabled {
			changeMazeSize((mazeSizeOriginal * 2) * 2)

		} else if g.buttonsAlgo[0].In(x, y) && g.buttonsAlgo[0].enabled {
			whichPath = 1

		} else if g.buttonsAlgo[1].In(x, y) && g.buttonsAlgo[1].enabled {
			whichPath = 0

		} else if g.buttonsAlgo[2].In(x, y) && g.buttonsAlgo[2].enabled {
			whichPath = 2

		} else if g.buttonsAlgo[3].In(x, y) && g.buttonsAlgo[1].enabled {
			whichPath = 4

		} else if g.buttonsAlgo[4].In(x, y) && g.buttonsAlgo[2].enabled {
			whichPath = 3
		}
	}

	if g.buttonBack.enabled {
		if inpututil.IsKeyJustPressed(ebiten.Key1) {
			p := []int{1, 2, 3, 4}

			saveToFile(p)

		}

		if inpututil.IsKeyJustPressed(ebiten.Key2) {
			grid := loadFromFile()

			fmt.Println(grid)
		}

		// Dijkstras
		if inpututil.IsKeyJustPressed(ebiten.KeyA) {
			whichPath = 0

			// A*
		} else if inpututil.IsKeyJustPressed(ebiten.KeyB) {
			whichPath = 1

			// Graph
		} else if inpututil.IsKeyJustPressed(ebiten.KeyC) {
			whichPath = 2

			// Maze Only
		} else if inpututil.IsKeyJustPressed(ebiten.KeyD) {
			whichPath = 3

			// Shortest Path
		} else if inpututil.IsKeyJustPressed(ebiten.KeyE) {

			whichPath = 4
		}
	}

	return nil
}

// This function is called every second to update what is drawn on the screen

func (g *Game) Draw(screen *ebiten.Image) {

	switch typeOfMaze {
	case 0:
		mainMenu(screen, g)
	case 1:
		OldMazeSystem(screen, g)
		backButton(screen, g)
		sizeMenu(screen, g)
		algoMenu(screen, g)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func NewGame() *Game {

	// Initialize the button
	buttonImage := ebiten.NewImage(100, 30)        // Set the size of the button
	buttonImage.Fill(color.RGBA{0, 255, 255, 250}) // Fill with a color

	buttonsMenu := makeMainMenuButtons()

	buttonsSize := gameSizeButtons()

	buttonsAlgo := gameAlgoButtons()

	buttonBack := &Button{
		image:   buttonImage,
		x:       (screenWidth / 2) - 50, // Position of the button
		y:       (screenHeight / 2),
		width:   100,
		height:  30,
		message: "Main Menu",
		enabled: false,
	}

	// Initialize the game.
	return &Game{
		buttonsMenu: buttonsMenu,
		buttonsSize: buttonsSize,
		buttonsAlgo: buttonsAlgo,
		buttonBack:  buttonBack,
		fontFace:    basicfont.Face7x13,
	}
}

func changeMazeSize(newSize int) {
	oldGameGridDFS = algorithms.DFS(newSize, nil)
	algorithms.MarkUnvisited(oldGameGridDFS)
	gameGridDFS = algorithms.DFS(newSize, oldGameGridDFS)
	dijkstrasPath = algorithms.Dijkstras(gameGridDFS, 20, 20, 20*newSize, 20*newSize)
	aStarPath = algorithms.AStar(gameGridDFS, 20, 20, 20*newSize, 20*newSize)
	absolutePathDijkstras, weightDijkstras = algorithms.AbsolutePath(dijkstrasPath)
	absolutePathAStar, weigthAStar = algorithms.AbsolutePath(aStarPath)
	graph = generation.MazeToGraph(gameGridDFS, 20, 20, float32(20*newSize), float32(20*newSize))
	graphPaths = generation.AllPaths(gameGridDFS, graph)
	mazeSize = newSize
	whichPath = 3
}

func main() {

	fmt.Println("Size of Dijkstras:", len(dijkstrasPath))
	fmt.Println("Size of A*:", len(aStarPath))

	fmt.Println("Size of absolute path", len(absolutePathDijkstras))
	fmt.Println(" ")
	changeMazeSize(mazeSizeOriginal)
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Single Agent Maze!")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
