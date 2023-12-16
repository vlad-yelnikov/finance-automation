package main

import (
	"fmt"
	"log"
	"net/http"
)

func router(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "OK")
	case http.MethodPost:
		fmt.Fprintf(w, "OK")
	}
}

func main() {
	port := "8080"
	http.HandleFunc("/", router)
	log.Println("Running on port", port)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
