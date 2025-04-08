package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name  string `json:"nombre,omitempty"`
	Email string `json:"email"`
}

type Message struct {
	Text string `json:"message"`
}

func main() {
	// user := User{"Sol", "test@gmail.com"}

	// data, err := json.Marshal(user)
	// if err != nil {
	// 	fmt.Errorf("Error: %v", err)
	// 	return

	// }
	// fmt.Println(string(data))

	// jsonString := `{"name":"Jose"}`
	// var userUnmarshal User

	// err = json.Unmarshal([]byte(jsonString), &userUnmarshal)
	// if err != nil {
	// 	fmt.Errorf("Error: %v", err)
	// 	return
	// }
	// fmt.Printf("%v", userUnmarshal)
	http.HandleFunc("/", messageFunc)
	http.HandleFunc("/users", handlerUser)
	http.ListenAndServe(":8080", nil)

}

// Encoder
func messageFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		message := Message{"Hello World"}
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func handlerUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

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
