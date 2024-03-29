package Middleware

import (
	"github.com/adrian-lorenz/nox-vault/security"
	"github.com/adrian-lorenz/nox-vault/tools"
	"github.com/gin-gonic/gin"
)

func SysWhitelist() gin.HandlerFunc {
	return func(c *gin.Context) {
		acp := security.CheckWhitelistsInt(tools.GetIP(c))
		if !acp {
			c.JSON(401, gin.H{})
			c.Abort()
			return
		}
		c.Next()
	}
}
