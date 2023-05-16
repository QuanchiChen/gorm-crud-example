package controllers

import (
	"example/go-crud/initializers"
	"example/go-crud/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatePosts(c *gin.Context) {
	var body struct {
		Title string
		Body  string
	}
	err := c.Bind(&body)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func GetAllPosts(c *gin.Context) {
	var posts []models.Post
	result := initializers.DB.Find(&posts)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "posts not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func GetAPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	result := initializers.DB.First(&post, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "post not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func UpdateAPost(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Title string
		Body  string
	}
	err := c.Bind(&body)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	var post models.Post
	initializers.DB.First(&post, id)
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func DeleteAPost(c *gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(&models.Post{}, id)
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully deleted",
	})
}
