package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = make(map[string]*User)

var mutex sync.Mutex

// Main entry point to start the server and define routes
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", router)
}

// Handler to get all users
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Handler to get a user by ID
func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	mutex.Lock()
	defer mutex.Unlock()
	if user, ok := users[params["id"]]; ok {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
		return
	}
	http.NotFound(w, r)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	json.NewDecoder(r.Body).Decode(&newUser)
	mutex.Lock()
	users[newUser.ID] = &newUser
	mutex.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedUser User
	json.NewDecoder(r.Body).Decode(&updatedUser)
	mutex.Lock()
	defer mutex.Unlock()
	if user, ok := users[params["id"]]; ok {
		user.Name = updatedUser.Name
		user.Email = updatedUser.Email
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
		return
	}
	http.NotFound(w, r)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := users[params["id"]]; ok {
		delete(users, params["id"])
		w.WriteHeader(http.StatusNoContent)
		return
	}
	http.NotFound(w, r)
}
