package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mamtaharris/mini-aspire/config"
	"github.com/mamtaharris/mini-aspire/internal/models/entities"
)

func authenticate(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.App.JwtSecret), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	// Pass the authenticated username to the next handler
	c.Set("username", username)
	c.Next()
}

// Middleware to authorize the user based on role
func authorize(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		username, _ := c.Get("username")
		user, err := getUserByUsername(username.(string))
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			c.Abort()
			return
		}
		if user.Role != role {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func getUserByUsername(username string) (entities.Users, error) {

	return entities.Users{}, nil
}
