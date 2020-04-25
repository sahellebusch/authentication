package main

import (
	"authentication/controllers"
	"authentication/models"
	"github.com/gin-gonic/gin"
)

func main() {

	db := models.SetupModels()
	defer db.Close()
	router := gin.New() // if you use Default and add middleware, it'll print twice
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	router.GET("/status", controllers.AreWeGood)

	v1 := router.Group("/v1")
	{
		v1.GET("/user", controllers.GetUsers)
	}

	router.Run() // listen and serve on 0.0.0.0:8080
}
