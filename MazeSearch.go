package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"sort"
	"strconv"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// Defining the normal font for use in the program
var (
	mplusNormalFont font.Face
)

// This function uses the A* Algorithm to find the shortest path from one node to another in a given maze
func aStar(gameGridDFS *[8][8]MazeSquare, startX int, startY int, finishX int, finishY int) []MazeSquare {

	// Storing the original start values
	originalStartX := startX
	originalStartY := startY

	// Marking every node unvisited
	markUnvisited(gameGridDFS)

	var bestPath []MazeSquare

	prevWeight := 0
	var nodePrevWeights []int

	var splitNodes []MazeSquare

	// Assigning the first node a weight of 0
	gameGridDFS[int(startX/20)-1][int(startY/20)-1].Weight = 0

	for !gameGridDFS[int(finishX/20)-1][int(finishY/20)-1].Visited {

		choosingNodes := make(map[MazeSquare]float64)

		// Assigning a new weight to the current node only if it is not the starting point
		if gameGridDFS[int(startX/20)-1][int(startY/20)-1] != gameGridDFS[int(originalStartX/20)-1][int(originalStartY/20)-1] {
			prevWeight += 1
			gameGridDFS[int(startX/20)-1][int(startY/20)-1].Weight = prevWeight
		}

		// Mark the current node as visited and add the node to the array of nodes for the path taken
		gameGridDFS[int(startX/20)-1][int(startY/20)-1].Visited = true
		bestPath = append(bestPath, gameGridDFS[int(startX/20)-1][int(startY/20)-1])

		// This if block checks if the current node has any neighbours and if so, adds them all sequentially to an array
		// It calculates the distance from the neighbour nodes to the end node
		if !gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasDown && !gameGridDFS[int(startX/20)-1+1][int(startY/20)-1].Visited {
			tempminDistance := euclideanDistance(float64(gameGridDFS[int(startX/20)-1+1][int(startY/20)-1].XCoordinate), float64(gameGridDFS[int(startX/20)-1+1][int(startY/20)-1].YCoordinate), float64(gameGridDFS[int(finishX/20)-1][int(startY/20)-1].XCoordinate), float64(gameGridDFS[int(finishX/20)-1][int(startY/20)-1].YCoordinate))
			choosingNodes[gameGridDFS[int(startX/20)-1+1][int(startY/20)-1]] = tempminDistance

		}

		if !gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasUp && !gameGridDFS[int(startX/20)-1-1][int(startY/20)-1].Visited {
			tempminDistance := euclideanDistance(float64(gameGridDFS[int(startX/20)-1-1][int(startY/20)-1].XCoordinate), float64(gameGridDFS[int(startX/20)-1-1][int(startY/20)-1].YCoordinate), float64(gameGridDFS[int(finishX/20)-1][int(startY/20)-1].XCoordinate), float64(gameGridDFS[int(finishX/20)-1][int(startY/20)-1].YCoordinate))
			choosingNodes[gameGridDFS[int(startX/20)-1-1][int(startY/20)-1]] = tempminDistance

		}

		if !gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasLeft && !gameGridDFS[int(startX/20)-1][int(startY/20)-1-1].Visited {
			tempminDistance := euclideanDistance(float64(gameGridDFS[int(startX/20)-1][int(startY/20)-1-1].XCoordinate), float64(gameGridDFS[int(startX/20)-1][int(startY/20)-1-1].YCoordinate), float64(gameGridDFS[int(finishX/20)-1][int(startY/20)-1].XCoordinate), float64(gameGridDFS[int(finishX/20)-1][int(startY/20)-1].YCoordinate))
			choosingNodes[gameGridDFS[int(startX/20)-1][int(startY/20)-1-1]] = tempminDistance

		}

		if !gameGridDFS[int(startX/20)-1][int(startY/20)-1].HasRight && !gameGridDFS[int(startX/20)-1][int(startY/20)-1+1].Visited {
			tempminDistance := euclideanDistance(float64(gameGridDFS[int(startX/20)-1][int(startY/20)-1+1].XCoordinate), float64(gameGridDFS[int(startX/20)-1][int(startY/20)-1+1].YCoordinate), float64(gameGridDFS[int(finishX/20)-1][int(startY/20)-1].XCoordinate), float64(gameGridDFS[int(finishX/20)-1][int(startY/20)-1].YCoordinate))
			choosingNodes[gameGridDFS[int(startX/20)-1][int(startY/20)-1+1]] = tempminDistance

		}

		keys := make([]MazeSquare, 0, len(choosingNodes))

		// The neighbouring nodes are added to a map and then sorted by the distances
		for key := range choosingNodes {
			keys = append(keys, key)
		}

		// This sorts the ( [Node] = Distance ) from highest distance to lowest distance
		sort.SliceStable(keys, func(i, j int) bool {
			return choosingNodes[keys[i]] > choosingNodes[keys[j]]
		})

		// This is adding the sorted nodes back to the array to check for all paths possible
		// This way, the shortest distance nodes are checked first and then the highest distance checked later
		for i := 0; i < len(keys); i++ {
			k := keys[i]
			splitNodes = append(splitNodes, gameGridDFS[int(k.YCoordinate/20)-1][int(k.XCoordinate/20)-1])
			nodePrevWeights = append(nodePrevWeights, prevWeight)
		}

		// If no path was possible from the current node, try a previous found neighbour of a node and set that as the new start
		if len(splitNodes) != 0 {
			nodePopped := splitNodes[len(splitNodes)-1]
			splitNodes = splitNodes[:len(splitNodes)-1]

			prevWeight = nodePrevWeights[len(nodePrevWeights)-1]
			nodePrevWeights = nodePrevWeights[:len(nodePrevWeights)-1]

			startY = int(nodePopped.XCoordinate)
			startX = int(nodePopped.YCoordinate)
		}

	}

	fmt.Println("A* Concluded")
	// Returns the path that the algorithm took to get from the start to the finish
	return bestPath
}

