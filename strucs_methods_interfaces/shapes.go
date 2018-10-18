package strucsmethodsinterfaces

// Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle struct
type Circle struct {
	Radius float64
}

// Perimeter function
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area function
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
