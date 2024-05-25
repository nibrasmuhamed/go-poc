package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestCreateUser(t *testing.T) {
	mockRepo := &MockUserRepository{users: make(map[string]User)}
	userRepository = mockRepo

	r := chi.NewRouter()
	r.Post("/users", createUser)

	user := User{Name: "John Doe", Email: "john@example.com"}
	body, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var responseUser User
	if err := json.NewDecoder(rr.Body).Decode(&responseUser); err != nil {
		t.Errorf("handler returned unexpected body: %v", rr.Body.String())
	}

	if responseUser.Name != user.Name || responseUser.Email != user.Email {
		t.Errorf("handler returned wrong user: got %v want %v", responseUser, user)
	}
}

func TestGetUser(t *testing.T) {
	mockRepo := &MockUserRepository{users: make(map[string]User)}
	userRepository = mockRepo

	user := User{ID: "1", Name: "John Doe", Email: "john@example.com"}
	mockRepo.users[user.ID] = user

	r := chi.NewRouter()
	r.Get("/users/{id}", getUser)

	req, _ := http.NewRequest("GET", "/users/1", nil)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var responseUser User
	if err := json.NewDecoder(rr.Body).Decode(&responseUser); err != nil {
		t.Errorf("handler returned unexpected body: %v", rr.Body.String())
	}

	if responseUser.Name != user.Name || responseUser.Email != user.Email {
		t.Errorf("handler returned wrong user: got %v want %v", responseUser, user)
	}
}
