package file

import (
	"encoding/gob"
	"log"
	"os"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

// This function, given an game grid, will save that game grid to a file
func SaveToFile(gameGrid [][]mazegrid.MazeSquare) {
	file, err := os.Create("gameGridDFS.gob")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(gameGrid)
	if err != nil {
		log.Fatalf("Failed to encode gameGridDFS: %v", err)
	}
}

// This function, will return a game grid by decoding the game grid file
func LoadFromFile() [][]mazegrid.MazeSquare {

	var gameGrid [][]mazegrid.MazeSquare

	file, err := os.Open("gameGridDFS.gob")
	if err != nil {
		log.Fatalf("Failed to get file: %v", err)
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&gameGrid)
	if err != nil {
		log.Fatalf("Failed to encode gameGridDFS: %v", err)
	}

	return gameGrid
}
