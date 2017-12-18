package main

import (
	"fmt"
	"math/rand"
	"github.com/MastersAcademy/go-course-2017/homeworks/andrew.shevtsov_andrewsdv/homework5/figure"
	"time"
)

var boxSize int
var box [][]figure.Figure
var figures = []figure.Figure{figure.Cube{}, figure.Pyramid{}, figure.Cone{}}

func main() {
	askBoxSize()
	createBox()
	fillBox()
	printBox()
}

func askBoxSize() {
	fmt.Print("Enter N value for box size NxN: ")
	fmt.Scanln(&boxSize)
	fmt.Printf("Box size: %dx%d", boxSize, boxSize)
	fmt.Println()
}

func createBox() {
	box = make([][]figure.Figure, boxSize)
	for i := range box {
		box[i] = make([]figure.Figure, boxSize)
	}
}

func fillBox() {
	for width := 0; width < boxSize; width = width + 1 {
		for height := 0; height < boxSize; height = height + 1 {
			box[width][height] = getRandomFigure()
		}
		fmt.Println()
	}
}

func printBox() {
	for width := 0; width < boxSize; width = width + 1 {
		for height := 0; height < boxSize; height = height + 1 {
			fmt.Print(box[width][height])
		}
		fmt.Println()
	}
}

func getRandomFigure() figure.Figure {
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)

	randomIndex := randomizer.Intn(cap(figures))
	return figures[randomIndex]
}
