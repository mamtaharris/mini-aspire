package validators

import (
	"github.com/gin-gonic/gin"
	"github.com/mamtaharris/mini-aspire/internal/models/requests"
)

type userReqValidator struct {
	validator ValidatorInterface
}

func NewUserValidator(validator ValidatorInterface) UserReqValidatorInterface {
	return &userReqValidator{
		validator: validator,
	}
}

//go:generate mockgen -package mocks -source=user.go -destination=mocks/user_mocks.go
type UserReqValidatorInterface interface {
	ValidateUserLoginReq(ctx *gin.Context) (requests.UserLoginReq, error)
}

func (v *userReqValidator) ValidateUserLoginReq(ctx *gin.Context) (requests.UserLoginReq, error) {
	var reqBody requests.UserLoginReq
	err := v.validator.ValidateUnknownParams(&reqBody, ctx)
	if err != nil {
		return requests.UserLoginReq{}, err
	}
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		return requests.UserLoginReq{}, err
	}
	return reqBody, nil
}
