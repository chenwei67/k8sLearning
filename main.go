package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received from %s\n", r.RemoteAddr)
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("err: %s", err)
	}
	_, _ = fmt.Fprintf(w, "You've hit %s", hostname)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}