package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func responseSize(url string) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(body))

}
func main() {
	go responseSize("https://www.google.com")
	go responseSize("https://www.yahoo.com")
	go responseSize("https://www.bing.com")
	time.Sleep(2 * time.Second)
}
