package repository

import (
	"context"
	log "github.com/sirupsen/logrus"
	. "github.com/thomas-bousquet/startup/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(userCollection *mongo.Collection) UserRepository {
	return UserRepository{
		collection: userCollection,
	}
}

func (repo UserRepository) CreateUser(user User) primitive.ObjectID {
	result, err := repo.collection.InsertOne(context.Background(),
		bson.M{
			"first_name": user.Firstname,
			"last_name":  user.Lastname,
			"password": user.Password,
			"email":      user.Email,
			"created_at": time.Now(),
		})

	if err != nil {
		log.Error(err)
		log.Errorf("Error when creating user with email '%s'", user.Email)
	}

	return result.InsertedID.(primitive.ObjectID)
}

func (repo UserRepository) UpdateUser(id string, user User) {
	objectId, _ := primitive.ObjectIDFromHex(id)
	_, err := repo.collection.UpdateOne(context.Background(), bson.M{"_id": objectId}, bson.M{"$set":
	bson.M{
		"first_name": user.Firstname,
		"last_name":  user.Lastname,
		"email":      user.Email,
		"updated_at": time.Now(),
	},
	})

	if err != nil {
		log.Error(err)
		log.Errorf("Error when updating user with email '%s'", user.Email)
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
	objectId, _ := primitive.ObjectIDFromHex(id)
	err := repo.collection.FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&user)

	if err != nil {
		log.Error(err)
	}

	return user
}
