package shapes

import (
	"math"
	"strconv"
)

type Cylinder struct {
	s string
	m int
	r int
	h int
}

func NewCylinder(conf map[string]string) (Shaper, error) {
	r, _ := strconv.Atoi(conf["R"])
	h, _ := strconv.Atoi(conf["H"])
	x, _ := strconv.Atoi(conf["X"])
	if r == 0 && h == 0 && x > 0 {
		r = x / 2
		h = x
	}
	v := math.Pi * float64(r*r*h)
	m := int(K * v)
	return &Cylinder{
			s: "Cylinder",
			m: m,
			r: r,
			h: h},
		nil
}

func (s *Cylinder) Shape() string {
	return s.s
}

func (s *Cylinder) Weight() int {
	return s.m
}

func (s *Cylinder) String() string {
	return "(" + s.Shape() + " m=" + strconv.Itoa(s.Weight()) + ")"
}
