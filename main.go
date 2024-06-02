package main

import (
	"fmt"
	functions "math-skills/Functions"
	"os"
)

// The main function reads a file containing integers and calculates statistical measures such as
// Average, Median, Variance, and Standard deviation then prints the results.
func main() {

	// Check if the correct number of command-line arguments is provided
	if len(os.Args) != 2 {
		fmt.Println("Usage: run main.go <filename.txt>")
		return
	}

	// Get the filename from the command-line arguments
	filename := os.Args[1]
	file := functions.Reader(filename)

	Average := functions.Average(file)
	Median := functions.Median(file)
	Variance := functions.Variance(file)
	StdDeviation := functions.StdDeviation(file)

	fmt.Printf("Average: %0.0f\n", Average)
	fmt.Printf("Median: %0.0f\n", Median)
	fmt.Printf("Variance: %0.0f\n", Variance)
	fmt.Printf("Standard Deviation: %0.0f\n", StdDeviation)
}
