package clients

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"time"
)

func NewMongoClient(logger *logrus.Logger) *mongo.Client {
	mongoUrl := os.Getenv("MONGODB_URL") + ":" + os.Getenv("MONGODB_PORT")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUrl).SetAuth(options.Credential{
		Username:                os.Getenv("MONGODB_USERNAME"),
		Password:                os.Getenv("MONGODB_PASSWORD"),
	}))

	if err != nil {
		logger.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = client.Connect(ctx); err != nil {
		logger.Fatal(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		logger.Fatal(err)
	}

	return client
}