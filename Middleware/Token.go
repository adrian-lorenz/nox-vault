package Middleware

import (
	"github.com/adrian-lorenz/nox-vault/globals"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"slices"
)

type TokenData struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
	jwt.RegisteredClaims
}

func TokenRequiredLst(authlist []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if globals.JWTKey == "" || globals.MasterKey == "" {
			c.JSON(http.StatusUnauthorized, gin.H{})
			c.Abort()
			return
		}

		tokenString := c.Request.Header.Get("x-token")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{})
			c.Abort()
			return
		}
		var td TokenData
		_, err := jwt.ParseWithClaims(tokenString, &td, func(token *jwt.Token) (interface{}, error) {
			secretKey := []byte(globals.JWTKey)

			return secretKey, nil
		})

		if !slices.Contains(authlist, td.Name) {
			c.JSON(http.StatusUnauthorized, gin.H{})
			c.Abort()
			return
		}
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{})
			c.Abort()
			return
		}
		c.Set("tokenData", td)
		c.Next()
	}
}
