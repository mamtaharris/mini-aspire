package handlers

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/mamtaharris/mini-aspire/internal/models/requests"
	"github.com/mamtaharris/mini-aspire/internal/services"
	sMock "github.com/mamtaharris/mini-aspire/internal/services/mocks"
	"github.com/mamtaharris/mini-aspire/internal/validators"
	vMock "github.com/mamtaharris/mini-aspire/internal/validators/mocks"
)

func TestUserHandler_Login(t *testing.T) {
	type fields struct {
		userService   services.UserService
		userValidator validators.UserReqValidatorInterface
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		mockUserService   func(ctrl *gomock.Controller) *sMock.MockUserService
		mockUserValidator func(ctrl *gomock.Controller) *vMock.MockUserReqValidatorInterface
	}{
		{
			name: "failed at service layer",
			mockUserService: func(ctrl *gomock.Controller) *sMock.MockUserService {
				userMock := sMock.NewMockUserService(ctrl)
				userMock.EXPECT().ValidateUserAndGenerateToken(gomock.Any(), gomock.Any()).Return("", errors.New("dummy"))
				return userMock
			},
			mockUserValidator: func(ctrl *gomock.Controller) *vMock.MockUserReqValidatorInterface {
				loanMock := vMock.NewMockUserReqValidatorInterface(ctrl)
				loanMock.EXPECT().ValidateUserLoginReq(gomock.Any()).Return(requests.UserLoginReq{}, nil)
				return loanMock
			},
		},
		{
			name: "failed at validator",
			mockUserService: func(ctrl *gomock.Controller) *sMock.MockUserService {
				userMock := sMock.NewMockUserService(ctrl)
				return userMock
			},
			mockUserValidator: func(ctrl *gomock.Controller) *vMock.MockUserReqValidatorInterface {
				loanMock := vMock.NewMockUserReqValidatorInterface(ctrl)
				loanMock.EXPECT().ValidateUserLoginReq(gomock.Any()).Return(requests.UserLoginReq{}, errors.New("dummy"))
				return loanMock
			},
		},
		{
			name: "happy case",
			mockUserService: func(ctrl *gomock.Controller) *sMock.MockUserService {
				userMock := sMock.NewMockUserService(ctrl)
				userMock.EXPECT().ValidateUserAndGenerateToken(gomock.Any(), gomock.Any()).Return("", nil)
				return userMock
			},
			mockUserValidator: func(ctrl *gomock.Controller) *vMock.MockUserReqValidatorInterface {
				loanMock := vMock.NewMockUserReqValidatorInterface(ctrl)
				loanMock.EXPECT().ValidateUserLoginReq(gomock.Any()).Return(requests.UserLoginReq{}, nil)
				return loanMock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tt.fields.userService = tt.mockUserService(ctrl)
			tt.fields.userValidator = tt.mockUserValidator(ctrl)
			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			tt.args.ctx = ctx

			h := NewUserHandler(
				tt.fields.userService,
				tt.fields.userValidator,
			)
			h.Login(tt.args.ctx)
		})
	}
}
