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
	"github.com/mamtaharris/mini-aspire/internal/constants"
	"github.com/mamtaharris/mini-aspire/internal/models/requests"
	vMock "github.com/mamtaharris/mini-aspire/internal/validators/mocks"
)

func Test_loanReqValidator_ValidateCreateLoanReq(t *testing.T) {
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
		want          requests.CreateLoanReq
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
			want:    requests.CreateLoanReq{},
			wantErr: true,
		},
		{
			name: "invalid request",
			args: args{ctx: &gin.Context{Request: &http.Request{Body: io.NopCloser(bytes.NewBuffer([]byte(`{
				"amount": 100
			}`)))}},
			},
			mockValidator: func(ctrl *gomock.Controller) *vMock.MockValidatorInterface {
				repaymentMock := vMock.NewMockValidatorInterface(ctrl)
				repaymentMock.EXPECT().ValidateUnknownParams(gomock.Any(), gomock.Any()).Return(nil)
				return repaymentMock
			},
			want:    requests.CreateLoanReq{},
			wantErr: true,
		},
		{
			name: "happy case",
			args: args{ctx: &gin.Context{Request: &http.Request{Body: io.NopCloser(bytes.NewBuffer([]byte(`{
				"amount": 100,
				"term": 1
			}`)))}},
			},
			mockValidator: func(ctrl *gomock.Controller) *vMock.MockValidatorInterface {
				repaymentMock := vMock.NewMockValidatorInterface(ctrl)
				repaymentMock.EXPECT().ValidateUnknownParams(gomock.Any(), gomock.Any()).Return(nil)
				return repaymentMock
			},
			want:    requests.CreateLoanReq{Amount: 100, Term: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tt.fields.validator = tt.mockValidator(ctrl)

			v := NewLoanValidator(
				tt.fields.validator,
			)
			got, err := v.ValidateCreateLoanReq(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("loanReqValidator.ValidateCreateLoanReq() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loanReqValidator.ValidateCreateLoanReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loanReqValidator_ValidateUpdateLoanReq(t *testing.T) {
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
		want          requests.UpdateLoanReq
		want1         int
		wantErr       bool
	}{
		{
			name: "unknown params present",
			args: args{ctx: &gin.Context{
				Request: &http.Request{Body: io.NopCloser(bytes.NewBuffer([]byte(`{
				"test": "test"
			}`)))},
				Params: gin.Params{gin.Param{Key: "id", Value: "123"}}},
			},
			mockValidator: func(ctrl *gomock.Controller) *vMock.MockValidatorInterface {
				repaymentMock := vMock.NewMockValidatorInterface(ctrl)
				repaymentMock.EXPECT().ValidateUnknownParams(gomock.Any(), gomock.Any()).Return(errors.New("dummy"))
				return repaymentMock
			},
			want:    requests.UpdateLoanReq{},
			wantErr: true,
		},
		{
			name: "invalid request - not a valid status",
			args: args{ctx: &gin.Context{
				Request: &http.Request{Body: io.NopCloser(bytes.NewBuffer([]byte(`{
					"status": "PENDING"
			}`)))},
				Params: gin.Params{gin.Param{Key: "id", Value: "123"}}},
			},
			mockValidator: func(ctrl *gomock.Controller) *vMock.MockValidatorInterface {
				repaymentMock := vMock.NewMockValidatorInterface(ctrl)
				repaymentMock.EXPECT().ValidateUnknownParams(gomock.Any(), gomock.Any()).Return(nil)
				return repaymentMock
			},
			want:    requests.UpdateLoanReq{},
			wantErr: true,
		},
		{
			name: "invalid request - not a number",
			args: args{ctx: &gin.Context{
				Request: &http.Request{Body: io.NopCloser(bytes.NewBuffer([]byte(`{
					"status": "APPROVED"
			}`)))},
				Params: gin.Params{gin.Param{Key: "id", Value: "abc"}}},
			},
			mockValidator: func(ctrl *gomock.Controller) *vMock.MockValidatorInterface {
				repaymentMock := vMock.NewMockValidatorInterface(ctrl)
				repaymentMock.EXPECT().ValidateUnknownParams(gomock.Any(), gomock.Any()).Return(nil)
				return repaymentMock
			},
			want:    requests.UpdateLoanReq{},
			wantErr: true,
		},
		{
			name: "invalid request - no body",
			args: args{ctx: &gin.Context{
				Request: &http.Request{Body: io.NopCloser(bytes.NewBuffer([]byte(`{
			}`)))},
				Params: gin.Params{gin.Param{Key: "id", Value: "123"}}},
			},
			mockValidator: func(ctrl *gomock.Controller) *vMock.MockValidatorInterface {
				repaymentMock := vMock.NewMockValidatorInterface(ctrl)
				repaymentMock.EXPECT().ValidateUnknownParams(gomock.Any(), gomock.Any()).Return(nil)
				return repaymentMock
			},
			want:    requests.UpdateLoanReq{},
			wantErr: true,
		},
		{
			name: "happy case",
			args: args{ctx: &gin.Context{
				Request: &http.Request{Body: io.NopCloser(bytes.NewBuffer([]byte(`{
				"status": "APPROVED"
			}`)))},
				Params: gin.Params{gin.Param{Key: "id", Value: "123"}}},
			},
			mockValidator: func(ctrl *gomock.Controller) *vMock.MockValidatorInterface {
				repaymentMock := vMock.NewMockValidatorInterface(ctrl)
				repaymentMock.EXPECT().ValidateUnknownParams(gomock.Any(), gomock.Any()).Return(nil)
				return repaymentMock
			},
			want:    requests.UpdateLoanReq{Status: constants.LoanStatus.Approved},
			want1:   123,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tt.fields.validator = tt.mockValidator(ctrl)

			v := NewLoanValidator(
				tt.fields.validator,
			)
			got, got1, err := v.ValidateUpdateLoanReq(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("loanReqValidator.ValidateUpdateLoanReq() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loanReqValidator.ValidateUpdateLoanReq() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("loanReqValidator.ValidateUpdateLoanReq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_loanReqValidator_ValidateGetLoanReq(t *testing.T) {
	type fields struct {
		validator ValidatorInterface
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "not a valid loan id",
			args: args{ctx: &gin.Context{
				Params: gin.Params{gin.Param{Key: "id", Value: "abc"}}},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "happy case",
			args: args{ctx: &gin.Context{
				Params: gin.Params{gin.Param{Key: "id", Value: "123"}}},
			},
			want:    123,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewLoanValidator(
				tt.fields.validator,
			)
			got, err := v.ValidateGetLoanReq(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("loanReqValidator.ValidateGetLoanReq() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("loanReqValidator.ValidateGetLoanReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loanReqValidator_ValidateRepayLoanReq(t *testing.T) {
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
		want          requests.RepayLoanReq
		want1         int
		want2         int
		wantErr       bool
	}{
		{
			name: "incorrect repayment id",
			args: args{ctx: &gin.Context{
				Request: &http.Request{Body: io.NopCloser(bytes.NewBuffer([]byte(`{
					"amount": 100
				}`)))},
				Params: gin.Params{gin.Param{Key: "loanID", Value: "123"}, gin.Param{Key: "repaymentID", Value: "abc"}}},
			},
			mockValidator: func(ctrl *gomock.Controller) *vMock.MockValidatorInterface {
				repaymentMock := vMock.NewMockValidatorInterface(ctrl)
				repaymentMock.EXPECT().ValidateUnknownParams(gomock.Any(), gomock.Any()).Return(nil)
				return repaymentMock
			},
			want:    requests.RepayLoanReq{},
			want1:   0,
			want2:   0,
			wantErr: true,
		},
		{
			name: "incorrect loan id",
			args: args{ctx: &gin.Context{
				Request: &http.Request{Body: io.NopCloser(bytes.NewBuffer([]byte(`{
					"amount": 100
				}`)))},
				Params: gin.Params{gin.Param{Key: "loanID", Value: "abc"}, gin.Param{Key: "repaymentID", Value: "123"}}},
			},
			mockValidator: func(ctrl *gomock.Controller) *vMock.MockValidatorInterface {
				repaymentMock := vMock.NewMockValidatorInterface(ctrl)
				repaymentMock.EXPECT().ValidateUnknownParams(gomock.Any(), gomock.Any()).Return(nil)
				return repaymentMock
			},
			want:    requests.RepayLoanReq{},
			want1:   0,
			want2:   0,
			wantErr: true,
		},
		{
			name: "incorrect request body",
			args: args{ctx: &gin.Context{
				Request: &http.Request{Body: io.NopCloser(bytes.NewBuffer([]byte(`{
			}`)))},
				Params: gin.Params{gin.Param{Key: "loanID", Value: "123"}, gin.Param{Key: "repaymentID", Value: "123"}}},
			},
			mockValidator: func(ctrl *gomock.Controller) *vMock.MockValidatorInterface {
				repaymentMock := vMock.NewMockValidatorInterface(ctrl)
				repaymentMock.EXPECT().ValidateUnknownParams(gomock.Any(), gomock.Any()).Return(nil)
				return repaymentMock
			},
			want:    requests.RepayLoanReq{},
			want1:   0,
			want2:   0,
			wantErr: true,
		},
		{
			name: "received unknown params ",
			args: args{ctx: &gin.Context{
				Request: &http.Request{Body: io.NopCloser(bytes.NewBuffer([]byte(`{
				"test": 100
			}`)))},
				Params: gin.Params{gin.Param{Key: "loanID", Value: "123"}, gin.Param{Key: "repaymentID", Value: "123"}}},
			},
			mockValidator: func(ctrl *gomock.Controller) *vMock.MockValidatorInterface {
				repaymentMock := vMock.NewMockValidatorInterface(ctrl)
				repaymentMock.EXPECT().ValidateUnknownParams(gomock.Any(), gomock.Any()).Return(errors.New("dummy"))
				return repaymentMock
			},
			want:    requests.RepayLoanReq{},
			want1:   0,
			want2:   0,
			wantErr: true,
		},
		{
			name: "happy case",
			args: args{ctx: &gin.Context{
				Request: &http.Request{Body: io.NopCloser(bytes.NewBuffer([]byte(`{
				"amount": 100
			}`)))},
				Params: gin.Params{gin.Param{Key: "loanID", Value: "123"}, gin.Param{Key: "repaymentID", Value: "123"}}},
			},
			mockValidator: func(ctrl *gomock.Controller) *vMock.MockValidatorInterface {
				repaymentMock := vMock.NewMockValidatorInterface(ctrl)
				repaymentMock.EXPECT().ValidateUnknownParams(gomock.Any(), gomock.Any()).Return(nil)
				return repaymentMock
			},
			want:    requests.RepayLoanReq{Amount: 100},
			want1:   123,
			want2:   123,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tt.fields.validator = tt.mockValidator(ctrl)

			v := NewLoanValidator(
				tt.fields.validator,
			)
			got, got1, got2, err := v.ValidateRepayLoanReq(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("loanReqValidator.ValidateRepayLoanReq() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loanReqValidator.ValidateRepayLoanReq() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("loanReqValidator.ValidateRepayLoanReq() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("loanReqValidator.ValidateRepayLoanReq() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
