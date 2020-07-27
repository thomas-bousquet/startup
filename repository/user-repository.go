package repository

import (
	"context"
	log "github.com/sirupsen/logrus"
	. "github.com/thomas-bousquet/startup/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(userCollection *mongo.Collection) UserRepository {
	return UserRepository{
		collection: userCollection,
	}
}

func (repo UserRepository) CreateUser(user User) {
	_, err := repo.collection.InsertOne(context.Background(), user)

	if err != nil {
		log.Error(err)
		log.Errorf("Error when creating user with email '%s'", user.Email)
	}
}
