package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mamtaharris/mini-aspire/config"
	"github.com/mamtaharris/mini-aspire/internal/repositories"
)

type AuthMiddleware struct {
	userRepo repositories.UserRepo
}

func NewAuthMiddleware(userRepo repositories.UserRepo) *AuthMiddleware {
	return &AuthMiddleware{userRepo: userRepo}
}

func (a *AuthMiddleware) Authenticate(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.App.JwtSecret), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	userID := claims["userID"].(float64)

	// Pass the authenticated username to the next handler
	c.Set("userID", int(userID))
	c.Next()
}

// Middleware to authorize the user based on role
func (a *AuthMiddleware) Authorize(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("userID")
		user, err := a.userRepo.GetByUserID(context.Background(), userID.(int))
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
