package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

// Create a struct to read the username and password from the request body
type Credentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Login(c *gin.Context) {
	var json Credentials
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
