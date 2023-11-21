package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":8080"

func checkHealth(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "OK")
}

func main() {
    http.HandleFunc("/health", checkHealth)
    log.Println("Running on port", port)
    err := http.ListenAndServe(port, nil)
    if err != nil {
        log.Fatal(err)
    }
}
