package repository

import "go.mongodb.org/mongo-driver/mongo"

type UserRepository struct {
	mongoClient *mongo.Client
}

func NewUserRepository(mongoClient *mongo.Client) UserRepository {
	return UserRepository{
		mongoClient: mongoClient,
	}
}
