package services

import (
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mamtaharris/mini-aspire/config"
	"github.com/mamtaharris/mini-aspire/internal/models/requests"
	"github.com/mamtaharris/mini-aspire/internal/repositories"
	"gorm.io/gorm"
)

type userService struct {
	userRepo repositories.UserRepo
}

func NewUserService(userRepo repositories.UserRepo) UserService {
	return &userService{userRepo: userRepo}
}

//go:generate mockgen -package mocks -source=user.go -destination=mocks/user_mocks.go

type UserService interface {
	ValidateUserAndGenerateToken(ctx context.Context, loginReq requests.UserLoginReq) (string, error)
}

func (u *userService) ValidateUserAndGenerateToken(ctx context.Context, loginReq requests.UserLoginReq) (string, error) {
	user, err := u.userRepo.GetByUsername(ctx, loginReq.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("credentials are incorrect")
		}
		return "", err
	}
	if user.Password != loginReq.Password {
		return "", errors.New("credentials are incorrect")
	}
	if !user.IsActive {
		return "", errors.New("user is not active")
	}
	return generateToken(user.ID)
}

func generateToken(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 2).Unix(), // Set expiration to 2 hours
	})

	tokenString, err := token.SignedString([]byte(config.App.JwtSecret))
	if err != nil {
		return "", errors.New("failed to generate token")
	}
	return tokenString, nil
}
