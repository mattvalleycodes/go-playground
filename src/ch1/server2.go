package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	var mx sync.Mutex
	var count int

	handler := func(w http.ResponseWriter, r *http.Request) {
		mx.Lock()
		count++
		mx.Unlock()

		fmt.Fprintf(w, "url was: %s\n", r.URL.Path)
	}

	countHandler := func(w http.ResponseWriter, r *http.Request) {
		mx.Lock()
		fmt.Fprintf(w, "count is: %d\n", count)
		mx.Unlock()
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/count", countHandler)

	log.Fatal(http.ListenAndServe(":1984", nil))
}