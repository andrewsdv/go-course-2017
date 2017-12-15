package shapes

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

type ShapeFactory func(conf map[string]string) (Shaper, error)

var shapeFactories = make(map[string]ShapeFactory)

func Register(name string, factory ShapeFactory) {
	if factory == nil {
		log.Panicf("Shape factory %s does not exist.", name)
	}
	_, registered := shapeFactories[name]
	if registered {
		log.Printf("Shape factory %s already registered. Ignoring.", name)
	}
	shapeFactories[name] = factory
}

func init() {
	Register("cone", NewCone)
	Register("cube", NewCube)
	Register("cylinder", NewCylinder)
	Register("prizm", NewPrizm)
	Register("pyramid", NewPyramid)
	Register("sphere", NewSphere)
}

func Create(conf map[string]string) (Shaper, error) {
	shapeName := conf["SHAPE"]
	shapeFactory, ok := shapeFactories[shapeName]
	if !ok {
		return nil, fmt.Errorf("Invalid shape name. Must be one of: %s", strings.Join(GetAvailable(), ", "))
	}
	return shapeFactory(conf)
}

func GetAvailable() []string {
	s := []string{}
	for k := range shapeFactories {
		s = append(s, k)
	}
	sort.Sort(sort.StringSlice(s))
	return s
}
