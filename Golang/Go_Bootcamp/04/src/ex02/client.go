package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
)

type Order1 struct {
	Money      int    `json:"money"`
	CandyType  string `json:"candyType"`
	CandyCount int    `json:"candyCount"`
}

type Response1 struct {
	Thanks string `json:"thanks,omitempty"`
	Change int    `json:"change,omitempty"`
	Error  string `json:"error,omitempty"`
}

func main() {
	log.SetFlags(0)
	candyType := flag.String("k", "", "Candy type")
	candyCount := flag.Int("c", 0, "Candy count")
	money := flag.Int("m", 0, "Money inserted")
	flag.Parse()

	// Чтение CA файла
	caCert, err := ioutil.ReadFile("minica.pem")
	if err != nil {
		log.Fatalf("Reading CA certificate failed: %s", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Настройка TLS
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}
	tlsConfig.BuildNameToCertificate()

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	order := Order1{
		Money:      *money,
		CandyType:  *candyType,
		CandyCount: *candyCount,
	}
	orderJSON, err := json.Marshal(order)
	if err != nil {
		log.Fatalf("JSON marshalling failed: %s", err)
	}

	resp, err := client.Post("https://localhost/buy_candy", "application/json", bytes.NewBuffer(orderJSON))
	if err != nil {
		log.Fatalf("Failed to send request: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response: %s", err)
	}
	var response Response1
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalf("Failed to unmarshal response: %s", err)
	}
	if response.Thanks != "" {
		log.Printf("%s Your change is %d", response.Thanks, response.Change)
	} else {
		log.Printf("Error: %s", response.Error)
	}
}
