package main

import (
	"fmt"
	"head-first-go-exercises/chapter12/datafile"
	"log"
	"os"
)

func main() {
	numbers, err := datafile.GetFloats(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Printf("Sum: %.2f\n", sum)
}
