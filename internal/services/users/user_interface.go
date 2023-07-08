package users

import (
	"context"

	"github.com/mamtaharris/mini-aspire/internal/models/requests"
	"github.com/mamtaharris/mini-aspire/internal/repositories/users"
)

type userService struct {
	userRepo users.UserRepo
}

func NewService(userRepo users.UserRepo) UserService {
	return &userService{userRepo: userRepo}
}

type UserService interface {
	ValidateUserAndGenerateToken(ctx context.Context, loginReq requests.UserLoginReq) (string, error)
}
