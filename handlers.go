package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"

	"github.com/google/uuid"
)

var (
	users = make(map[string]User) // In-memory store for users
	mu    sync.Mutex              // Mutex to avoid race condition
)

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.ID = uuid.New().String() // Generate a unique ID for the user
	mu.Lock()
	users[user.ID] = user
	mu.Unlock()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// ListUsers lists all users
func ListUsers(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	values := make([]User, 0, len(users))
	for _, v := range users {
		values = append(values, v)
	}
	mu.Unlock()
	json.NewEncoder(w).Encode(values)
}

// GetUser retrieves a single user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/user/")
	mu.Lock()
	user, ok := users[id]
	mu.Unlock()
	if !ok {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// UpdateUser updates an existing user's details
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/user/")
	var updateUser User
	if err := json.NewDecoder(r.Body).Decode(&updateUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mu.Lock()
	_, ok := users[id]
	if ok {
		updateUser.ID = id // Ensure the ID remains unchanged
		users[id] = updateUser
	}
	mu.Unlock()
	if ok {
		json.NewEncoder(w).Encode(updateUser)
	} else {
		http.NotFound(w, r)
	}
}

// DeleteUser removes a user from the store
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/user/")
	mu.Lock()
	_, ok := users[id]
	if ok {
		delete(users, id)
	}
	mu.Unlock()
	if ok {
		w.WriteHeader(http.StatusOK)
	} else {
		http.NotFound(w, r)
	}
}
