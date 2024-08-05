package middlewares

import (
	"net/http"

	"example.com/rest/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	// validates token is part of request
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		// stops the current request
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}
	// verifies token
	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
		return
	}

	context.Set("userId", userId)
	// allows next function to execute
	context.Next()
}
