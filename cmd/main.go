package main 


import (
	"log"
	"github.com/spf13/viper"
	"github.com/Algalyq/Go_boilerplate/data"
	"github.com/Algalyq/Go_boilerplate/pkg/handler"
	"github.com/Algalyq/Go_boilerplate/pkg/repository"
	"github.com/Algalyq/Go_boilerplate/pkg/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal(err.Error())
	}

	// Connect to MongoDB Atlas
	clientOptions := options.Client().ApplyURI("mongodb+srv://algalyq:2003720an@cluster0.bosxgra.mongodb.net/?retryWrites=true&w=majority")
	mongoClient, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB Atlas: %v", err)
	}
	defer mongoClient.Disconnect(context.Background())

	repo := repository.NewRepository(mongoClient)
	services := service.NewService(repo)
	handler := handler.NewHandler(services)

	srv := new(data.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf(err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
