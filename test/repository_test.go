package test

import (
	"nickm980/repository"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client = nil

func init() {
	client = repository.Connect()
}

// Creates, finds, then deletes an entry from the mongodb database
func TestFullRepository(t *testing.T) {

}

func TestRepositoryFindByField(t *testing.T) {
	var r repository.Repository = *repository.NewRepository("database", "collection")
	result := r.FindByField("name", "-")

	if result == nil {
		t.Errorf("Repository could not find the requested resource")
	}
}

func TestRepositoryShouldFailToFindById(t *testing.T) {
	var r repository.Repository = *repository.NewRepository("database", "collection")
	result := r.FindByField("name", "-")

	if result != nil {
		t.Errorf("Repository could not find the requested resource")
	}
}
