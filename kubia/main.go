package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func Home(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "You've hit %s v2\n", hostname)
	return
}

func main() {
	http.HandleFunc("/", Home)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
