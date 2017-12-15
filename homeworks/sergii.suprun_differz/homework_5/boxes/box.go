package boxes

// Boxer is common interface for all boxes
type Boxer interface {
	IsEmpty() bool
	GetState() int
	PrintWeight() string
	String() string
}

// Generator can fill in new boxes
type Generator interface {
	Generate()
}

// Shaker makes sorting
type Shaker interface {
	Shake(int)
}

// BlackBoxer combined box
type BlackBoxer interface {
	Boxer
	Generator
	Shaker
}
