package main

import "fmt"

type numberP int //number is a type alias

func (n *numberP) double() {
	*n = *n * 2
}

func main() {
	four := numberP(4)
	fmt.Println("Before doubling:", four)
	four.double()
	fmt.Println("After doubling:", four)
}
