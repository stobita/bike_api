package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/stobita/bike_api/lib"
)

// TokenAuthMiddleware authentication middleware
func TokenAuthMiddleware(c *gin.Context) {
	authHeader := string(c.GetHeader("Authorization"))
	if authHeader == "" {
		c.JSON(400, lib.ErrorResponse("Auth header empty"))
		c.Abort()
		return
	}
	tokenString := strings.Split(authHeader, " ")[1]
	result, userID := lib.TokenAuthenticate(tokenString)
	if !result {
		c.JSON(401, lib.ErrorResponse("Invalid Token"))
		c.Abort()
		return
	}
	c.Set("userId", userID)
	c.Next()
}
