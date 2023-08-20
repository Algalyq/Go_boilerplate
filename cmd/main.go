package main 


import (
	"log"
	"github.com/Algalyq/Go_boilerplate"
	"github.com/Algalyq/Go_boilerplate/pkg/handler"
)

func main(){
	handler := new(handler.Handler)
	srv := new(Go_boilerplate.Server)
	if err := srv.Run("8080", handler.InitRoutes()); err != nil{
		log.Fatalf(err.Error())
	}
}