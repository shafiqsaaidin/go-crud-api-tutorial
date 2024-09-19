package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shafiqsaaidin/go-crud-api/controllers"
	"github.com/shafiqsaaidin/go-crud-api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.CreatePost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetSinglePost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	r.Run() // listen and serve on 0.0.0.0:8080
}
