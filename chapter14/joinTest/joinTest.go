package main

import (
	"fmt"
	"head-first-go-exercises/chapter14/prose"
)

func main() {
	phrases := []string{"My parents", "a rodeo clown"}
	fmt.Println("A photo of", prose.JoinWithcommas(phrases))
	phrases = []string{"My parents", "a rodeo clown", "a prize bull"}
	fmt.Println("A photo of", prose.JoinWithcommas(phrases))
}
