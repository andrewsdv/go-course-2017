package main

import (
	"fmt"
	"math/rand"
	"github.com/MastersAcademy/go-course-2017/homeworks/andrew.shevtsov_andrewsdv/homework5/models"
	"time"
	"sort"
	"github.com/MastersAcademy/go-course-2017/homeworks/andrew.shevtsov_andrewsdv/homework5/models/figures"
)

var boxSize int
var shakenCorner int
var box []models.SortableRow
var objects = []models.Figure{figures.Cube{}, figures.Pyramid{}, figures.Cone{}}

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
	box = make([]models.SortableRow, boxSize)
	for i := range box {
		box[i] = make([]models.Figure, boxSize)
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

func getRandomFigure() models.Figure {
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)

	randomIndex := randomizer.Intn(cap(objects))
	return objects[randomIndex]
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
	sortHorizontally(shakenCorner)
	sortVertically(shakenCorner)
}

func sortHorizontally(shakenCorner int) {
	sortableRow := make(models.SortableRow, boxSize)
	for height := 0; height < boxSize; height = height + 1 {
		sortableRow = box[height]
		sort.Sort(sortableRow)

		if shakenCorner == 2 || shakenCorner == 4 {
			reverseSlice(sortableRow)
		}
	}
}

func sortVertically(shakenCorner int) {
	for height := 0; height < boxSize; height = height + 1 {
		sortableColumn := make(models.SortableRow, boxSize)
		for width := 0; width < boxSize; width = width + 1 {
			sortableColumn[width] = box[width][height]
		}

		sort.Sort(sortableColumn)

		if shakenCorner == 3 || shakenCorner == 4 {
			reverseSlice(sortableColumn)
		}

		for row := 0; row < boxSize; row = row + 1 {
			box[row][height] = sortableColumn[row]
		}
	}
}

func reverseSlice(sortableRow models.SortableRow) {
	for i := len(sortableRow)/2 - 1; i >= 0; i-- {
		opp := len(sortableRow) - 1 - i
		sortableRow[i], sortableRow[opp] = sortableRow[opp], sortableRow[i]
	}
}
