package functions

import "math"

// Finds the standard deviation by getting the square root of variance
func StdDeviation(data []int) float64 {

	return math.Sqrt(Variance(data))
}
