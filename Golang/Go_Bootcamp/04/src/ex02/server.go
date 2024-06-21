package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

unsigned int i;
unsigned int argscharcount = 0;

char *ask_cow(char phrase[]) {
  int phrase_len = strlen(phrase);
  char *buf = (char *)malloc(sizeof(char) * (160 + (phrase_len + 2) * 3));
  strcpy(buf, " ");

  for (i = 0; i < phrase_len + 2; ++i) {
    strcat(buf, "_");
  }

  strcat(buf, "\n< ");
  strcat(buf, phrase);
  strcat(buf, " ");
  strcat(buf, ">\n ");

  for (i = 0; i < phrase_len + 2; ++i) {
    strcat(buf, "-");
  }
  strcat(buf, "\n");
  strcat(buf, "        \\   ^__^\n");
  strcat(buf, "         \\  (oo)\\_______\n");
  strcat(buf, "            (__)\\       )\\/\\\n");
  strcat(buf, "                ||----w |\n");
  strcat(buf, "                ||     ||\n");
  return buf;
}
*/
import "C"
import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"unsafe"
)

var candyPrices = map[string]int{
	"CE": 10, // Cool Eskimo
	"AA": 15, // Apricot Aardvark
	"NT": 17, // Natural Tiger
	"DE": 21, // Dazzling Elderberry
	"YR": 23, // Yellow Rambutan
}

type Order struct {
	Money      int    `json:"money"`
	CandyType  string `json:"candyType"`
	CandyCount int    `json:"candyCount"`
}

type Response struct {
	Thanks string `json:"thanks,omitempty"`
	Change int    `json:"change,omitempty"`
	Error  string `json:"error,omitempty"`
}

func buyCandy(w http.ResponseWriter, r *http.Request) {
	var order Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid order format", http.StatusBadRequest)
		return
	}

	if order.CandyCount < 0 {
		json.NewEncoder(w).Encode(Response{Error: "Invalid candy count"})
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	price, exists := candyPrices[order.CandyType]
	if !exists {
		json.NewEncoder(w).Encode(Response{Error: "Invalid candy type"})
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	totalPrice := price * order.CandyCount
	if totalPrice > order.Money {
		needed := totalPrice - order.Money
		json.NewEncoder(w).Encode(Response{Error: "You need " + strconv.Itoa(needed) + " more money!"})
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}

	change := order.Money - totalPrice
	phrase := C.CString("Thank you!")
	defer C.free(unsafe.Pointer(phrase))

	cow := C.ask_cow(phrase)
	defer C.free(unsafe.Pointer(cow))

	json.NewEncoder(w).Encode(Response{Thanks: C.GoString(cow), Change: change})
}

func main() {
	http.HandleFunc("/buy_candy", buyCandy)
	err := http.ListenAndServeTLS(":443", "localhost/cert.pem", "localhost/key.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err)
	}
}

