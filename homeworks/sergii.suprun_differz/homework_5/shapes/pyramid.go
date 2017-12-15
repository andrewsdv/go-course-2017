package shapes

import "strconv"

type Pyramid struct {
	s string
	m int
	a int
	h int
}

func NewPyramid(conf map[string]string) (Shaper, error) {
	a, _ := strconv.Atoi(conf["A"])
	h, _ := strconv.Atoi(conf["H"])
	x, _ := strconv.Atoi(conf["X"])
	if a == 0 && h == 0 && x > 0 {
		a = x 
		h = x
	}
	v := float64(a*a*h) / 3
	m := int(K * v)
	return &Pyramid{
			s: "Pyramid",
			m: m,
			a: a,
			h: h},
		nil
}

func (s *Pyramid) Shape() string {
	return s.s
}

func (s *Pyramid) Weight() int {
	return s.m
}

func (s *Pyramid) String() string {
	return "(" + s.Shape() + " m=" + strconv.Itoa(s.Weight()) + ")"
}
