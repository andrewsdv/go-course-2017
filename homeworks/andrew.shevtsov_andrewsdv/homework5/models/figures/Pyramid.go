package figures

import "fmt"

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