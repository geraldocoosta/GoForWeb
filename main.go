package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Handle)
	fmt.Println(http.ListenAndServe(":8080", nil))
}

func Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
