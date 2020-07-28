package repository

import (
	"context"
	log "github.com/sirupsen/logrus"
	. "github.com/thomas-bousquet/startup/model"
	"go.mongodb.org/mongo-driver/bson"
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

func (repo UserRepository) FindUserByEmail(email string) User {
	user := User{}
	err := repo.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)

	if err != nil {
		log.Error(err)
	}

	return user
}

func (repo UserRepository) FindUser(id string) User {
	user := User{}
	err := repo.collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&user)

	if err != nil {
		log.Error(err)
	}

	return user
}
