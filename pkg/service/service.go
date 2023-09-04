package service 

import (
	"github.com/Algalyq/Go_boilerplate/pkg/repository"
	"github.com/Algalyq/Go_boilerplate/data"
)

type Authorization interface {

	CreateUser(user data.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	RefreshToken(username,password string) (string,error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
    return &Service{}  
}