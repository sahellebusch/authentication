package controllers

import (
	"authentication/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user []models.User
	db.Find(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}
