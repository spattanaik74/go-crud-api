package main

import (
	"fmt"
	"log"

	"crud-api/initializers"
	"crud-api/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal(" could not load environmnet variables", err)
	}

	initializers.ConnectDB(&config)

	
}


func main() {

	initializers.DB.AutoMigrate(&models.User{}, &models.Post{})
	fmt.Println("? Migration Complete")
}