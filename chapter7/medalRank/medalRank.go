package main

import (
	"fmt"
)

func main() {
	ranks := map[string]int{"bronze": 3, "silver": 2, "gold": 1}
	for medal, rank := range ranks {
		fmt.Printf("this %s medal is worth %d\n", medal, rank)
	}
}
