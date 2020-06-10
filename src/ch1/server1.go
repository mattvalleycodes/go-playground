package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "url is: %s\n", r.URL.Path)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":1984", nil))
}