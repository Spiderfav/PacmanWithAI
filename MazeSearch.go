package main

import (
	"fmt"
	"image/color"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	mplusNormalFont font.Face
	mplusBigFont    font.Face
)

func dijkstras(gameGridDFS *[8][8]MazeSquare, startX int, startY int, finishX int, finishY int) []MazeSquare {

	var pathTaken []MazeSquare

	markUnvisited(gameGridDFS)

	prevWeight := 0
	var nodePrevWeights []int

	var splitNodes []MazeSquare

	gameGridDFS[int(startX/20)-1][int(startY/20)-1].Weight = 0

	// While the bottom right hand corner of the grid does not have a distance assigned and all the paths have not been taken

	for !gameGridDFS[int(finishX/20)-1][int(finishY/20)-1].Visited {

		if gameGridDFS[int(startX/20)-1][int(startY/20)-1] != gameGridDFS[0][0] {
			prevWeight += 1
			gameGridDFS[int(startX/20)-1][int(startY/20)-1].Weight = prevWeight
		}

		gameGridDFS[int(startX/20)-1][int(startY/20)-1].Visited = true
		pathTaken = append(pathTaken, gameGridDFS[int(startX/20)-1][int(startY/20)-1])

		if !gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasDown && !gameGridDFS[int(startX/20)-1+1][int(startY/20)-1].Visited {
			splitNodes = append(splitNodes, gameGridDFS[int(startX/20)-1+1][int(startY/20)-1])
			nodePrevWeights = append(nodePrevWeights, prevWeight)
		}

		if !gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasUp && !gameGridDFS[int(startX/20)-1-1][int(startY/20)-1].Visited {
			splitNodes = append(splitNodes, gameGridDFS[int(startX/20)-1-1][int(startY/20)-1])
			nodePrevWeights = append(nodePrevWeights, prevWeight)

		}

		if !gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasLeft && !gameGridDFS[int(startX/20)-1][int(startY/20)-1-1].Visited {
			splitNodes = append(splitNodes, gameGridDFS[int(startX/20)-1][int(startY/20)-1-1])
			nodePrevWeights = append(nodePrevWeights, prevWeight)

		}

		if !gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasRight && !gameGridDFS[int(startX/20)-1][int(startY/20)-1+1].Visited {
			splitNodes = append(splitNodes, gameGridDFS[int(startX/20)-1][int(startY/20)-1+1])
			nodePrevWeights = append(nodePrevWeights, prevWeight)

		}

		if len(splitNodes) != 0 {
			nodePopped := splitNodes[len(splitNodes)-1]
			splitNodes = splitNodes[:len(splitNodes)-1]

			prevWeight = nodePrevWeights[len(nodePrevWeights)-1]
			nodePrevWeights = nodePrevWeights[:len(nodePrevWeights)-1]

			startY = int(nodePopped.XCoordinate)
			startX = int(nodePopped.YCoordinate)
		}

	}

	return pathTaken
}

func drawDijkstrasOld(screen *ebiten.Image, pathTaken []MazeSquare) {
	fmt.Println(pathTaken)

	prevX := pathTaken[0].XCoordinate + 10
	prevY := pathTaken[0].YCoordinate + 10

	for i := 1; i < len(pathTaken); i++ {
		vector.StrokeLine(screen, prevX, prevY, pathTaken[i].XCoordinate+10, pathTaken[i].YCoordinate+10, 1, color.RGBA{255, 0, 0, 250}, false)
		prevX = pathTaken[i].XCoordinate + 10
		prevY = pathTaken[i].YCoordinate + 10
	}

}

func drawDijkstras(screen *ebiten.Image, pathTaken []MazeSquare) {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	mplusNormalFont, _ = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    8,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})

	for i := 0; i < len(pathTaken); i++ {
		vector.DrawFilledCircle(screen, pathTaken[i].XCoordinate+10, pathTaken[i].YCoordinate+10, 2, color.RGBA{255, 0, 0, 250}, true)
		text.Draw(screen, strconv.Itoa(i), mplusNormalFont, int(pathTaken[i].XCoordinate)+10, int(pathTaken[i].YCoordinate)+10, color.RGBA{255, 0, 255, 250})
	}

}
