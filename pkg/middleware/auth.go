package middleware

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mamtaharris/mini-aspire/config"
	"github.com/mamtaharris/mini-aspire/internal/models/entities"
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

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.App.JwtSecret), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	username := claims["userId"].(string)

	// Pass the authenticated username to the next handler
	c.Set("userId", username)
	c.Next()
}

// Middleware to authorize the user based on role
func (a *AuthMiddleware) Authorize(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		username, _ := c.Get("username")
		user, err := a.getUserByUsername(username.(string))
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

func (a *AuthMiddleware) getUserByUsername(username string) (entities.Users, error) {
	user, err := a.userRepo.GetByUsername(context.Background(), username)
	if err != nil {
		return entities.Users{}, err
	}
	return user, nil
}
