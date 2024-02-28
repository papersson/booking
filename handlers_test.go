package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUsersHandler(t *testing.T) {
	// Testing POST to create a new user
	var jsonStr = []byte(`{"name":"Test User","email":"test@example.com"}`)
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(usersHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Add more tests as needed, for example, testing the GET method,
	// or testing the userHandler for different scenarios.
}
