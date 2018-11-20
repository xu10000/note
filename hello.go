package main

import (
	"fmt"
	"net/http"
	"log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, world")
}


func main() {
	http.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal("ListenAndServe err: ", err)
	}
}
