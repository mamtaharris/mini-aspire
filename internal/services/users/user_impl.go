package users

import (
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mamtaharris/mini-aspire/config"
)

func (u *userService) ValidateUserAndGenerateToken(ctx context.Context, username string) (string, error) {

	user, err := u.userRepository.GetByUsername(ctx, username)
	if err != nil {
		return "", err
	}
	if !user.IsActive {
		return "", errors.New("user is not active")
	}
	return "", nil
}

func generateToken(username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 2).Unix(), // Set expiration to 2 hours
	})

	tokenString, _ := token.SignedString([]byte(config.App.JwtSecret))
	return tokenString
}
