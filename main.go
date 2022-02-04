package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kkwon1/webauthn-playground-service/db"
	"github.com/kkwon1/webauthn-playground-service/model"
	"github.com/kkwon1/webauthn-playground-service/protocol"
	"github.com/segmentio/ksuid"
)

func main() {
	fmt.Print("Hello World!")

	router := gin.Default()
	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = []string{"*"}
	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true

	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS", "GET", "POST", "PUT")

	// Register the middleware
	router.Use(cors.New(corsConfig))

	router.GET("/test", test)
	router.POST("/register", registerUser)

	router.Run("localhost:8080")
}

func test(c *gin.Context) {
	db.Connect()

	c.JSON(200, gin.H{
		"message": "test",
	})
}

type RegisterUserRequest struct {
	Username string
}

func registerUser(c *gin.Context) {
	var req RegisterUserRequest
	c.BindJSON(&req)
	username := req.Username
	
	// TODO: perform username validation. Invalid characters, trim, etc.
	if (username == "") {
		c.String(http.StatusBadRequest, "username was not provided")
	}

	userExists := db.FindUser(username)

	if (userExists) {
		c.String(http.StatusBadRequest, "username already exists")
	} else {
		user := model.UserEntity {
			Name: "Kevin",
			DisplayName: "Kevin",
			ID: ksuid.New().String(),
		}

		db.SaveUser(user)
		options := createPublicKeyCredentialCreationOption("test", "test")
		c.JSON(http.StatusOK, options)
	}
}



func createPublicKeyCredentialCreationOption(id string, username string) model.PublicKeyCredentialCreationOptions {
	rp := model.RelyingPartyEntity {
		Name: "localhost",
		ID: "localhost",
	}

	userEntity := model.UserEntity {
		Name: "Kevin",
		DisplayName: "Kevin",
		ID: ksuid.New().String(),
	}

	options := model.PublicKeyCredentialCreationOptions {
		Challenge: protocol.CreateChallenge(),
		RelyingParty: rp,
		User: userEntity,
		Parameters: defaultRegistrationCredentialParameters(),
	}

	return options;
}



func defaultRegistrationCredentialParameters() []protocol.CredentialParameter {
	return []protocol.CredentialParameter{
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: protocol.AlgES256,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: protocol.AlgES384,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: protocol.AlgES512,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: protocol.AlgRS256,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: protocol.AlgRS384,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: protocol.AlgRS512,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: protocol.AlgPS256,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: protocol.AlgPS384,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: protocol.AlgPS512,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: protocol.AlgEdDSA,
		},
	}
}