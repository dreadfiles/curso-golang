package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello! I like %s", request.URL.Path[1:])
	fmt.Fprintf(response, "\n")
	fmt.Fprintf(response, "Hello! I like %s", request.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