// This function, given the respective x and y values of two nodes, calculates the euclidean distance between them
func euclideanDistance(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	// The euclidean distance is calculated by the square root of the dot product of the difference of the two vectors
	// u = (x1, y1)      v = (x2, y2)     uv = u-v
	// uv . uv = total
	// sqrt(total) = distance

	differenceX := (x2) - (x1)
	differenceY := (y2) - (y1)

	fakeDotProduct := (differenceX * differenceX) + (differenceY * differenceY)

	return math.Sqrt(fakeDotProduct)

}

// This function uses Dijkstras Algorithm to find the shortest path from one node to another in a given maze
func dijkstras(gameGridDFS *[8][8]MazeSquare, startX int, startY int, finishX int, finishY int) []MazeSquare {

	// Storing the original start values
	originalStartX := startX
	originalStartY := startY

	// Marking every node unvisited
	markUnvisited(gameGridDFS)

	var pathTaken []MazeSquare

	prevWeight := 0
	var nodePrevWeights []int

	var splitNodes []MazeSquare

	// Assigning the first node a weight of 0
	gameGridDFS[int(startX/20)-1][int(startY/20)-1].Weight = 0

	// While the end node of the grid has not been visited

	for !gameGridDFS[int(finishX/20)-1][int(finishY/20)-1].Visited {

		// Assigning a new weight to the current node only if it is not the starting point
		if gameGridDFS[int(startX/20)-1][int(startY/20)-1] != gameGridDFS[(originalStartX/20)-1][(originalStartY/20)-1] {
			prevWeight += 1
			gameGridDFS[int(startX/20)-1][int(startY/20)-1].Weight = prevWeight
		}

		// Mark the current node as visited and add the node to the array of nodes for the path taken
		gameGridDFS[int(startX/20)-1][int(startY/20)-1].Visited = true
		pathTaken = append(pathTaken, gameGridDFS[int(startX/20)-1][int(startY/20)-1])

		// This if block checks if the current node has any neighbours and if so, adds them all sequentially to an array
		// It also stores the current weight at the given node for backtracking (that way the weight is correct)
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

		// If no path was possible from the current node, try a previous found neighbour of a node and set that as the new start
		if len(splitNodes) != 0 {
			nodePopped := splitNodes[len(splitNodes)-1]
			splitNodes = splitNodes[:len(splitNodes)-1]

			prevWeight = nodePrevWeights[len(nodePrevWeights)-1]
			nodePrevWeights = nodePrevWeights[:len(nodePrevWeights)-1]

			startY = int(nodePopped.XCoordinate)
			startX = int(nodePopped.YCoordinate)
		}

	}

	fmt.Println("Dijkstra Concluded")

	// Returns the path that the algorithm took to get from the start to the finish
	return pathTaken
}

// This function draws lines for each path taken
func drawPathsLines(screen *ebiten.Image, pathTaken []MazeSquare) {
	prevX := pathTaken[0].XCoordinate + 10
	prevY := pathTaken[0].YCoordinate + 10

	for i := 1; i < len(pathTaken); i++ {
		vector.StrokeLine(screen, prevX, prevY, pathTaken[i].XCoordinate+10, pathTaken[i].YCoordinate+10, 1, color.RGBA{0, 255, 0, 250}, false)
		prevX = pathTaken[i].XCoordinate + 10
		prevY = pathTaken[i].YCoordinate + 10
	}

}

// This function draws circles with their position in the path
// It also draws the start node and end node and the total cost
func drawPaths(screen *ebiten.Image, pathTaken []MazeSquare) {

	// Here we are defining the font to be used from the general golang fonts
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

	// For every node searched by the algorithms, draw a circle with their postion
	for i := 0; i < len(pathTaken); i++ {
		vector.DrawFilledCircle(screen, pathTaken[i].XCoordinate+10, pathTaken[i].YCoordinate+10, 2, color.RGBA{255, 0, 0, 250}, true)
		text.Draw(screen, strconv.Itoa(i), mplusNormalFont, int(pathTaken[i].XCoordinate)+10, int(pathTaken[i].YCoordinate)+10, color.RGBA{255, 0, 255, 250})

	}

	text.Draw(screen, "Start node is "+strconv.Itoa(int(pathTaken[0].XCoordinate))+","+strconv.Itoa(int(pathTaken[0].YCoordinate)), mplusNormalFont, 10, int(gameGridDFS[len(gameGridDFS)-1][len(gameGridDFS)-1].YCoordinate)+40, color.RGBA{0, 0, 0, 250})
	text.Draw(screen, "End node is "+strconv.Itoa(int(pathTaken[len(pathTaken)-1].XCoordinate))+","+strconv.Itoa(int(pathTaken[len(pathTaken)-1].YCoordinate)), mplusNormalFont, 10, int(gameGridDFS[len(gameGridDFS)-1][len(gameGridDFS)-1].YCoordinate)+50, color.RGBA{0, 0, 0, 250})

	text.Draw(screen, "Path cost to desired node is "+strconv.Itoa(int(pathTaken[len(pathTaken)-1].Weight)), mplusNormalFont, 10, 10, color.RGBA{0, 0, 0, 250})
}
