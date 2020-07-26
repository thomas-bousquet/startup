package repository

import "go.mongodb.org/mongo-driver/mongo"

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(userCollection *mongo.Collection) UserRepository {
	return UserRepository{
		collection: userCollection,
	}
}
