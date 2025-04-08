package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type User struct {
	Name  string `json:"nombre,omitempty"`
	Email string `json:"email"`
}

type Message struct {
	Text string `json:"message"`
}

func main() {
	router := chi.NewRouter()
	router.Get("/", messageFunc)
	router.Post("/users", handlerUser)

	fmt.Println("Server running on port 8000")
	http.ListenAndServe(":8000", router)
}

func messageFunc(w http.ResponseWriter, r *http.Request) {
	message := Message{"Hello World"}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func handlerUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("Error decoding JSON")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"message": "User created",
		"user":    user,
	})

}
