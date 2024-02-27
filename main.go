package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"time"
)

type MonoTransaction struct {
	ID              string `json:"id"`
	Time            int64  `json:"time"`
	Description     string `json:"description"`
	MCC             int    `json:"mcc"`
	OriginalMCC     int    `json:"originalMcc"`
	Amount          int    `json:"amount"`
	OperationAmount int    `json:"operationAmount"`
	CurrencyCode    int    `json:"currencyCode"`
	CommissionRate  int    `json:"commissionRate"`
	CashbackAmount  int    `json:"cashbackAmount"`
	Balance         int    `json:"balance"`
	Hold            bool   `json:"hold"`
	ReceiptID       string `json:"receiptId"`
}

type ProcessedTransaction struct {
	Time        string  `json:"time"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	MCC         int     `json:"mcc"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
}

func processTransaction(transaction MonoTransaction) (ProcessedTransaction, bool) {
	if transaction.Amount >= 0 {
		return ProcessedTransaction{}, false
	}
	time := time.Now().Format("2006-01-02T15:04:05") // time should be in my timezone
	amount := math.Abs(float64(transaction.Amount) / 100.0)
	currencyCodes := map[int]string{
		980: "UAH",
		840: "USD",
		978: "EUR",
		985: "PLN",
		710: "ZAR",
	}
	currency := currencyCodes[transaction.CurrencyCode]
	descriptions := map[string]string{
		"Аврора":  "aurora",
		"YouTube": "services",
		"ФОП Волошина Надія Олександрівна": "treatment",
		"Інтернет і ТБ": "services",
	}
	categoryByDescription, exists := descriptions[transaction.Description]
	if exists {
		return ProcessedTransaction{
			Time:        time,
			Description: transaction.Description,
			Category:    categoryByDescription,
			MCC:         transaction.MCC,
			Amount:      amount,
			Currency:    currency,
		}, true
	}
	categories := map[int]string{
		5812: "restaurants",
		5814: "restaurants",
		5816: "games",
		5499: "food",
		5462: "food",
		5441: "food",
		5411: "food",
		// 4900: "apartment bills", temporally disabled
		5399: "aurora",
		5977: "aurora",
		2842: "aurora",
		4112: "travel",
		4131: "transport",
		4121: "transport",
		5912: "treatment",
		4814: "services",
		4816: "services",
		7841: "services",
		5631: "киця",
		5992: "киця",
		5942: "books",
		5651: "clothes/shoes",
		5699: "clothes/shoes",
		8398: "donations",
		7832: "etc",
		7922: "etc",
	}
	category, exists := categories[transaction.MCC]
	if !exists {
		category = "unknown"
	}
	return ProcessedTransaction{
		Time:        time,
		Description: transaction.Description,
		Category:    category,
		MCC:         transaction.MCC,
		Amount:      amount,
		Currency:    currency,
	}, true
}

func sendRequest(transaction ProcessedTransaction) {
	url := ""
	requestBody, err := json.Marshal(transaction)
	if err != nil {
		log.Println(err)
		return
	}
	go http.Post(url, "application/json", bytes.NewBuffer(requestBody))
}

func processPostRequest(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	var transaction MonoTransaction
	unmarshalErr := json.Unmarshal(body, &transaction)
	if unmarshalErr != nil {
		log.Println(unmarshalErr)
		return
	}
	log.Printf("transaction: %+v\n", transaction)
	processedTransaction, processed := processTransaction(transaction)
	if processed {
		sendRequest(processedTransaction)
		fmt.Fprint(w, "OK")
	} else {
		fmt.Fprint(w, "OK")
	}

}

func router(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprint(w, "OK")
	case http.MethodPost:
		processPostRequest(w, r)
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
