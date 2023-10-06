package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"crud-api/controllers"
	"crud-api/initializers"
	"crud-api/routes"
)

var (
	server *gin.Engine

	PostController      controllers.PostController
	PostRouteController routes.PostRouteController
)



func init() {
	config, err := initializers.LoadConfig(".")

	if err != nil {
		log.Fatal("? Coul not load env variables", err)
	}

	initializers.ConnectDB(&config)

	PostController = controllers.NewPostController(initializers.DB)
	PostRouteController = routes.NewRoutePostController(PostController)

	server = gin.Default()
}


func main() {
	config, err := initializers.LoadConfig(".")

	if err != nil {
		log.Fatal("? could not load env variables", err)
	}

	router := server.Group("/api")

	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang"

		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	PostRouteController.PostRoute(router)

	log.Fatal(server.Run(":" + config.ServerPort))

}