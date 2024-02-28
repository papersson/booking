package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/user/", userHandler)

	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
