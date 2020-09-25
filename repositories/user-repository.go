package repositories

import (
	"context"
	log "github.com/sirupsen/logrus"
	. "startup/errors"
	. "startup/models"
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
			"role": [1]string{"user"},
		})

	if err != nil {
		// @TODO: Wrap error so we keep the root cause
		log.Errorf("Error when creating user with email '%s'", user.Email)
		log.Error(err)
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
