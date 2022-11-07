package controller

import (
	"example.com/go-crud/initialize"
	"example.com/go-crud/models"
	"github.com/gin-gonic/gin"
)

// Create
func PostsCreate(c *gin.Context) {

	var body struct {
		Body string
		Name string
	}
	c.Bind(&body)
	post := models.Post{Name: body.Name, Body: body.Body}

	result := initialize.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}

// Read All
func PostIndex(c *gin.Context) {
	var posts []models.Post
	initialize.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

// Read Single
func PostShow(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initialize.DB.First(&post, id)

	c.JSON(200, gin.H{
		"posts": post,
	})
}

//Update

func PostUpdate(c *gin.Context) {

	id := c.Param("id")

	var body struct {
		Body string
		Name string
	}
	c.Bind(&body)

	var post models.Post
	initialize.DB.First(&post, id)

	initialize.DB.Model(&post).Updates(models.Post{Name: body.Name, Body: body.Body})
	c.JSON(200, gin.H{
		"posts": post,
	})

}

//Delete

func PostDelete(c *gin.Context) {

	id := c.Param("id")

	initialize.DB.Delete(&models.Post{}, id)
	c.Status(200)

}
