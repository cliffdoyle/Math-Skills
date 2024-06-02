package functions

//Average calculates the average (mean) of a slice of integers.
func Average(data []int) float64 {

	sum := 0

	for _, num := range data {
		sum += num
	}

	return float64(sum) / float64(len(data))
}
