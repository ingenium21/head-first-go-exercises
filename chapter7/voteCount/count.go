package main

import (
	"fmt"
	"head-first-go-exercises/chapter7/datafile"
	"log"
)

func main() {
	lines, err := datafile.GetStrings("./chapter7/voteCount/votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	var names []string
	var counts []int
	for _, line := range lines {
		matched := false
		for i, name := range names {
			if name == line {
				counts[i]++
				matched = true
			}
		}
		if !matched {
			names = append(names, line)
			counts = append(counts, 1)
		}
	}
	for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i])
	}
}
