package main

import (
	"fmt"
	"net/http"
)

func rootGateWay(w http.ResponseWriter, r *http.Request) {
	println("Welcome to Chris's homepage!")
}

func defaultGateWay(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello jingjing")
	println("Welcome to Chris's homepage!")
}
func main() {
	http.HandleFunc("/", defaultGateWay)
	http.ListenAndServe(":8080", nil)
}
