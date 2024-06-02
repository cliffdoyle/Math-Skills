package functions

//Variance calculates the variance of a slice of integers
func Variance(data []int) float64 {

	avg := Average(data)//Calculate the mean of numbers in data
	sum := 0.0//Initialize a variable to keep track of the squared differences from the mean.

	for _, num := range data {

		sum += Square(float64(num) - avg)//Calculate the difference between the number and the average ,then square it and add to sum
	}

	return sum / float64(len(data))//Calculate variance by dividing the sum of Squared differences by the number of elements in data

}
