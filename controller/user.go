package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kkwon1/webauthn-playground-service/cmd"
	"github.com/kkwon1/webauthn-playground-service/db"
	"github.com/kkwon1/webauthn-playground-service/model"
	"github.com/segmentio/ksuid"
)

type RegisterUserRequest struct {
	Username string
}

func RegisterUser(c *gin.Context) {
	var req RegisterUserRequest
	c.BindJSON(&req)
	username := req.Username

	// TODO: perform username validation. Invalid characters, trim, etc.
	if username == "" {
		c.String(http.StatusBadRequest, "username was not provided")
	}

	userExists := db.FindUser(username)

	if userExists {
		c.String(http.StatusBadRequest, "username already exists")
	} else {
		user := model.UserEntity{
			Name:        "Kevin",
			DisplayName: "Kevin",
			ID:          ksuid.New().String(),
		}

		db.SaveUser(user)
		options := cmd.CreatePublicKeyCredentialCreationOption("test", "test")
		c.JSON(http.StatusOK, options)
	}
}
