package service 

import ("github.com/Algalyq/Go_boilerplate/pkg/repository")

type Authorization interface {

}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
    return &Service{}  
}