package main

import "fmt"

type number int //number is a type alias

//these are examples of type methods
func (n number) add(otherNumber int) {
	fmt.Printf("%d + %d = %d\n", n, otherNumber, int(n)+otherNumber)
}

func (n number) subtract(otherNumber int) {
	fmt.Printf("%d - %d = %d\n", n, otherNumber, int(n)-otherNumber)
}

func main() {
	ten := number(10)
	ten.add(4)
	ten.subtract(5)
	four := number(4)
	four.add(3)
	four.subtract(2)
}
