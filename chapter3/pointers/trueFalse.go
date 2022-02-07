package main

import "fmt"

func negate(myBooloean *bool) {
	*myBooloean = !*myBooloean
}

func main() {
	truth := true
	negate(&truth)
	fmt.Println(truth)
	lies := false
	negate(&lies)
	fmt.Println(lies)
}
