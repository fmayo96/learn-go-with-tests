package structsmethodsinterfaces

import "math"

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Rectangle struct {
	Base   float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Base + 2*r.Height
}

func (r Rectangle) Area() float64 {
	return r.Base * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2.0)
}
