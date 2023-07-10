package validators

import (
	"github.com/gin-gonic/gin"
	"github.com/mamtaharris/mini-aspire/internal/models/requests"
)

type userReqValidator struct{}

func NewUserValidator() UserReqValidatorInterface {
	return &userReqValidator{}
}

//go:generate mockgen -package mocks -source=loan_interface.go -destination=mocks/user_interface_mocks.go
type UserReqValidatorInterface interface {
	ValidateUserLoginReq(ctx *gin.Context) (requests.UserLoginReq, error)
}

func (v *userReqValidator) ValidateUserLoginReq(ctx *gin.Context) (requests.UserLoginReq, error) {
	var reqBody requests.UserLoginReq
	err := ValidateUnknownParams(&reqBody, ctx)
	if err != nil {
		return reqBody, err
	}
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		return reqBody, err
	}
	return reqBody, nil
}
