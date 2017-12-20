package figures

import "fmt"

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