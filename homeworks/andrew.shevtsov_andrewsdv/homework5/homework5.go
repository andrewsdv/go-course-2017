package main

import (
	"fmt"
	"math/rand"
	"github.com/MastersAcademy/go-course-2017/homeworks/andrew.shevtsov_andrewsdv/homework5/figure"
	"time"
	"sort"
)

type SortableRow []figure.Figure

func (it SortableRow) Len() int {
		return len(it)
	}
func (it SortableRow) Swap(i, j int) {
		it[i], it[j] = it[j], it[i]
	}
func (it SortableRow) Less(i, j int) bool {
		return (it[i]).Weight() > (it[j]).Weight()
	}

var boxSize int
var shakenCorner int
var box [][]figure.Figure
var figures = []figure.Figure{figure.Cube{}, figure.Pyramid{}, figure.Cone{}}

func main() {
	askBoxSize()
	createBox()
	fillBox()
	printBox()
	askShakenCorner()
	sortBox(shakenCorner)
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

func askShakenCorner() {
	fmt.Println()
	fmt.Println("Corners map: ")
	fmt.Println("12")
	fmt.Println("34")
	fmt.Println()

	fmt.Print("Enter which corner to shake around: ")
	fmt.Scanln(&shakenCorner)
	fmt.Println()
}

func sortBox(shakenCorner int) {
	sortableRow := make(SortableRow, boxSize)
	for width := 0; width < boxSize; width = width + 1 {
		for height := 0; height < boxSize; height = height + 1 {
			sortableRow[height] = box[width][height]
		}
		sort.Sort(sortableRow)

		if shakenCorner == 2 || shakenCorner == 4 {
			reverseSlice(sortableRow)
		}

		for height := 0; height < boxSize; height = height + 1 {
			box[width][height] = sortableRow[height]
		}
	}
}
func reverseSlice(sortableRow SortableRow) {
	for i := len(sortableRow)/2 - 1; i >= 0; i-- {
		opp := len(sortableRow) - 1 - i
		sortableRow[i], sortableRow[opp] = sortableRow[opp], sortableRow[i]
	}
}
