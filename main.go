package main

import (
	"example/go-crud/controllers"
	"example/go-crud/initializers"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/posts", controllers.GetAllPosts)
	r.GET("/posts/:id", controllers.GetAPost)
	r.POST("/posts", controllers.CreatePosts)
	r.PUT("/posts/:id", controllers.UpdateAPost)
	r.DELETE("/posts/:id", controllers.DeleteAPost)

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
