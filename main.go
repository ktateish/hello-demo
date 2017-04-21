package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	color = "blue"
)

func main() {
	host, err := os.Hostname()
	if err != nil {
		log.Fatal("os.Hostname(): %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<body>\n")
		fmt.Fprintf(w, "<h1 style='color: %s;'>", color)
		fmt.Fprintf(w, "Hello, this is %s</h1>\n", host)
		fmt.Fprintf(w, "</body>\n")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
