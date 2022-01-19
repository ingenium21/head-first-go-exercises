package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileinfo, err := os.Stat(".\\chapter2\\exercise2\\my.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileinfo.Size())
}
