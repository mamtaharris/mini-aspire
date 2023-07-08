package user

import (
	"github.com/gin-gonic/gin"
	"github.com/mamtaharris/mini-aspire/internal/models/requests"
)

type userReqValidator struct{}

func NewValidator() UserReqValidatorInterface {
	return &userReqValidator{}
}

//go:generate mockgen -package mocks -source=loan_interface.go -destination=mocks/loan_interface_mocks.go
type UserReqValidatorInterface interface {
	ValidateUserLoginReq(ctx *gin.Context) (requests.UserLoginReq, error)
}
