package main

import (
	"fmt"
	"log"
	"net/http"
)

type Counter uint64

func (c *Counter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	*c++
	fmt.Fprintln(w, *c)
}

func main() {
	var c Counter
	http.Handle("/count", &c)

	addr := ":8080"
	log.Printf("server starting on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
