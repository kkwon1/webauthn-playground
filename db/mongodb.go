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


func Connect() {
	// TODO: Configure this elsewhere
	mongodb_uri := "mongodb://localhost:27017"

	client, err := mongo.NewClient(options.Client().ApplyURI(mongodb_uri))
	if err != nil {
			log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	fmt.Print("I have connected with the DB!")
	if err != nil {
			log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	
		/*
					List databases
	*/
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
			log.Fatal(err)
	}
	fmt.Println(databases)

}

func FindUser(username string) bool {
	log.Print("did not find: " + username)
	return false;
}

func SaveUser(user model.UserEntity) {
	log.Print("Saving user")
}