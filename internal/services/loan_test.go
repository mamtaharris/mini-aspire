package services

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/mamtaharris/mini-aspire/internal/constants"
	"github.com/mamtaharris/mini-aspire/internal/models/entities"
	"github.com/mamtaharris/mini-aspire/internal/models/requests"
	"github.com/mamtaharris/mini-aspire/internal/models/responses"
	"github.com/mamtaharris/mini-aspire/internal/repositories"
	rMock "github.com/mamtaharris/mini-aspire/internal/repositories/mocks"
)

func Test_loanService_CreateLoan(t *testing.T) {
	type fields struct {
		loanRepo      repositories.LoanRepo
		repaymentRepo repositories.RepaymentRepo
	}
	type args struct {
		ctx context.Context
		req requests.CreateLoanReq
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		mockLoanRepo      func(ctrl *gomock.Controller) *rMock.MockLoanRepo
		mockRepaymentRepo func(ctrl *gomock.Controller) *rMock.MockRepaymentRepo
		sleep             time.Duration
		want              responses.LoanResp
		wantErr           bool
	}{
		{
			name: "could not get list of repayment schedule for loan",
			mockLoanRepo: func(ctrl *gomock.Controller) *rMock.MockLoanRepo {
				loanMock := rMock.NewMockLoanRepo(ctrl)
				loanMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(entities.Loans{Amount: 100, Term: 1}, nil)
				return loanMock
			},
			mockRepaymentRepo: func(ctrl *gomock.Controller) *rMock.MockRepaymentRepo {
				repaymentMock := rMock.NewMockRepaymentRepo(ctrl)
				repaymentMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(entities.Repayments{}, nil)
				repaymentMock.EXPECT().GetAllRepaymentsForLoanID(gomock.Any(), gomock.Any()).Return([]entities.Repayments{}, errors.New("dummy"))
				return repaymentMock
			},
			want:    responses.LoanResp{},
			wantErr: true,
		},
		{
			name: "repayment entry not created",
			mockLoanRepo: func(ctrl *gomock.Controller) *rMock.MockLoanRepo {
				loanMock := rMock.NewMockLoanRepo(ctrl)
				loanMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(entities.Loans{Amount: 100, Term: 1}, nil)
				return loanMock
			},
			mockRepaymentRepo: func(ctrl *gomock.Controller) *rMock.MockRepaymentRepo {
				repaymentMock := rMock.NewMockRepaymentRepo(ctrl)
				repaymentMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(entities.Repayments{}, errors.New("dummy"))
				return repaymentMock
			},
			want:    responses.LoanResp{},
			wantErr: true,
		},
		{
			name: "loan entry not cretaed",
			mockLoanRepo: func(ctrl *gomock.Controller) *rMock.MockLoanRepo {
				loanMock := rMock.NewMockLoanRepo(ctrl)
				loanMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(entities.Loans{Amount: 100, Term: 1}, errors.New("dummy"))
				return loanMock
			},
			mockRepaymentRepo: func(ctrl *gomock.Controller) *rMock.MockRepaymentRepo {
				repaymentMock := rMock.NewMockRepaymentRepo(ctrl)
				return repaymentMock
			},
			want:    responses.LoanResp{},
			wantErr: true,
		},
		{
			name: "happy case",
			mockLoanRepo: func(ctrl *gomock.Controller) *rMock.MockLoanRepo {
				loanMock := rMock.NewMockLoanRepo(ctrl)
				loanMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(entities.Loans{Amount: 100, Term: 1}, nil)
				return loanMock
			},
			mockRepaymentRepo: func(ctrl *gomock.Controller) *rMock.MockRepaymentRepo {
				repaymentMock := rMock.NewMockRepaymentRepo(ctrl)
				repaymentMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(entities.Repayments{}, nil)
				repaymentMock.EXPECT().GetAllRepaymentsForLoanID(gomock.Any(), gomock.Any()).Return([]entities.Repayments{{Amount: 100}}, nil)
				return repaymentMock
			},
			want:    responses.LoanResp{Amount: 100, Term: 1, Repayments: []responses.Repayments{{Amount: 100}}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tt.fields.loanRepo = tt.mockLoanRepo(ctrl)
			tt.fields.repaymentRepo = tt.mockRepaymentRepo(ctrl)

			l := NewLoanService(
				tt.fields.loanRepo,
				tt.fields.repaymentRepo,
			)
			got, err := l.CreateLoan(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("loanService.CreateLoan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loanService.CreateLoan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loanService_UpdateLoan(t *testing.T) {
	type fields struct {
		loanRepo      repositories.LoanRepo
		repaymentRepo repositories.RepaymentRepo
	}
	type args struct {
		ctx    context.Context
		req    requests.UpdateLoanReq
		loanID int
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		mockLoanRepo      func(ctrl *gomock.Controller) *rMock.MockLoanRepo
		mockRepaymentRepo func(ctrl *gomock.Controller) *rMock.MockRepaymentRepo
		want              responses.LoanResp
		wantErr           bool
	}{
		{
			name: "could not get repayment schedule for loan",
			mockLoanRepo: func(ctrl *gomock.Controller) *rMock.MockLoanRepo {
				loanMock := rMock.NewMockLoanRepo(ctrl)
				loanMock.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(entities.Loans{Amount: 100, Term: 1}, nil)
				loanMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(entities.Loans{Amount: 100, Term: 1}, nil)
				return loanMock
			},
			mockRepaymentRepo: func(ctrl *gomock.Controller) *rMock.MockRepaymentRepo {
				repaymentMock := rMock.NewMockRepaymentRepo(ctrl)
				repaymentMock.EXPECT().GetAllRepaymentsForLoanID(gomock.Any(), gomock.Any()).Return([]entities.Repayments{{Amount: 100}}, errors.New("dummy"))
				return repaymentMock
			},
			want:    responses.LoanResp{},
			wantErr: true,
		},
		{
			name: "could not update loan",
			mockLoanRepo: func(ctrl *gomock.Controller) *rMock.MockLoanRepo {
				loanMock := rMock.NewMockLoanRepo(ctrl)
				loanMock.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(entities.Loans{Amount: 100, Term: 1}, nil)
				loanMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(entities.Loans{Amount: 100, Term: 1}, errors.New("dummy"))
				return loanMock
			},
			mockRepaymentRepo: func(ctrl *gomock.Controller) *rMock.MockRepaymentRepo {
				repaymentMock := rMock.NewMockRepaymentRepo(ctrl)
				return repaymentMock
			},
			want:    responses.LoanResp{},
			wantErr: true,
		},
		{
			name: "could not get loan",
			mockLoanRepo: func(ctrl *gomock.Controller) *rMock.MockLoanRepo {
				loanMock := rMock.NewMockLoanRepo(ctrl)
				loanMock.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(entities.Loans{Amount: 100, Term: 1}, errors.New("dummy "))
				return loanMock
			},
			mockRepaymentRepo: func(ctrl *gomock.Controller) *rMock.MockRepaymentRepo {
				repaymentMock := rMock.NewMockRepaymentRepo(ctrl)
				return repaymentMock
			},
			want:    responses.LoanResp{},
			wantErr: true,
		},
		{
			name: "happy case",
			mockLoanRepo: func(ctrl *gomock.Controller) *rMock.MockLoanRepo {
				loanMock := rMock.NewMockLoanRepo(ctrl)
				loanMock.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(entities.Loans{Amount: 100, Term: 1}, nil)
				loanMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(entities.Loans{Amount: 100, Term: 1}, nil)
				return loanMock
			},
			mockRepaymentRepo: func(ctrl *gomock.Controller) *rMock.MockRepaymentRepo {
				repaymentMock := rMock.NewMockRepaymentRepo(ctrl)
				repaymentMock.EXPECT().GetAllRepaymentsForLoanID(gomock.Any(), gomock.Any()).Return([]entities.Repayments{{Amount: 100}}, nil)
				return repaymentMock
			},
			want:    responses.LoanResp{Amount: 100, Term: 1, Repayments: []responses.Repayments{{Amount: 100}}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tt.fields.loanRepo = tt.mockLoanRepo(ctrl)
			tt.fields.repaymentRepo = tt.mockRepaymentRepo(ctrl)

			l := NewLoanService(
				tt.fields.loanRepo,
				tt.fields.repaymentRepo,
			)
			got, err := l.UpdateLoan(tt.args.ctx, tt.args.req, tt.args.loanID)
			if (err != nil) != tt.wantErr {
				t.Errorf("loanService.UpdateLoan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loanService.UpdateLoan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loanService_GetLoan(t *testing.T) {
	type fields struct {
		loanRepo      repositories.LoanRepo
		repaymentRepo repositories.RepaymentRepo
	}
	type args struct {
		ctx    context.Context
		loanID int
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		mockLoanRepo      func(ctrl *gomock.Controller) *rMock.MockLoanRepo
		mockRepaymentRepo func(ctrl *gomock.Controller) *rMock.MockRepaymentRepo
		want              responses.LoanResp
		wantErr           bool
	}{
		{
			name: "could not get repayment schedule for loan",
			mockLoanRepo: func(ctrl *gomock.Controller) *rMock.MockLoanRepo {
				loanMock := rMock.NewMockLoanRepo(ctrl)
				loanMock.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(entities.Loans{Amount: 100, Term: 1}, nil)
				return loanMock
			},
			mockRepaymentRepo: func(ctrl *gomock.Controller) *rMock.MockRepaymentRepo {
				repaymentMock := rMock.NewMockRepaymentRepo(ctrl)
				repaymentMock.EXPECT().GetAllRepaymentsForLoanID(gomock.Any(), gomock.Any()).Return([]entities.Repayments{{Amount: 100}}, errors.New("dummy"))
				return repaymentMock
			},
			want:    responses.LoanResp{},
			wantErr: true,
		},
		{
			name: "could not get loan",
			mockLoanRepo: func(ctrl *gomock.Controller) *rMock.MockLoanRepo {
				loanMock := rMock.NewMockLoanRepo(ctrl)
				loanMock.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(entities.Loans{Amount: 100, Term: 1}, errors.New("dummy"))
				return loanMock
			},
			mockRepaymentRepo: func(ctrl *gomock.Controller) *rMock.MockRepaymentRepo {
				repaymentMock := rMock.NewMockRepaymentRepo(ctrl)
				return repaymentMock
			},
			want:    responses.LoanResp{},
			wantErr: true,
		},
		{
			name: "happy case",
			mockLoanRepo: func(ctrl *gomock.Controller) *rMock.MockLoanRepo {
				loanMock := rMock.NewMockLoanRepo(ctrl)
				loanMock.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(entities.Loans{Amount: 100, Term: 1}, nil)
				return loanMock
			},
			mockRepaymentRepo: func(ctrl *gomock.Controller) *rMock.MockRepaymentRepo {
				repaymentMock := rMock.NewMockRepaymentRepo(ctrl)
				repaymentMock.EXPECT().GetAllRepaymentsForLoanID(gomock.Any(), gomock.Any()).Return([]entities.Repayments{{Amount: 100}}, nil)
				return repaymentMock
			},
			want:    responses.LoanResp{Amount: 100, Term: 1, Repayments: []responses.Repayments{{Amount: 100}}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tt.fields.loanRepo = tt.mockLoanRepo(ctrl)
			tt.fields.repaymentRepo = tt.mockRepaymentRepo(ctrl)

			l := NewLoanService(
				tt.fields.loanRepo,
				tt.fields.repaymentRepo,
			)
			got, err := l.GetLoan(tt.args.ctx, tt.args.loanID)
			if (err != nil) != tt.wantErr {
				t.Errorf("loanService.GetLoan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loanService.GetLoan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loanService_RepayLoan(t *testing.T) {
	type fields struct {
		loanRepo      repositories.LoanRepo
		repaymentRepo repositories.RepaymentRepo
	}
	type args struct {
		ctx         context.Context
		req         requests.RepayLoanReq
		loanID      int
		repaymentID int
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		mockLoanRepo      func(ctrl *gomock.Controller) *rMock.MockLoanRepo
		mockRepaymentRepo func(ctrl *gomock.Controller) *rMock.MockRepaymentRepo
		want              responses.LoanResp
		wantErr           bool
	}{
		{
			name: "not all repayments of loan are complete",
			mockLoanRepo: func(ctrl *gomock.Controller) *rMock.MockLoanRepo {
				loanMock := rMock.NewMockLoanRepo(ctrl)
				loanMock.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(entities.Loans{Amount: 100, Term: 2, Status: constants.LoanStatus.Approved}, nil)
				loanMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(entities.Loans{Amount: 100, Term: 1}, nil)
				return loanMock
			},
			mockRepaymentRepo: func(ctrl *gomock.Controller) *rMock.MockRepaymentRepo {
				repaymentMock := rMock.NewMockRepaymentRepo(ctrl)
				repaymentMock.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(entities.Repayments{Status: constants.RepaymentStatus.Pending}, nil).AnyTimes()
				repaymentMock.EXPECT().GetAllRepaymentsForLoanID(gomock.Any(), gomock.Any()).Return([]entities.Repayments{{Amount: 100}}, nil).AnyTimes()
				repaymentMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(entities.Repayments{Amount: 100}, nil)
				return repaymentMock
			},
			want:    responses.LoanResp{Amount: 100, Term: 1, Repayments: []responses.Repayments{{Amount: 100}}},
			wantErr: false,
		},
		{
			name: "could not get all repayment",
			mockLoanRepo: func(ctrl *gomock.Controller) *rMock.MockLoanRepo {
				loanMock := rMock.NewMockLoanRepo(ctrl)
				loanMock.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(entities.Loans{Amount: 100, Term: 1, Status: constants.LoanStatus.Approved}, nil)
				return loanMock
			},
			mockRepaymentRepo: func(ctrl *gomock.Controller) *rMock.MockRepaymentRepo {
				repaymentMock := rMock.NewMockRepaymentRepo(ctrl)
				repaymentMock.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(entities.Repayments{Status: constants.RepaymentStatus.Pending}, nil)
				repaymentMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(entities.Repayments{Amount: 100}, nil)
				repaymentMock.EXPECT().GetAllRepaymentsForLoanID(gomock.Any(), gomock.Any()).Return([]entities.Repayments{{Amount: 100}}, errors.New("dummy"))
				return repaymentMock
			},
			want:    responses.LoanResp{},
			wantErr: true,
		},
		{
			name: "could not update repayment",
			mockLoanRepo: func(ctrl *gomock.Controller) *rMock.MockLoanRepo {
				loanMock := rMock.NewMockLoanRepo(ctrl)
				loanMock.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(entities.Loans{Amount: 100, Term: 1, Status: constants.LoanStatus.Approved}, nil)
				return loanMock
			},
			mockRepaymentRepo: func(ctrl *gomock.Controller) *rMock.MockRepaymentRepo {
				repaymentMock := rMock.NewMockRepaymentRepo(ctrl)
				repaymentMock.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(entities.Repayments{Status: constants.RepaymentStatus.Pending}, nil)
				repaymentMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(entities.Repayments{Amount: 100}, errors.New("dummy"))
				return repaymentMock
			},
			want:    responses.LoanResp{},
			wantErr: true,
		},
		{
			name: "could not get repayment",
			mockLoanRepo: func(ctrl *gomock.Controller) *rMock.MockLoanRepo {
				loanMock := rMock.NewMockLoanRepo(ctrl)
				loanMock.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(entities.Loans{Amount: 100, Term: 1, Status: constants.LoanStatus.Approved}, nil)
				return loanMock
			},
			mockRepaymentRepo: func(ctrl *gomock.Controller) *rMock.MockRepaymentRepo {
				repaymentMock := rMock.NewMockRepaymentRepo(ctrl)
				repaymentMock.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(entities.Repayments{Status: constants.RepaymentStatus.Pending}, errors.New("dummy"))
				return repaymentMock
			},
			want:    responses.LoanResp{},
			wantErr: true,
		},
		{
			name: "could not get loan",
			mockLoanRepo: func(ctrl *gomock.Controller) *rMock.MockLoanRepo {
				loanMock := rMock.NewMockLoanRepo(ctrl)
				loanMock.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(entities.Loans{Amount: 100, Term: 1, Status: constants.LoanStatus.Approved}, errors.New("dummy"))
				return loanMock
			},
			mockRepaymentRepo: func(ctrl *gomock.Controller) *rMock.MockRepaymentRepo {
				repaymentMock := rMock.NewMockRepaymentRepo(ctrl)
				return repaymentMock
			},
			want:    responses.LoanResp{},
			wantErr: true,
		},
		{
			name: "repayment already paid",
			mockLoanRepo: func(ctrl *gomock.Controller) *rMock.MockLoanRepo {
				loanMock := rMock.NewMockLoanRepo(ctrl)
				loanMock.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(entities.Loans{Amount: 100, Term: 1, Status: constants.LoanStatus.Approved}, nil)
				return loanMock
			},
			mockRepaymentRepo: func(ctrl *gomock.Controller) *rMock.MockRepaymentRepo {
				repaymentMock := rMock.NewMockRepaymentRepo(ctrl)
				repaymentMock.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(entities.Repayments{Status: constants.RepaymentStatus.Paid}, nil)
				return repaymentMock
			},
			want:    responses.LoanResp{},
			wantErr: true,
		},
		{
			name: "incorrect repayment amount",
			mockLoanRepo: func(ctrl *gomock.Controller) *rMock.MockLoanRepo {
				loanMock := rMock.NewMockLoanRepo(ctrl)
				loanMock.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(entities.Loans{Amount: 100, Term: 1, Status: constants.LoanStatus.Approved}, nil)
				return loanMock
			},
			mockRepaymentRepo: func(ctrl *gomock.Controller) *rMock.MockRepaymentRepo {
				repaymentMock := rMock.NewMockRepaymentRepo(ctrl)
				repaymentMock.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(entities.Repayments{Amount: 1, Status: constants.RepaymentStatus.Pending}, nil)
				return repaymentMock
			},
			want:    responses.LoanResp{},
			wantErr: true,
		},
		{
			name: "incorrect repayment id",
			mockLoanRepo: func(ctrl *gomock.Controller) *rMock.MockLoanRepo {
				loanMock := rMock.NewMockLoanRepo(ctrl)
				loanMock.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(entities.Loans{Amount: 100, Term: 1, Status: constants.LoanStatus.Approved}, nil)
				return loanMock
			},
			mockRepaymentRepo: func(ctrl *gomock.Controller) *rMock.MockRepaymentRepo {
				repaymentMock := rMock.NewMockRepaymentRepo(ctrl)
				repaymentMock.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(entities.Repayments{LoanID: 1, Status: constants.RepaymentStatus.Pending}, nil)
				return repaymentMock
			},
			want:    responses.LoanResp{},
			wantErr: true,
		},
		{
			name: "loan not approved",
			mockLoanRepo: func(ctrl *gomock.Controller) *rMock.MockLoanRepo {
				loanMock := rMock.NewMockLoanRepo(ctrl)
				loanMock.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(entities.Loans{Amount: 100, Term: 1, Status: constants.LoanStatus.Pending}, nil)
				return loanMock
			},
			mockRepaymentRepo: func(ctrl *gomock.Controller) *rMock.MockRepaymentRepo {
				repaymentMock := rMock.NewMockRepaymentRepo(ctrl)
				return repaymentMock
			},
			want:    responses.LoanResp{},
			wantErr: true,
		},
		{
			name: "happy case",
			mockLoanRepo: func(ctrl *gomock.Controller) *rMock.MockLoanRepo {
				loanMock := rMock.NewMockLoanRepo(ctrl)
				loanMock.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(entities.Loans{Amount: 100, Term: 1, Status: constants.LoanStatus.Approved}, nil)
				loanMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(entities.Loans{Amount: 100, Term: 1, Status: constants.LoanStatus.Approved}, nil)
				return loanMock
			},
			mockRepaymentRepo: func(ctrl *gomock.Controller) *rMock.MockRepaymentRepo {
				repaymentMock := rMock.NewMockRepaymentRepo(ctrl)
				repaymentMock.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(entities.Repayments{Status: constants.RepaymentStatus.Pending}, nil)
				repaymentMock.EXPECT().GetAllRepaymentsForLoanID(gomock.Any(), gomock.Any()).Return([]entities.Repayments{{Amount: 100}}, nil).AnyTimes()
				repaymentMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(entities.Repayments{Amount: 100}, nil)
				return repaymentMock
			},
			want:    responses.LoanResp{Amount: 100, Term: 1, Repayments: []responses.Repayments{{Amount: 100}}, Status: constants.LoanStatus.Approved},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tt.fields.loanRepo = tt.mockLoanRepo(ctrl)
			tt.fields.repaymentRepo = tt.mockRepaymentRepo(ctrl)

			l := NewLoanService(
				tt.fields.loanRepo,
				tt.fields.repaymentRepo,
			)
			got, err := l.RepayLoan(tt.args.ctx, tt.args.req, tt.args.loanID, tt.args.repaymentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("loanService.RepayLoan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loanService.RepayLoan() = %v, want %v", got, tt.want)
			}
		})
	}
}
