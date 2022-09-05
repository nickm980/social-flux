package users

import (
	"context"
	"log"
	"nickm980/repository"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	databaseName string = "exproject"
)

type User struct {
	Name     string `bson:"name" json:"name"`
	Password string `bson:"password" json:"-"`
	Email    string `bson:"email" json:"email"`
}

type serverError struct {
	StatusCode rune   `json:"statusCode"`
	Success    bool   `json:"success"`
	Message    string `json:"message"`
}

type successfulCreation struct {
	StatusCode rune   `json:"statusCode"`
	Success    bool   `json:"success"`
	Message    string `json:"message"`
}

func CreateUser(name string, email string, password string) (*User, interface{}) {
	col := repository.Client.Database(databaseName).Collection("users")

	user := User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	alreadyExistingUser, _ := FindUserByName(name)
	log.Default().Println(alreadyExistingUser)

	if alreadyExistingUser != nil {
		return nil, serverError{
			StatusCode: 400,
			Success:    false,
			Message:    "The user '" + name + "' already exists",
		}
	}

	_, err := col.InsertOne(context.TODO(), user)

	if err != nil {
		log.Default().Fatal("Database failed to insert a user")
		return nil, serverError{
			StatusCode: 500,
			Success:    false,
			Message:    "There was an internal error with the database query. Please report this issue.",
		}
	}

	return &user, &successfulCreation{
		StatusCode: 200,
		Success:    true,
		Message:    "The user has been created",
	}
}

func FindUserByName(name string) (*User, interface{}) {
	var result User
	col := repository.Client.Database(databaseName).Collection("users")
	col.FindOne(context.TODO(), bson.D{{Key: "name", Value: name}}).Decode(&result)

	if result.Name == "" {
		return nil, serverError{StatusCode: 400, Success: false, Message: "The user could not be found"}
	}

	return &result, nil
}
