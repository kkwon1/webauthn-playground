package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kkwon1/webauthn-playground-service/cmd"
	"github.com/kkwon1/webauthn-playground-service/db"
	"github.com/segmentio/ksuid"
)

type RegisterUserRequest struct {
	Username string `validate:"required,alphanumeric"`
}

func RegisterUser(c *gin.Context) {
	mongoDb := db.Connect()

	var req RegisterUserRequest
	c.BindJSON(&req)
	username := req.Username

	err := validateUsername(username)
	if (err != nil) {
		c.String(http.StatusBadRequest, "Invalid username. Please user alphanumeric values only")
		return
	}

	userExists := db.FindUser(mongoDb, username)

	if userExists {
		c.String(http.StatusBadRequest, "username already exists")
	} else {
		userId := ksuid.New().String()
		options := cmd.CreatePublicKeyCredentialCreationOption(userId, username)

		userDO := db.UserDO {
			UserID: userId,
			Username: username,
			PublicKeyCredentialCreationOptions: options,
		}

		db.SaveUser(mongoDb, userDO)
		c.JSON(http.StatusOK, options)
	}
}

func validateUsername(username string) error {
	if (username =="") {
		fmt.Println("Invalid username dumb ass")
		fmt.Println(username)
		return errors.New("Invalid username")
	}

	return nil
}
