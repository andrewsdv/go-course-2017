package shapes

import "strconv"
import "math"

type Prizm struct {
	s string
	m int
	a int
	h int
}

func NewPrizm(conf map[string]string) (Shaper, error) {
	a, _ := strconv.Atoi(conf["A"])
	h, _ := strconv.Atoi(conf["H"])
	x, _ := strconv.Atoi(conf["X"])
	if a == 0 && h == 0 && x > 0 {
		a = x
		h = x
	}
	v := float64(a*a*h) * math.Sqrt(3) / 4
	m := int(K * v)
	return &Prizm{
			s: "Prizm",
			m: m,
			a: a,
			h: h},
		nil
}

func (s *Prizm) Shape() string {
	return s.s
}

func (s *Prizm) Weight() int {
	return s.m
}

func (s *Prizm) String() string {
	return "(" + s.Shape() + " m=" + strconv.Itoa(s.Weight()) + ")"
}
