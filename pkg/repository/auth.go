package repository

import (
	"context"
	"fmt"
	"github.com/Algalyq/Go_boilerplate/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthMongoDB struct {
	collection *mongo.Collection
}

func NewAuthMongoDB(client *mongo.Client, dbName, collectionName string) *AuthMongoDB {
	db := client.Database(dbName)
	collection := db.Collection(collectionName)
	return &AuthMongoDB{collection: collection}
}

func (a *AuthMongoDB) CreateUser(user data.User) (string, error) {
	

	// Insert the user document into the MongoDB collection
	insertResult, err := a.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return "", err
	}

	// Get the ID of the inserted document
	insertedID := insertResult.InsertedID
	return fmt.Sprintf("%v", insertedID), nil
}

func (a *AuthMongoDB) GetUser(username, password string) (data.User, error) {
	var user data.User

	// Define a filter to find the user by username and password
	filter := bson.M{"username": username, "password": password}

	// Find the user document in the MongoDB collection
	err := a.collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return data.User{}, err
	}

	return user, nil
}
