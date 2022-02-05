package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kkwon1/webauthn-playground-service/controller"
)

func main() {
	router := setupRouter()

	router.Run("localhost:8080")
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowMethods("OPTIONS")
	router.Use(cors.New(corsConfig))

	return router
}

func setupEndpoints(router *gin.Engine) {
	router.POST("/register", controller.RegisterUser)
}
