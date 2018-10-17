package strucsmethodsinterfaces

// Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter function
func Perimeter(width float64, height float64) float64 {
	return 2 * (width + height)
}

// Area function
func Area(width float64, height float64) float64 {
	return width * height
}
