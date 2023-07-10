package validators

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/mamtaharris/mini-aspire/internal/models/requests"
	vMock "github.com/mamtaharris/mini-aspire/internal/validators/mocks"
)

func Test_userReqValidator_ValidateUserLoginReq(t *testing.T) {
	type fields struct {
		validator ValidatorInterface
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		mockValidator func(ctrl *gomock.Controller) *vMock.MockValidatorInterface
		want          requests.UserLoginReq
		wantErr       bool
	}{
		{
			name: "unknown params present",
			args: args{ctx: &gin.Context{Request: &http.Request{Body: io.NopCloser(bytes.NewBuffer([]byte(`{
				"test": "test"
			}`)))}},
			},
			mockValidator: func(ctrl *gomock.Controller) *vMock.MockValidatorInterface {
				repaymentMock := vMock.NewMockValidatorInterface(ctrl)
				repaymentMock.EXPECT().ValidateUnknownParams(gomock.Any(), gomock.Any()).Return(errors.New("dummy"))
				return repaymentMock
			},
			want:    requests.UserLoginReq{},
			wantErr: true,
		},
		{
			name: "invalid request",
			args: args{ctx: &gin.Context{Request: &http.Request{Body: io.NopCloser(bytes.NewBuffer([]byte(`{
				"username": "test"
			}`)))}},
			},
			mockValidator: func(ctrl *gomock.Controller) *vMock.MockValidatorInterface {
				repaymentMock := vMock.NewMockValidatorInterface(ctrl)
				repaymentMock.EXPECT().ValidateUnknownParams(gomock.Any(), gomock.Any()).Return(nil)
				return repaymentMock
			},
			want:    requests.UserLoginReq{},
			wantErr: true,
		},
		{
			name: "happy case",
			args: args{ctx: &gin.Context{Request: &http.Request{Body: io.NopCloser(bytes.NewBuffer([]byte(`{
				"username": "test",
				"password": "test"
			}`)))}},
			},
			mockValidator: func(ctrl *gomock.Controller) *vMock.MockValidatorInterface {
				repaymentMock := vMock.NewMockValidatorInterface(ctrl)
				repaymentMock.EXPECT().ValidateUnknownParams(gomock.Any(), gomock.Any()).Return(nil)
				return repaymentMock
			},
			want:    requests.UserLoginReq{Username: "test", Password: "test"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tt.fields.validator = tt.mockValidator(ctrl)

			v := NewUserValidator(
				tt.fields.validator,
			)
			got, err := v.ValidateUserLoginReq(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("userReqValidator.ValidateUserLoginReq() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userReqValidator.ValidateUserLoginReq() = %v, want %v", got, tt.want)
			}
		})
	}
}
