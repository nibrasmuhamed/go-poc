package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID    string `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}

type UserRepository interface {
	InsertUser(ctx context.Context, user User) (*mongo.InsertOneResult, error)
	FindUserByID(ctx context.Context, id string) (*User, error)
}

type MongoUserRepository struct {
	collection *mongo.Collection
}

func (r *MongoUserRepository) InsertUser(ctx context.Context, user User) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(ctx, user)
}

func (r *MongoUserRepository) FindUserByID(ctx context.Context, id string) (*User, error) {
	var user User
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	return &user, err
}
