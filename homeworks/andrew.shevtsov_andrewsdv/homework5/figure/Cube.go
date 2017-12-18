package figure

import "fmt"

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