package controllers

import "github.com/gin-gonic/gin"

func AreWeGood(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "we good",
	})
}
