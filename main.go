package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kkwon1/webauthn-playground-service/db"
	"github.com/kkwon1/webauthn-playground-service/model"
	"github.com/kkwon1/webauthn-playground-service/protocol"
	"github.com/segmentio/ksuid"
)

func main() {
	fmt.Print("Hello World!")

	router := gin.Default()
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
		Name: "Kevin Test",
		ID: "123",
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
	}

	return options;
}