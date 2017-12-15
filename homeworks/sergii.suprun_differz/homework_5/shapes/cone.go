package shapes

import (
	"math"
	"strconv"
)

type Cone struct {
	s string
	m int
	r int
	h int
}

func NewCone(conf map[string]string) (Shaper, error) {
	r, _ := strconv.Atoi(conf["R"])
	h, _ := strconv.Atoi(conf["H"])
	x, _ := strconv.Atoi(conf["X"])
	if r == 0 && h == 0 && x > 0 {
		r = x / 2
		h = x
	}
	v := math.Pi * float64(r*r*h) / 3
	m := int(K * v)
	return &Cone{
			s: "Cone",
			m: m,
			r: r,
			h: h},
		nil
}

func (s *Cone) Shape() string {
	return s.s
}

func (s *Cone) Weight() int {
	return s.m
}

func (s *Cone) String() string {
	return "(" + s.Shape() + " m=" + strconv.Itoa(s.Weight()) + ")"
}
