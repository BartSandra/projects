package main_test

import (
	"banner-service/pkg/config"
	"banner-service/pkg/db"
	"banner-service/pkg/handlers"
	"banner-service/pkg/middleware"
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
	"log"
)

func SetupRouter() *mux.Router {
	cfg := config.Load() 

	dbConn, err := db.Initialize(cfg)
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}

	router := mux.NewRouter()
	router.Use(middleware.JWTAuthentication(cfg))
	router.HandleFunc("/signin", handlers.Signin).Methods("POST")
	router.HandleFunc("/banner", handlers.CreateBanner(dbConn)).Methods("POST")
	router.HandleFunc("/banner/{id}", handlers.UpdateBanner(dbConn)).Methods("PATCH")
	router.HandleFunc("/banner/{id}", handlers.DeleteBanner(dbConn)).Methods("DELETE")
	router.HandleFunc("/user_banner", handlers.GetUserBanner(dbConn)).Methods("GET")
	router.HandleFunc("/banner", handlers.GetBanners(dbConn)).Methods("GET")

	return router
}

func TestE2E2(t *testing.T) {
	router := SetupRouter()

	server := httptest.NewServer(router)
	defer server.Close()

	userCreds := map[string]string{
		"username": "testuser",
		"password": "testpass",
	}
	userBody, _ := json.Marshal(userCreds)
	userResp, err := http.Post(server.URL+"/signin", "application/json", bytes.NewBuffer(userBody))
	if err != nil {
		t.Fatalf("Failed to send POST request to /signin: %v", err)
	}
	defer userResp.Body.Close()

	if userResp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK for signin; got %v", userResp.Status)
	}

}

func TestE2E(t *testing.T) {
	router := SetupRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	testSignIn(t, server)

	testUserAccess(t, server)

	testGetBanner(t, server)
}

func testSignIn(t *testing.T, server *httptest.Server) {
	userCreds := map[string]string{
		"username": "testuser",
		"password": "testpass",
	}
	body, _ := json.Marshal(userCreds)
	resp, err := http.Post(server.URL+"/signin", "application/json", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Failed to send POST request to /signin: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK for signin; got %v", resp.Status)
	}
}

func testUserAccess(t *testing.T, server *httptest.Server) {

	userToken := "validUserTokenHere"
	resp, err := sendAuthorizedRequest("GET", server.URL+"/user_banner?feature_id=1&tag_id=10", userToken, nil)
	if err != nil {
		t.Fatalf("Error sending request with valid user token: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK for valid user access; got %v", resp.Status)
	}
	resp.Body.Close()
}

func testGetBanner(t *testing.T, server *httptest.Server) {
	
	userToken := "validUserTokenHere"
	resp, err := sendAuthorizedRequest("GET", server.URL+"/user_banner?feature_id=1&tag_id=10", userToken, nil)
	if err != nil {
		t.Fatalf("Failed to retrieve banner: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK for retrieving banner; got %v", resp.Status)
	}

}

func sendAuthorizedRequest(method, url, token string, body []byte) (*http.Response, error) {
    client := &http.Client{}
    req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
    if err != nil {
        return nil, err
    }
    req.Header.Set("Authorization", "Bearer " + token)
    req.Header.Set("Content-Type", "application/json")
    return client.Do(req)
}
