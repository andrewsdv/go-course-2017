package main

import (
	"fmt"
	"math/rand"
)

var boxSize int
var box [][]Figure
var figures = []Figure{Cube{}, Pyramid{}, Cone{}}

type Figure interface {
	Name() string
	Weight() int
}

/* Cube */

type Cube struct {
}

func (cube Cube) Name() string {
	return "Cube   "
}

func (cube Cube) Weight() int {
	return 3
}

func (cube Cube) String() string {
	return fmt.Sprintf("[%v %v kg] ", cube.Name(), cube.Weight())
}

/* Pyramid */

type Pyramid struct {
}

func (pyramid Pyramid) Name() string {
	return "Pyramid"
}

func (pyramid Pyramid) Weight() int {
	return 6
}

func (pyramid Pyramid) String() string {
	return fmt.Sprintf("[%v %v kg] ", pyramid.Name(), pyramid.Weight())
}

/* Cone */

type Cone struct {
}

func (cone Cone) Name() string {
	return "Cone   "
}

func (cone Cone) Weight() int {
	return 9
}

func (cone Cone) String() string {
	return fmt.Sprintf("[%v %v kg] ", cone.Name(), cone.Weight())
}

/* Main */

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
	box = make([][]Figure, boxSize)
	for i := range box {
		box[i] = make([]Figure, boxSize)
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

func getRandomFigure() Figure {
	randomIndex := rand.Intn(cap(figures))
	return figures[randomIndex]
}
