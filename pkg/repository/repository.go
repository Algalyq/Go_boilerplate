package repository

import (
    "go.mongodb.org/mongo-driver/mongo"
	"github.com/Algalyq/Go_boilerplate/data"
)


type Authorization interface {
	CreateUser(user data.User) (int, error)
	GetUser(username,password string) (data.User, error) 
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