package middlewares

import (
	"Gofinal1/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("my_secret_key")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Request does not contain an access token"})
			c.Abort()
			return
		}

		// Remove "Bearer " prefix from the token
		tokenString = tokenString[7:]

		claims := &controllers.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Set user ID in the context
		c.Set("user_id", claims.UserID) // เก็บค่า user_id ใน context
		c.Next()
	}
}
