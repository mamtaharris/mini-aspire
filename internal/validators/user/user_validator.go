package user

import (
	"github.com/gin-gonic/gin"
	"github.com/mamtaharris/mini-aspire/internal/models/requests"
	"github.com/mamtaharris/mini-aspire/internal/validators"
)

func (v *userReqValidator) ValidateUserLoginReq(ctx *gin.Context) (requests.UserLoginReq, error) {
	var reqBody requests.UserLoginReq
	err := validators.ValidateUnknownParams(&reqBody, ctx)
	if err != nil {
		return reqBody, err
	}
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		return reqBody, err
	}
	return reqBody, nil
}
