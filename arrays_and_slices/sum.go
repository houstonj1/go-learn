package slices

// Sum function
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll function
func SumAll(numbersToSum ...[]int) (sums []int) {
	return
}
