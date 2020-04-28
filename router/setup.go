package router

import (
	"authentication/controllers"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.New() // if you use Default and add middleware, it'll print twice

	if os.Getenv("GO_ENV") != "test" {
		router.Use(gin.Logger())
		router.Use(gin.Recovery())
	}

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	router.GET("/status", controllers.AreWeGood)

	v1 := router.Group("/v1")
	{
		v1.GET("/user", controllers.GetUsers)
	}

	return router
}
