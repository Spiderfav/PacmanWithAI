package file

import (
	"encoding/gob"
	"log"
	"os"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

// This function, given an game grid, will save that game grid to a file.
// Can only save to one file
func SaveToFile(gameGrid [][]mazegrid.MazeSquare) {
	file, err := os.Create("Maze.gob")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(gameGrid)
	if err != nil {
		log.Fatalf("Failed to encode Maze: %v", err)
	}
}

// This function, will return a game grid by decoding the game grid file.
// It will only load one file at a time.
func LoadFromFile() [][]mazegrid.MazeSquare {

	var gameGrid [][]mazegrid.MazeSquare

	file, err := os.Open("Maze.gob")
	if err != nil {
		log.Fatalf("Failed to get file: %v", err)
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&gameGrid)
	if err != nil {
		log.Fatalf("Failed to encode Maze: %v", err)
	}

	return gameGrid
}
