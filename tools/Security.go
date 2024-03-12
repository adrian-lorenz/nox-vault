package tools

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	token string = "e40966ce65ab4662b7b132d3af6a2482"
)

func TokenRequired() gin.HandlerFunc {

	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("x-token")
		// Token fehlt
		if tokenString == "" {
			Respond(c, http.StatusUnauthorized, gin.H{"error": "A token is required"})
			c.Abort()
			return
		}

		if tokenString != token {
			log.Error("Couldn't handle this token")
			Respond(c, http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		log.Infoln("token ok")
		c.Next()
	}
}
