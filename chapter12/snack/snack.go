package main

import "fmt"

func snack() {
	defer fmt.Println("closing refrigerator")
	fmt.Println("opening refrigerator")
	panic("refrigerator is broken")
}

func main() {
	snack()
}
