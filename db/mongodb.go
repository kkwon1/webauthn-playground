package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kkwon1/webauthn-playground-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDb struct {
	Client *mongo.Client
}

type UserDO struct {
	UserID string
	Username string
	PublicKeyCredentialCreationOptions model.PublicKeyCredentialCreationOptions
}

func Connect() MongoDb {
	// TODO: Configure this elsewhere
	mongodb_uri := "mongodb://localhost:27017"

	client, err := mongo.NewClient(options.Client().ApplyURI(mongodb_uri))
	if err != nil {
			log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("I have connected with the DB!")
	return MongoDb {
		Client: client,
	}
}

func FindUser(mongoDb MongoDb, username string) bool {
	users_collection := mongoDb.Client.Database("authdb").Collection("Users")

	var foundUser UserDO
	err := users_collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&foundUser)
	if (err != nil) {
		fmt.Println(err)
		return false
	}
	return true
}

func SaveUser(mongoDb MongoDb, userDO UserDO) error {
	users_collection := mongoDb.Client.Database("authdb").Collection("Users")

	insertResult, err := users_collection.InsertOne(context.Background(), userDO)
	if (err != nil) {
		fmt.Println(err)
		return err
	}

	fmt.Println(insertResult)
	return nil
}