package main

import (
	"fmt"
	"head-first-go-exercises/chapter7/datafile"
	"log"
	"sort"
	"strings"
)

// main is the entry point for the application.
//using maps to store the names and counts
func main() {
	counts := make(map[string]int)
	lines, err := datafile.GetStrings("./chapter7/voteCount/votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range lines {
		//all caps the line to make it easier to compare
		counts[strings.ToUpper(line)]++
	}
	var names []string
	for name := range counts {
		names = append(names, name)
	}
	// sort the names
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("Votes for %s: %d\n", name, counts[name])
	}
}
