package handlers

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/mamtaharris/mini-aspire/internal/models/requests"
	"github.com/mamtaharris/mini-aspire/internal/models/responses"
	"github.com/mamtaharris/mini-aspire/internal/services"
	sMock "github.com/mamtaharris/mini-aspire/internal/services/mocks"
	"github.com/mamtaharris/mini-aspire/internal/validators"
	vMock "github.com/mamtaharris/mini-aspire/internal/validators/mocks"
)

func TestLoanHandler_CreateLoanHandler(t *testing.T) {
	type fields struct {
		loanSvc          services.LoanService
		loanReqValidator validators.LoanReqValidatorInterface
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name              string
		fields            fields
		mockLoanService   func(ctrl *gomock.Controller) *sMock.MockLoanService
		mockLoanValidator func(ctrl *gomock.Controller) *vMock.MockLoanReqValidatorInterface
		args              args
	}{
		{
			name: "failed at service layer",
			mockLoanService: func(ctrl *gomock.Controller) *sMock.MockLoanService {
				loanMock := sMock.NewMockLoanService(ctrl)
				loanMock.EXPECT().CreateLoan(gomock.Any(), gomock.Any()).Return(responses.LoanResp{}, errors.New("dummy"))
				return loanMock
			},
			mockLoanValidator: func(ctrl *gomock.Controller) *vMock.MockLoanReqValidatorInterface {
				loanMock := vMock.NewMockLoanReqValidatorInterface(ctrl)
				loanMock.EXPECT().ValidateCreateLoanReq(gomock.Any()).Return(requests.CreateLoanReq{}, nil)
				return loanMock
			},
		},
		{
			name: "failed at validator",
			mockLoanService: func(ctrl *gomock.Controller) *sMock.MockLoanService {
				loanMock := sMock.NewMockLoanService(ctrl)
				return loanMock
			},
			mockLoanValidator: func(ctrl *gomock.Controller) *vMock.MockLoanReqValidatorInterface {
				loanMock := vMock.NewMockLoanReqValidatorInterface(ctrl)
				loanMock.EXPECT().ValidateCreateLoanReq(gomock.Any()).Return(requests.CreateLoanReq{}, errors.New("dummy"))
				return loanMock
			},
		},
		{
			name: "happy case",
			mockLoanService: func(ctrl *gomock.Controller) *sMock.MockLoanService {
				loanMock := sMock.NewMockLoanService(ctrl)
				loanMock.EXPECT().CreateLoan(gomock.Any(), gomock.Any()).Return(responses.LoanResp{}, nil)
				return loanMock
			},
			mockLoanValidator: func(ctrl *gomock.Controller) *vMock.MockLoanReqValidatorInterface {
				loanMock := vMock.NewMockLoanReqValidatorInterface(ctrl)
				loanMock.EXPECT().ValidateCreateLoanReq(gomock.Any()).Return(requests.CreateLoanReq{}, nil)
				return loanMock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tt.fields.loanSvc = tt.mockLoanService(ctrl)
			tt.fields.loanReqValidator = tt.mockLoanValidator(ctrl)
			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			tt.args.ctx = ctx

			h := NewLoanHandler(
				tt.fields.loanSvc,
				tt.fields.loanReqValidator,
			)
			h.CreateLoanHandler(tt.args.ctx)
		})
	}
}

