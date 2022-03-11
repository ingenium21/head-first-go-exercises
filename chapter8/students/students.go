package main

import "fmt"

//create student struct
type student struct {
	name  string
	grade float64
}

func printInfo(s student) {
	fmt.Printf("Name: %s\n", s.name)
	fmt.Printf("Grade: %.2f\n", s.grade)
}

func main() {
	//create student
	var s student
	s.name = "Alonzo Cole"
	s.grade = 92.3
	printInfo(s)
}
