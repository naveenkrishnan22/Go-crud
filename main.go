package main

import (
	"example.com/go-crud/controller"
	"example.com/go-crud/initialize"
	"github.com/gin-gonic/gin"
)

func init() {
	initialize.LoadEnvVariable()
	initialize.ConnectToDB()
}
func main() {
	r := gin.Default()
	r.POST("/posts", controller.PostsCreate)
	r.GET("/posts", controller.PostIndex)
	r.GET("/posts/:id", controller.PostShow)
	r.PUT("/posts/:id", controller.PostUpdate)
	r.DELETE("/posts/:id", controller.PostDelete)

	r.Run() // listen and serve on 0.0.0.0:8080

}
