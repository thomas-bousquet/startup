package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	. "github.com/thomas-bousquet/user-service/errors"
	. "github.com/thomas-bousquet/user-service/models"
	uuid "github.com/thomas-bousquet/user-service/utils/id-generator"
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
		collection: mongoClient.Database("user-service").Collection("users"),
		logger:     logger,
	}
}

func (repo UserRepository) CreateUser(user User) (*string, error) {
	documentId := uuid.New().String()
	_, err := repo.collection.InsertOne(context.Background(),
		bson.M{
			"id":         documentId,
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"password":   user.Password,
			"email":      user.Email,
			"created_at": time.Now(),
			"role":       [1]string{"user"},
		})

	if err != nil {
		repo.logger.Errorf("%v", err)
		return nil, fmt.Errorf("error when creating new user with email %q: %v", user.Email, err)
	}

	return &documentId, nil
}

func (repo UserRepository) UpdateUser(id string, user User) error {
	_, err := repo.collection.UpdateOne(context.Background(), bson.M{"id": id}, bson.M{"$set": bson.M{
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
		"updated_at": time.Now(),
	},
	})

	if err != nil {
		repo.logger.Errorf("%v", err)
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
			repo.logger.Errorf("%v", result.Err())
			return nil, fmt.Errorf("error occured when finding user with email %q: %v", email, result.Err())
		}
	}

	err := result.Decode(&user)

	if err != nil {
		repo.logger.Errorf("%v", err)
		return nil, NewUnexpectedError(nil, nil)
	}

	return &user, nil
}

func (repo UserRepository) FindUserWithRole(id string, role string) (*User, error) {
	return repo.doFindUser(bson.M{"id": id, "role": role})
}

func (repo UserRepository) FindUser(id string) (*User, error) {
	return repo.doFindUser(bson.M{"id": id})
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
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return nil, nil
		} else {
			repo.logger.Errorf("%v", result.Err())
			return nil, NewUnexpectedError(nil, nil)
		}
	}

	err := result.Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