func TestLoanHandler_UpdateLoanHandler(t *testing.T) {
	type fields struct {
		loanSvc          services.LoanService
		loanReqValidator validators.LoanReqValidatorInterface
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		mockLoanService   func(ctrl *gomock.Controller) *sMock.MockLoanService
		mockLoanValidator func(ctrl *gomock.Controller) *vMock.MockLoanReqValidatorInterface
	}{
		{
			name: "failed at service layer",
			mockLoanService: func(ctrl *gomock.Controller) *sMock.MockLoanService {
				loanMock := sMock.NewMockLoanService(ctrl)
				loanMock.EXPECT().UpdateLoan(gomock.Any(), gomock.Any(), gomock.Any()).Return(responses.LoanResp{}, errors.New("dummy"))
				return loanMock
			},
			mockLoanValidator: func(ctrl *gomock.Controller) *vMock.MockLoanReqValidatorInterface {
				loanMock := vMock.NewMockLoanReqValidatorInterface(ctrl)
				loanMock.EXPECT().ValidateUpdateLoanReq(gomock.Any()).Return(requests.UpdateLoanReq{}, 0, nil)
				return loanMock
			},
		},
		{
			name: "failed at validator",
			mockLoanService: func(ctrl *gomock.Controller) *sMock.MockLoanService {
				loanMock := sMock.NewMockLoanService(ctrl)
				return loanMock
			},
			mockLoanValidator: func(ctrl *gomock.Controller) *vMock.MockLoanReqValidatorInterface {
				loanMock := vMock.NewMockLoanReqValidatorInterface(ctrl)
				loanMock.EXPECT().ValidateUpdateLoanReq(gomock.Any()).Return(requests.UpdateLoanReq{}, 0, errors.New("dummy"))
				return loanMock
			},
		},
		{
			name: "happy case",
			mockLoanService: func(ctrl *gomock.Controller) *sMock.MockLoanService {
				loanMock := sMock.NewMockLoanService(ctrl)
				loanMock.EXPECT().UpdateLoan(gomock.Any(), gomock.Any(), gomock.Any()).Return(responses.LoanResp{}, nil)
				return loanMock
			},
			mockLoanValidator: func(ctrl *gomock.Controller) *vMock.MockLoanReqValidatorInterface {
				loanMock := vMock.NewMockLoanReqValidatorInterface(ctrl)
				loanMock.EXPECT().ValidateUpdateLoanReq(gomock.Any()).Return(requests.UpdateLoanReq{}, 0, nil)
				return loanMock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tt.fields.loanSvc = tt.mockLoanService(ctrl)
			tt.fields.loanReqValidator = tt.mockLoanValidator(ctrl)
			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			tt.args.ctx = ctx

			h := NewLoanHandler(
				tt.fields.loanSvc,
				tt.fields.loanReqValidator,
			)
			h.UpdateLoanHandler(tt.args.ctx)
		})
	}
}

func TestLoanHandler_GetLoanHandler(t *testing.T) {
	type fields struct {
		loanSvc          services.LoanService
		loanReqValidator validators.LoanReqValidatorInterface
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		mockLoanService   func(ctrl *gomock.Controller) *sMock.MockLoanService
		mockLoanValidator func(ctrl *gomock.Controller) *vMock.MockLoanReqValidatorInterface
	}{
		{
			name: "failed at service layer",
			mockLoanService: func(ctrl *gomock.Controller) *sMock.MockLoanService {
				loanMock := sMock.NewMockLoanService(ctrl)
				loanMock.EXPECT().GetLoan(gomock.Any(), gomock.Any()).Return(responses.LoanResp{}, errors.New("dummy"))
				return loanMock
			},
			mockLoanValidator: func(ctrl *gomock.Controller) *vMock.MockLoanReqValidatorInterface {
				loanMock := vMock.NewMockLoanReqValidatorInterface(ctrl)
				loanMock.EXPECT().ValidateGetLoanReq(gomock.Any()).Return(0, nil)
				return loanMock
			},
		},
		{
			name: "failed at validator",
			mockLoanService: func(ctrl *gomock.Controller) *sMock.MockLoanService {
				loanMock := sMock.NewMockLoanService(ctrl)
				return loanMock
			},
			mockLoanValidator: func(ctrl *gomock.Controller) *vMock.MockLoanReqValidatorInterface {
				loanMock := vMock.NewMockLoanReqValidatorInterface(ctrl)
				loanMock.EXPECT().ValidateGetLoanReq(gomock.Any()).Return(0, errors.New("dummy"))
				return loanMock
			},
		},
		{
			name: "happy case",
			mockLoanService: func(ctrl *gomock.Controller) *sMock.MockLoanService {
				loanMock := sMock.NewMockLoanService(ctrl)
				loanMock.EXPECT().GetLoan(gomock.Any(), gomock.Any()).Return(responses.LoanResp{}, nil)
				return loanMock
			},
			mockLoanValidator: func(ctrl *gomock.Controller) *vMock.MockLoanReqValidatorInterface {
				loanMock := vMock.NewMockLoanReqValidatorInterface(ctrl)
				loanMock.EXPECT().ValidateGetLoanReq(gomock.Any()).Return(0, nil)
				return loanMock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tt.fields.loanSvc = tt.mockLoanService(ctrl)
			tt.fields.loanReqValidator = tt.mockLoanValidator(ctrl)
			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			tt.args.ctx = ctx

			h := NewLoanHandler(
				tt.fields.loanSvc,
				tt.fields.loanReqValidator,
			)
			h.GetLoanHandler(tt.args.ctx)
		})
	}
}

