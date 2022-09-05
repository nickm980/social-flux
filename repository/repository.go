package repository

import (
	"context"
	"log"
	utils "nickm980/utils"
	"os/user"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const connUri string = "mongodb+srv://nickm980:1234@cluster0.7h1oahg.mongodb.net/test"

var Client *mongo.Client = nil

func Connect() *mongo.Client {
	var result *mongo.Client
	if Client != nil {
		result = Client
	} else {
		result = connectRaw()
	}
	return result
}

func connectRaw() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(connUri))
	utils.CheckErr(err, "Error")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	utils.CheckErr(err, "an error has occured")

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	Client = client
	return client
}

func NewRepository(database string, collection string) *Repository {
	return &Repository{Database: "exproject", Collection: "users"}

}

type Repository struct {
	Database   string
	Collection string
}

func (r *Repository) FindByField(field string, value string) interface{} {
	var result user.User

	col := Client.Database(r.Database).Collection(r.Collection)
	col.FindOne(context.TODO(), bson.D{{Key: "name", Value: "myname"}}).Decode(&result)

	return result
}
