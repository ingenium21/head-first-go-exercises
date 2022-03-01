package main

import (
	"fmt"
	"head-first-go-exercises/keyboard"
	"log"
)

func toCelsius(temp float64) {
	celsius := (temp - 32) * 5 / 9
	fmt.Printf("%0.2f degrees Celsius\n", celsius)
}

func toFarenheight(temp float64) {
	farenheit := (temp * 9 / 5) + 32
	fmt.Printf("%0.2f degrees Celsius\n", farenheit)
}

func main() {
	fmt.Print("Enter a temperature: ")
	temp, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Choose a conversion: ")
	fmt.Println("1) Celsius to Farenheit")
	fmt.Println("2) Farenheit to Celsius")
	var choice int
	fmt.Scanln(&choice)
	if choice == 2 {
		toCelsius(temp)
	} else if choice == 1 {
		toFarenheight(temp)
	} else {
		fmt.Println("please enter a valid choice")
	}
}
