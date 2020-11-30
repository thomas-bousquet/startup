package repositories

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	. "github.com/thomas-bousquet/startup/errors"
	. "github.com/thomas-bousquet/startup/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserRepository struct {
	collection *mongo.Collection
	logger     *logrus.Logger
}

func NewUserRepository(mongoClient *mongo.Client, logger *logrus.Logger) UserRepository {
	return UserRepository{
		collection: mongoClient.Database("startup").Collection("users"),
		logger:     logger,
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
			"role":       [1]string{"user"},
		})

	if err != nil {
		repo.logger.Error(err)
		return primitive.ObjectID{}, fmt.Errorf("error when creating new user with email %q: %v", user.Email, err)
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
		repo.logger.Error(err)
		return fmt.Errorf("error when updating user with email %q: %v", user.Email, err)
	}

	return nil
}

func (repo UserRepository) FindUserByEmail(email string) (*User, error) {
	var user User
	result := repo.collection.FindOne(context.Background(), bson.M{"email": email})

	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, nil
		} else {
			return nil, NewUnexpectedError()
		}
	}

	err := result.Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo UserRepository) FindUserWithRole(id string, role string) (*User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	return repo.doFindUser(bson.M{"_id": objectId, "role": role})
}

func (repo UserRepository) FindUser(id string) (*User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	return repo.doFindUser(bson.M{"_id": objectId})
}

func (repo UserRepository) FindUsers() ([]User, error) {

	results, err := repo.collection.Find(context.Background(), bson.M{})

	if err != nil {
		return nil, err
	}

	var users []User
	for results.Next(context.Background()) {
		var user User
		err := results.Decode(&user)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repo UserRepository) doFindUser(query primitive.M) (*User, error) {
	var user User

	result := repo.collection.FindOne(context.Background(), query)

	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, nil
		} else {
			return nil, NewUnexpectedError()
		}
	}

	err := result.Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo UserRepository) AuthenticateUser(email string, password string) (*User, error) {
	var user User
	result := repo.collection.FindOne(context.Background(), bson.M{"email": email, "password": password})

	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, nil
		} else {
			return nil, NewUnexpectedError()
		}
	}

	err := result.Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
