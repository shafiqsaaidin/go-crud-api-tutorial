package main

import (
	"github.com/shafiqsaaidin/go-crud-api/initializers"
	"github.com/shafiqsaaidin/go-crud-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
