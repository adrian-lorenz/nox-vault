package Middleware

import (
	"github.com/adrian-lorenz/nox-vault/globals"
	"github.com/adrian-lorenz/nox-vault/security"
	"github.com/adrian-lorenz/nox-vault/tools"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
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
		if globals.JWTKey == "" || globals.MasterKey == "" && !globals.Look {
			c.JSON(http.StatusUnauthorized, gin.H{})
			log.Error("No JWTKey or MasterKey set")
			c.Abort()
			return
		}
		iip := tools.GetIP(c)
		if !security.CheckWhitelistsInt(iip) {
			c.JSON(http.StatusUnauthorized, gin.H{})
			log.Error("IP not whitelisted", iip)
			c.Abort()
			return
		}

		tokenString := c.Request.Header.Get("x-token")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{})
			log.Error("No token found")
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
			log.Error("No permission")
			c.Abort()
			return
		}
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{})
			log.Error("Token error", err)
			c.Abort()
			return
		}
		c.Set("tokenData", td)
		c.Next()
	}
}
