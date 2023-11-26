package main

import (
	"fmt"
	"log"
	"net/http"
)

func get(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "OK")
}

func main() {
    const port string = "8080"
    http.HandleFunc("/finance", get)
    log.Println("Running on port", port)
    err := http.ListenAndServe("localhost:8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}
