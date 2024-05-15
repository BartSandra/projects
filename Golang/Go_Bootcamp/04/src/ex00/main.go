package main

import (
	"encoding/json"
	"net/http"
)

type OrderRequest struct {
	Money      int    `json:"money"`
	CandyType  string `json:"candyType"`
	CandyCount int    `json:"candyCount"`
}

type SuccessResponse struct {
	Thanks string `json:"thanks"`
	Change int    `json:"change"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

var candyPrices = map[string]int{
	"CE": 10,
	"AA": 15,
	"NT": 17,
	"DE": 21,
	"YR": 23,
}

func buyCandyHandler(w http.ResponseWriter, r *http.Request) {
	var req OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	price, exists := candyPrices[req.CandyType]
	if !exists {
		resp := ErrorResponse{Error: "Invalid candy type"}
		json.NewEncoder(w).Encode(resp)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	totalPrice := price * req.CandyCount
	if totalPrice > req.Money {
		resp := ErrorResponse{Error: "You need more money!"}
		json.NewEncoder(w).Encode(resp)
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}

	resp := SuccessResponse{
		Thanks: "Thank you!",
		Change: req.Money - totalPrice,
	}
	json.NewEncoder(w).Encode(resp)
	w.WriteHeader(http.StatusCreated)
}

func main() {
	http.HandleFunc("/buy_candy", buyCandyHandler)
	http.ListenAndServe(":3333", nil)
}
