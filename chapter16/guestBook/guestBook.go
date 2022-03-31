package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Guestbook struct {
	SignatureCount int
	Signatures     []string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	return lines
}

func createHandler(writer http.ResponseWriter, request *http.Request) {
	signature := request.FormValue("signature")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("C:\\Users\\renat\\Documents\\GitHub\\head-first-go-exercises\\chapter16\\guestBook\\signatures.txt", options, os.FileMode(0600))
	check(err)
	_, err = fmt.Fprintln(file, signature)
	check(err)
	err = file.Close()
	check(err)
	http.Redirect(writer, request, "/guestbook", http.StatusFound)
}

func newHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("C:\\Users\\renat\\Documents\\GitHub\\head-first-go-exercises\\chapter16\\guestBook\\new.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures := getStrings("C:\\Users\\renat\\Documents\\GitHub\\head-first-go-exercises\\chapter16\\guestBook\\signatures.txt")
	guestbook := Guestbook{
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}
	html, err := template.ParseFiles("C:\\Users\\renat\\Documents\\GitHub\\head-first-go-exercises\\chapter16\\guestBook\\view.html")
	check(err)
	err = html.Execute(writer, guestbook)
	check(err)
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
