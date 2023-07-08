package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mamtaharris/mini-aspire/internal/models/responses"
	"github.com/mamtaharris/mini-aspire/internal/services/users"
	"github.com/mamtaharris/mini-aspire/internal/validators/user"
)

type UserHandler struct {
	userService   users.UserService
	userValidator user.UserReqValidatorInterface
}

func NewUserHandler(userService users.UserService, userValidator user.UserReqValidatorInterface) *UserHandler {
	return &UserHandler{userService: userService, userValidator: userValidator}
}

func (h *UserHandler) Login(ctx *gin.Context) {
	req, err := h.userValidator.ValidateUserLoginReq(ctx)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, responses.ErrorResp{Error: err.Error()})
		return
	}
	response, err := h.userService.ValidateUserAndGenerateToken(ctx, req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResp{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, responses.Token{Token: response})
}
