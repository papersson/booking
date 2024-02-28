package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUser(t *testing.T) {
	// Setup
	user := User{
		Name:  "John Doe",
		Email: "john@example.com",
	}
	userJSON, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("Unable to marshal JSON: %v", err)
	}

	req, err := http.NewRequest("POST", "/users/create", bytes.NewBuffer(userJSON))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Execute
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUser)
	handler.ServeHTTP(rr, req)

	// Verify
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var createdUser User
	if err := json.NewDecoder(rr.Body).Decode(&createdUser); err != nil {
		t.Fatalf("Could not decode response: %v", err)
	}

	if createdUser.Name != user.Name || createdUser.Email != user.Email {
		t.Errorf("Handler returned unexpected body: got name=%v email=%v, want name=%v email=%v",
			createdUser.Name, createdUser.Email, user.Name, user.Email)
	}

	// Assuming your CreateUser handler assigns a UUID, check if it's not empty
	if createdUser.ID == "" {
		t.Errorf("Expected non-empty user ID")
	}
}
