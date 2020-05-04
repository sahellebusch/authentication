package controllers

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"net/http"
	"time"
)

type Credentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type JWTResponse struct {
	JWT string
}

type LoginToken struct {
	jwt.Payload
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

var privateKey, _ = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
var publicKey = &privateKey.PublicKey
var hs = jwt.NewES256(
	jwt.ECDSAPublicKey(publicKey),
	jwt.ECDSAPrivateKey(privateKey),
)

func sign(id uint, username string) (string, error) {
	now := time.Now()
	pl := LoginToken{
		Payload: jwt.Payload{
			Issuer:         "coolcat",
			Subject:        "login",
			Audience:       jwt.Audience{},
			ExpirationTime: jwt.NumericDate(now.Add(7 * 24 * time.Hour)),
			NotBefore:      jwt.NumericDate(now.Add(30 * time.Minute)),
			IssuedAt:       jwt.NumericDate(now),
			JWTID:          uuid.NewV4().String(),
		},
		ID:       id,
		Username: username,
	}
	token, err := jwt.Sign(pl, hs)
	return string(token), err
}

func Login(c *gin.Context) {
	var creds Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := sign(uint(1), "boom")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	response := &JWTResponse{JWT: token}
	c.JSON(http.StatusOK, response)
}