func TestLoanHandler_RepayLoanHandler(t *testing.T) {
	type fields struct {
		loanSvc          services.LoanService
		loanReqValidator validators.LoanReqValidatorInterface
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		mockLoanService   func(ctrl *gomock.Controller) *sMock.MockLoanService
		mockLoanValidator func(ctrl *gomock.Controller) *vMock.MockLoanReqValidatorInterface
	}{
		{
			name: "failed at service layer",
			mockLoanService: func(ctrl *gomock.Controller) *sMock.MockLoanService {
				loanMock := sMock.NewMockLoanService(ctrl)
				loanMock.EXPECT().RepayLoan(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(responses.LoanResp{}, errors.New("dummy"))
				return loanMock
			},
			mockLoanValidator: func(ctrl *gomock.Controller) *vMock.MockLoanReqValidatorInterface {
				loanMock := vMock.NewMockLoanReqValidatorInterface(ctrl)
				loanMock.EXPECT().ValidateRepayLoanReq(gomock.Any()).Return(requests.RepayLoanReq{}, 0, 0, nil)
				return loanMock
			},
		},
		{
			name: "failed at validator",
			mockLoanService: func(ctrl *gomock.Controller) *sMock.MockLoanService {
				loanMock := sMock.NewMockLoanService(ctrl)
				return loanMock
			},
			mockLoanValidator: func(ctrl *gomock.Controller) *vMock.MockLoanReqValidatorInterface {
				loanMock := vMock.NewMockLoanReqValidatorInterface(ctrl)
				loanMock.EXPECT().ValidateRepayLoanReq(gomock.Any()).Return(requests.RepayLoanReq{}, 0, 0, errors.New("dummy"))
				return loanMock
			},
		},
		{
			name: "happy case",
			mockLoanService: func(ctrl *gomock.Controller) *sMock.MockLoanService {
				loanMock := sMock.NewMockLoanService(ctrl)
				loanMock.EXPECT().RepayLoan(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(responses.LoanResp{}, nil)
				return loanMock
			},
			mockLoanValidator: func(ctrl *gomock.Controller) *vMock.MockLoanReqValidatorInterface {
				loanMock := vMock.NewMockLoanReqValidatorInterface(ctrl)
				loanMock.EXPECT().ValidateRepayLoanReq(gomock.Any()).Return(requests.RepayLoanReq{}, 0, 0, nil)
				return loanMock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tt.fields.loanSvc = tt.mockLoanService(ctrl)
			tt.fields.loanReqValidator = tt.mockLoanValidator(ctrl)
			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			tt.args.ctx = ctx

			h := NewLoanHandler(
				tt.fields.loanSvc,
				tt.fields.loanReqValidator,
			)
			h.RepayLoanHandler(tt.args.ctx)
		})
	}
}
