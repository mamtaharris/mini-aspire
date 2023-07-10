package services

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mamtaharris/mini-aspire/config"
	"github.com/mamtaharris/mini-aspire/internal/models/entities"
	"github.com/mamtaharris/mini-aspire/internal/models/requests"
	"github.com/mamtaharris/mini-aspire/internal/repositories"
	rMock "github.com/mamtaharris/mini-aspire/internal/repositories/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_userService_ValidateUserAndGenerateToken(t *testing.T) {
	type fields struct {
		userRepo repositories.UserRepo
	}
	type args struct {
		ctx      context.Context
		loginReq requests.UserLoginReq
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		mockUserRepo func(ctrl *gomock.Controller) *rMock.MockUserRepo
		want         string
		wantErr      bool
	}{
		{
			name: "could not get user",
			mockUserRepo: func(ctrl *gomock.Controller) *rMock.MockUserRepo {
				userMock := rMock.NewMockUserRepo(ctrl)
				userMock.EXPECT().GetByUsername(gomock.Any(), gomock.Any()).Return(entities.Users{}, errors.New("dummy"))
				return userMock
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "incorrect creds",
			mockUserRepo: func(ctrl *gomock.Controller) *rMock.MockUserRepo {
				userMock := rMock.NewMockUserRepo(ctrl)
				userMock.EXPECT().GetByUsername(gomock.Any(), gomock.Any()).Return(entities.Users{Password: "aa"}, nil)
				return userMock
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "user not active",
			mockUserRepo: func(ctrl *gomock.Controller) *rMock.MockUserRepo {
				userMock := rMock.NewMockUserRepo(ctrl)
				userMock.EXPECT().GetByUsername(gomock.Any(), gomock.Any()).Return(entities.Users{}, nil)
				return userMock
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tt.fields.userRepo = tt.mockUserRepo(ctrl)

			u := NewUserService(
				tt.fields.userRepo,
			)
			got, err := u.ValidateUserAndGenerateToken(tt.args.ctx, tt.args.loginReq)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.ValidateUserAndGenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("userService.ValidateUserAndGenerateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_generateToken(t *testing.T) {
	config.InitConfig()
	token, err := generateToken(1)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}
