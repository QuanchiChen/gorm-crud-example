package main

import (
	"example/go-crud/initializers"
	"example/go-crud/models"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatal(err)
	}
}
