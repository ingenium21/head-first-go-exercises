package main

import (
	"fmt"
	"keyboard"
	"log"
)

func main() {
	var status = ""
	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("a grade of", grade, "is", status)
}
