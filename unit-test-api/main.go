package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

var userRepository UserRepository

func main() {
	// MongoDB connection (skipped for brevity)

	// Setup repository (replace this with actual MongoDB setup)
	// userRepository = &MongoUserRepository{collection: mongoCollection}

	// Setup router
	r := chi.NewRouter()

	r.Post("/users", createUser)
	r.Get("/users/{id}", getUser)
	// r.Get("/users", getAllUsers)
	http.ListenAndServe(":3000", r)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := userRepository.InsertUser(ctx, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.ID = res.InsertedID.(string)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := userRepository.FindUserByID(ctx, id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
