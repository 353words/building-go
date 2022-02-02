package main

import (
	"fmt"
	"log"
	"net/http"
)

func byHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, Founder)
}

func main() {
	http.HandleFunc("/by", byHandler)

	addr := ":8080"
	log.Printf("server starting on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
