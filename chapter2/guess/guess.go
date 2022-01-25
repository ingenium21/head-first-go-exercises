package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("Ive chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	fmt.Println(target)

	reader := bufio.NewReader(os.Stdin)

	for guesses := 0; guesses < 10; guesses++ {
		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Oops.  Your guess was LOW")
		} else if guess > target {
			fmt.Println("Oops.  Your guess was HIGH")
		} else {
			fmt.Println("You guessed right!")
			break
		}
		if guesses == 9 {
			fmt.Println("Sorry, you ran out of guesses :(")
			fmt.Println("the correct answer was: ", target)
		}
	}

}
