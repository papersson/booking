package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users", ListUsers)         // Changed from usersHandler
	http.HandleFunc("/users/create", CreateUser) // New route for creating users
	http.HandleFunc("/user/", GetUser)           // Might need logic to differentiate based on method or split further
	http.HandleFunc("/user/update/", UpdateUser) // New route for updating users
	http.HandleFunc("/user/delete/", DeleteUser) // New route for deleting users

	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
