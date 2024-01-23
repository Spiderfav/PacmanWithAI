package main

import (
	"encoding/gob"
	"log"
	"os"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

func saveToFile(gameGrid [][]mazegrid.MazeSquare) {
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

func loadFromFile() [][]mazegrid.MazeSquare {

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
