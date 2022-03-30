package main

import (
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello, World!")
}

func spanishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hola, Mundo!")
}

func frenchHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Bonjour, le monde!")
}

func hindiHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Namaste, World!")
}

func main() {
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/hola", spanishHandler)
	http.HandleFunc("/salut", frenchHandler)
	http.HandleFunc("/namaste", hindiHandler)

	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
