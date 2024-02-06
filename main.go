package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/algorithms"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/characters"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/generation"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/input"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)

const (
	screenWidth  = 1920
	screenHeight = 1080
)

type Game struct {
	buttonsMenu []*input.Button
	buttonsSize []*input.Button
	buttonsAlgo []*input.Button
	buttonBack  *input.Button
	fontFace    font.Face
	Ghosts      characters.Character
	Player      characters.Character
}

const mazeSizeOriginal = 8

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
	g.Ghosts.UpdateCount()

	// Check if the button is clicked
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if g.buttonsMenu[0].Enabled {
			if g.buttonsMenu[0].In(x, y) {
				typeOfMaze = 1
				g.buttonBack.Enabled = true
				input.ChangeStateButtons(g.buttonsSize[:], true)
				input.ChangeStateButtons(g.buttonsAlgo[:], true)
				input.ChangeStateButtons(g.buttonsMenu[:], false)
				return nil

			} else if g.buttonsMenu[1].In(x, y) {
				gameGridDFS = loadFromFile()
				mazeSize = len(gameGridDFS[0])
				changeMazeSize(0, true, g)
			}

		} else if g.buttonBack.Enabled {

			if g.buttonBack.In(x, y) {
				typeOfMaze = 0
				g.buttonBack.Enabled = false
				input.ChangeStateButtons(g.buttonsSize[:], false)
				input.ChangeStateButtons(g.buttonsAlgo[:], false)
				input.ChangeStateButtons(g.buttonsMenu[:], true)
				return nil

			} else if g.buttonsSize[0].In(x, y) {
				changeMazeSize(mazeSizeOriginal, false, g)

			} else if g.buttonsSize[1].In(x, y) {
				changeMazeSize(mazeSizeOriginal*2, false, g)

			} else if g.buttonsSize[2].In(x, y) {
				changeMazeSize((mazeSizeOriginal*2)*2, false, g)

			} else if g.buttonsSize[3].In(x, y) {
				saveToFile(gameGridDFS)

				// A*
			} else if g.buttonsAlgo[0].In(x, y) {
				whichPath = 1

				// Dijkstras
			} else if g.buttonsAlgo[1].In(x, y) {
				whichPath = 0

				// Graph
			} else if g.buttonsAlgo[2].In(x, y) {
				whichPath = 2

				// Shortest Path
			} else if g.buttonsAlgo[3].In(x, y) {
				whichPath = 4

				// Maze Only
			} else if g.buttonsAlgo[4].In(x, y) {
				whichPath = 3
			}

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

		gameMenu(screen, g)
		drawSprite(screen, g.Ghosts)

	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func NewGame() *Game {

	ghost := characters.Character{}
	ghost.Init(gameGridDFS[mazeSize/2][mazeSize/2].NodePosition)
	// Initialize the button
	buttonImage := ebiten.NewImage(100, 30)        // Set the size of the button
	buttonImage.Fill(color.RGBA{0, 255, 255, 250}) // Fill with a color

	buttonsMenu := input.MakeMainMenuButtons(screenWidth, screenHeight)

	buttonsSize := input.GameSizeButtons(screenWidth, screenHeight)

	buttonsAlgo := input.GameAlgoButtons(screenWidth, screenHeight)

	buttonBack := &input.Button{
		Image:   buttonImage,
		X:       (screenWidth / 2) - 50, // Position of the button
		Y:       (screenHeight / 2),
		Width:   100,
		Height:  30,
		Message: "Main Menu",
		Enabled: false,
	}

	// Initialize the game.
	return &Game{
		buttonsMenu: buttonsMenu,
		buttonsSize: buttonsSize,
		buttonsAlgo: buttonsAlgo,
		buttonBack:  buttonBack,
		fontFace:    basicfont.Face7x13,
		Ghosts:      ghost,
	}
}

func changeMazeSize(newSize int, loadedMaze bool, g *Game) {
	if g != nil {
		ghost := characters.Character{}
		ghost.Init(gameGridDFS[mazeSize/2][mazeSize/2].NodePosition)
		g.Ghosts = ghost
	}

	if !loadedMaze {
		oldGameGridDFS = algorithms.DFS(newSize, nil)
		algorithms.MarkUnvisited(oldGameGridDFS)
		gameGridDFS = algorithms.DFS(newSize, oldGameGridDFS)

	} else {
		newSize = mazeSize
	}

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
	changeMazeSize(mazeSizeOriginal, false, nil)
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Single Agent Maze!")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
