package users

import (
	"github.com/gin-gonic/gin"
	"github.com/mamtaharris/mini-aspire/internal/repositories/users"
)

type UserHandler struct {
	userRepository users.UserRepo
}

func NewHandler(userRepo users.UserRepo) *UserHandler {
	return &UserHandler{userRepository: userRepo}
}

func (h *UserHandler) Login(ctx *gin.Context) {

}
