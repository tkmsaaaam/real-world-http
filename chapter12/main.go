package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "Hello World")
	})
	log.Println("Start running server at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
