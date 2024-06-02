package functions

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// Obtain the integers from file
func Reader(filename string) []int {
	//Create an empty slice to store the integers
	var data []int

	//Open the file for reading
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//Create a scanner to read from the file line by line
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() != "" {
			//Convert the line to an integer
			num, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			//Append the integer to the slice
			data = append(data, num)
		}
	}

	//Check if there was any error scanning the file
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	//Return the slice containing the integers from the file
	return data
}
