package main

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/google/uuid"
)

var (
	users = make(map[string]User) // In-memory store for users
	mu    sync.Mutex              // Mutex to avoid race condition
)

// usersHandler handles the creation of a new user and lists all users
func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
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
	case "GET":
		mu.Lock()
		values := make([]User, 0, len(users))
		for _, v := range users {
			values = append(values, v)
		}
		mu.Unlock()
		json.NewEncoder(w).Encode(values)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// userHandler handles the retrieval, update, and deletion of a user
func userHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/user/"):]

	mu.Lock()
	user, ok := users[id]
	mu.Unlock()

	if !ok {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(user)
	case "PUT":
		var updateUser User
		if err := json.NewDecoder(r.Body).Decode(&updateUser); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		updateUser.ID = user.ID // Ensure the ID remains unchanged
		mu.Lock()
		users[user.ID] = updateUser
		mu.Unlock()
		json.NewEncoder(w).Encode(updateUser)
	case "DELETE":
		mu.Lock()
		delete(users, user.ID)
		mu.Unlock()
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
