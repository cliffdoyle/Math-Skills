package functions

import "sort"

// Median calculates the median of a slice of integers.
func Median(data []int) float64 {
	//Sort the data in ascending order
	sort.Ints(data)

	n := len(data)
	//Check if the count of numbers is even
	if n%2 == 0 {
		//If even, return the average of two middle numbers
		return float64(data[n/2-1]+data[n/2]) / 2.0
	}
	//If odd return the middle number
	return float64(data[n/2])
}
