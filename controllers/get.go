package controllers

import (
	"authentication/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

// GET /books
// Get all books
func GetUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user []models.User
	db.Find(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// TODO figure out how to log and get a user
