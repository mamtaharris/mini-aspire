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

	return generateToken(user.ID)
}

func generateToken(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": userId,
		"exp":      time.Now().Add(time.Hour * 2).Unix(), // Set expiration to 2 hours
	})

	tokenString, err := token.SignedString([]byte(config.App.JwtSecret))
	if err != nil {
		return "", errors.New("failed to generate token")
	}
	return tokenString, nil
}
