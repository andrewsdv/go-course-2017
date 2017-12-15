package shapes

import "strconv"
import "math"

type Sphere struct {
	s string
	m int
	r int
}

func NewSphere(conf map[string]string) (Shaper, error) {
	r, _ := strconv.Atoi(conf["R"])
	x, _ := strconv.Atoi(conf["X"])
	if r == 0 && x > 0 {
		r = x / 2
	}
	v := float64(r*r*r) * math.Pi * 4 / 3
	m := int(K * v)
	return &Sphere{
			s: "Sphere",
			m: m,
			r: r},
		nil
}

func (s *Sphere) Shape() string {
	return s.s
}

func (s *Sphere) Weight() int {
	return s.m
}

func (s *Sphere) String() string {
	return "(" + s.Shape() + " m=" + strconv.Itoa(s.Weight()) + ")"
}
