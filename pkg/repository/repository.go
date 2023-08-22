package repository

import (
    "go.mongodb.org/mongo-driver/mongo"
)


type Authorization interface {
}

type Searching interface {
}

type Repository struct {
	Authorization
	Searching
}

func NewRepository(db *mongo.Client) *Repository {
    return &Repository{
	}
}