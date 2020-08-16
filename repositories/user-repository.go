package repositories

import (
	"context"
	log "github.com/sirupsen/logrus"
	. "github.com/thomas-bousquet/startup/models"
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

func (repo UserRepository) CreateUser(user User) (primitive.ObjectID, error) {
	result, err := repo.collection.InsertOne(context.Background(),
		bson.M{
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"password":   user.Password,
			"email":      user.Email,
			"created_at": time.Now(),
		})

	if err != nil {
		log.Errorf("Error when creating user with email '%s'", user.Email)
		return primitive.ObjectID{}, err
	}

	id := result.InsertedID.(primitive.ObjectID)

	return id, nil
}

func (repo UserRepository) UpdateUser(id string, user User) error {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = repo.collection.UpdateOne(context.Background(), bson.M{"_id": objectId}, bson.M{"$set":
	bson.M{
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
		"updated_at": time.Now(),
	},
	})

	if err != nil {
		log.Errorf("Error when updating user with email '%s'", user.Email)
		return err
	}

	return nil
}

func (repo UserRepository) FindUserByEmail(email string) (*User, error) {
	user := User{}
	err := repo.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo UserRepository) FindUser(id string) (*User, error) {
	user := User{}
	objectId, err := primitive.ObjectIDFromHex(id)
	err = repo.collection.FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
