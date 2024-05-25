package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type MockUserRepository struct {
	users map[string]User
}

func (m *MockUserRepository) InsertUser(ctx context.Context, user User) (*mongo.InsertOneResult, error) {
	m.users[user.ID] = user
	return &mongo.InsertOneResult{InsertedID: user.ID}, nil
}

func (m *MockUserRepository) FindUserByID(ctx context.Context, id string) (*User, error) {
	user, exists := m.users[id]
	if !exists {
		return nil, mongo.ErrNoDocuments
	}
	return &user, nil
}
