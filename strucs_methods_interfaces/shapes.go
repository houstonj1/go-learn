package strucsmethodsinterfaces

import "math"

// Shape interface
type Shape interface {
	Area() float64
}

// Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter function for Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area function for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle struct
type Circle struct {
	Radius float64
}

// Area function for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle struct
type Triangle struct {
	Base   float64
	Height float64
}

// Area function for Triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
