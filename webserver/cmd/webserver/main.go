package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("hello!")
	io.WriteString(w, "hello world!")
}

func main() {
	http.HandleFunc("/", getHello)
	log.Fatal(http.ListenAndServe(":1234", nil))
}
