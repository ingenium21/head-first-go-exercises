//takes a variadic number of float64s and returns the average
package main

import "fmt"

func average(numbers ...float64) {
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	average := sum / float64(len(numbers))
	fmt.Printf("Average: %.2f\n", average)
}

func main() {
	average(100, 50)
	average(90.7, 89.7, 98.5, 92.3)
}
