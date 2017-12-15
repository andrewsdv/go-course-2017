package shapes

import "strconv"

type Cube struct {
	s string
	m int
	a int
}

func NewCube(conf map[string]string) (Shaper, error) {
	a, _ := strconv.Atoi(conf["A"])
	x, _ := strconv.Atoi(conf["X"])
	if a == 0 && x > 0 {
		a = x
	}
	v := float64(a * a * a)
	m := int(K * v)
	return &Cube{
			s: "Cube",
			m: m,
			a: a},
		nil
}

func (s *Cube) Shape() string {
	return s.s
}

func (s *Cube) Weight() int {
	return s.m
}

func (s *Cube) String() string {
	return "(" + s.Shape() + " m=" + strconv.Itoa(s.Weight()) + ")"
}
