package main

import (
	"fmt"
	"head-first-go-exercises/chapter6/datafile"
	"log"
)

// grab an array of numbers and calculate the average and print it
func main() {
	numbers, err := datafile.GetFloats("./chapter6/steakAverage/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	average := sum / float64(len(numbers))
	fmt.Printf("Average: %.2f\n", average)
}
