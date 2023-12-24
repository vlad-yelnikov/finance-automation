package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Transaction struct {
	ID              string  `json:"id"`
	Time            int64   `json:"time"`
	Description     string  `json:"description"`
	MCC             int     `json:"mcc"`
	OriginalMCC     int     `json:"originalMcc"`
	Amount          int     `json:"amount"`
	OperationAmount int     `json:"operationAmount"`
	CurrencyCode    int     `json:"currencyCode"`
	CommissionRate  float64 `json:"commissionRate"`
	CashbackAmount  int     `json:"cashbackAmount"`
	Balance         int     `json:"balance"`
	Hold            bool    `json:"hold"`
	ReceiptID       string  `json:"receiptId"`
}

func handleTransaction(w http.ResponseWriter, r *http.Request)  {
	body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		var transaction Transaction
		unmarshalErr := json.Unmarshal(body, &transaction)
		if unmarshalErr != nil {
			log.Println(unmarshalErr)
			return
		}
		fmt.Printf("%+v\n", transaction)
		fmt.Fprint(w, "OK")
}

func router(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprint(w, "OK")
	case http.MethodPost:
		handleTransaction(w, r)
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
