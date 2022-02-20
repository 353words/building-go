package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	flag.Parse() // Support -h

	resp, err := http.Get("https://quotamiki.appspot.com/quote")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error: %s", resp.Status)
	}
	io.Copy(os.Stdout, resp.Body)
}
