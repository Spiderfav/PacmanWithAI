package main

import (
	"context"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/algorithms"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/characters"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/input"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)

type Maze struct {
	Size int
	Grid [][]mazegrid.MazeSquare
}

const (
	screenWidth  = 1920
	screenHeight = 1080
)

type Game struct {
	Maze        Maze
	buttonsMenu []*input.Button
	buttonsSize []*input.Button
	buttonsAlgo []*input.Button
	buttonBack  *input.Button
	fontFace    font.Face
	Ghosts      characters.NPC
	Player      characters.Character
}

const mazeSizeOriginal = 8

// var whichPath = 3

var menuOrGame = 0

func (g *Game) Update() error {
	g.Ghosts.UpdateCount()

	if menuOrGame == 1 {
		g.Ghosts.Move(g.Player.GetPosition(), g.Maze.Grid)
	}
	// Check if the button is clicked
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if g.buttonsMenu[0].Enabled {
			if g.buttonsMenu[0].In(x, y) {
				menuOrGame = 1
				g.buttonBack.Enabled = true
				input.ChangeStateButtons(g.buttonsSize[:], true)
				input.ChangeStateButtons(g.buttonsAlgo[:], true)
				input.ChangeStateButtons(g.buttonsMenu[:], false)
				return nil

			} else if g.buttonsMenu[1].In(x, y) {
				g.Maze.Grid = loadFromFile()
				g.Maze.Size = len(g.Maze.Grid[0])
				changeMazeSize(0, true, g)
			}

		} else if g.buttonBack.Enabled {

			if g.buttonBack.In(x, y) {
				menuOrGame = 0
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
				saveToFile(g.Maze.Grid)

				// A*
				// } else if g.buttonsAlgo[0].In(x, y) {
				// 	// Change the ghost's algorithm
				// 	//TO-DO: Reset function for Ghosts
				// 	whichPath = 1

				// 	// Dijkstras
				// } else if g.buttonsAlgo[1].In(x, y) {
				// 	whichPath = 0

				// 	// Graph
				// } else if g.buttonsAlgo[2].In(x, y) {
				// 	whichPath = 2

				// 	// Shortest Path
				// } else if g.buttonsAlgo[3].In(x, y) {
				// 	whichPath = 4

				// 	// Maze Only
				// } else if g.buttonsAlgo[4].In(x, y) {
				// 	whichPath = 3
				// }
			}

		}
	}

	return nil
}

// This function is called every second to update what is drawn on the screen

func (g *Game) Draw(screen *ebiten.Image) {

	switch menuOrGame {
	case 0:
		mainMenu(screen, g)

	case 1:
		gameMenu(screen, g)
		characters.DrawSprite(screen, g.Ghosts.Attributes)
		characters.DrawSprite(screen, g.Player)
		drawPathsLines(screen, g.Ghosts.Path)

	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func NewGame() *Game {
	oldGameGridDFS := algorithms.DFS(mazeSizeOriginal, nil)
	algorithms.MarkUnvisited(oldGameGridDFS)
	gameGridDFS := algorithms.DFS(mazeSizeOriginal, oldGameGridDFS)
	maze := Maze{mazeSizeOriginal, gameGridDFS}

	pacman := characters.Character{}
	pacman.Init(gameGridDFS[0][0].NodePosition)

	ghost := characters.NPC{}
	ghost.Init(gameGridDFS[mazeSizeOriginal/2][mazeSizeOriginal/2].NodePosition, algorithms.AStarAlgo, pacman.GetPosition(), gameGridDFS)

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
		Maze:        maze,
		Player:      pacman,
	}
}

func changeMazeSize(newSize int, loadedMaze bool, g *Game) {

	if !loadedMaze {
		oldGameGridDFS := algorithms.DFS(newSize, nil)
		algorithms.MarkUnvisited(oldGameGridDFS)
		g.Maze.Grid = algorithms.DFS(newSize, oldGameGridDFS)
		g.Maze.Size = newSize

	}

	g.Player.SetPosition(g.Maze.Grid[0][0].NodePosition)

	if g.Ghosts.CancelFunc != nil {
		g.Ghosts.CancelFunc()
	}

	g.Ghosts.Ctx, g.Ghosts.CancelFunc = context.WithCancel(context.Background())

	g.Ghosts.UpdatePosition(g.Maze.Grid[g.Maze.Size/2][g.Maze.Size/2].NodePosition, g.Player.GetPosition(), g.Maze.Grid)

	// whichPath = 3
}

func main() {

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Single Agent Maze!")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
