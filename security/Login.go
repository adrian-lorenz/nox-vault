package security

import (
	"github.com/adrian-lorenz/nox-vault/database"
	"github.com/adrian-lorenz/nox-vault/globals"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	type Login struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// check username and password
	var user database.User
	result := database.DB.Where(database.User{Username: json.Username}).First(&user)
	if result.Error != nil {
		c.JSON(401, gin.H{"message": "Username or password wrong"})
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(json.Password)) != nil {
		c.JSON(401, gin.H{"message": "Username or password wrong"})
		return
	}
	// create jwt token
	claims := jwt.MapClaims{
		"name": user.Username,
		"uuid": user.UUID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(globals.JWTKey))
	if err != nil {
		c.JSON(500, gin.H{"message": "Error creating token"})
		return
	}

	c.JSON(200, gin.H{"token": tokenString})
}
