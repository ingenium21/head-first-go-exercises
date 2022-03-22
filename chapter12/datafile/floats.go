// Package modules allows reading data samples from files
package datafile

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// GetFloats reads a float64 from each line of a file
func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := OpenFile(fileName)
	if err != nil {
		return numbers, err
	}
	defer CloseFile(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		numbers = append(numbers, number)
	}

	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil
}

//OpenFile opens the file and returns a pointer to it, along with an error
func OpenFile(fileName string) (*os.File, error) {
	fmt.Println("Opening:" + fileName)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return file, nil
}

//CloseFile closes the file
func CloseFile(file *os.File) {
	fmt.Println("Closing:" + file.Name())
	file.Close()
}
