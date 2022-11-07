package main

import (
	"example.com/go-crud/initialize"
	"example.com/go-crud/models"
)

func init() {
	initialize.LoadEnvVariable()
	initialize.ConnectToDB()
}

func main() {
	initialize.DB.AutoMigrate(&models.Post{})
}
