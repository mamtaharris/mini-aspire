package loan

import (
	"context"

	"github.com/mamtaharris/mini-aspire/internal/models/requests"
	"github.com/mamtaharris/mini-aspire/internal/models/responses"
	loanR "github.com/mamtaharris/mini-aspire/internal/repositories/loan"
	repaymentR "github.com/mamtaharris/mini-aspire/internal/repositories/repayment"
)

type loanService struct {
	loanRepo      loanR.LoanRepo
	repaymentRepo repaymentR.RepaymentRepo
}

func NewService(loanRepo loanR.LoanRepo, repaymentRepo repaymentR.RepaymentRepo) LoanService {
	return &loanService{
		loanRepo:      loanRepo,
		repaymentRepo: repaymentRepo,
	}
}

//go:generate mockgen -package mocks -source=loan_interface.go -destination=mocks/loan_interface_mocks.go
type LoanService interface {
	CreateLoan(ctx context.Context, req requests.CreateLoanReq) (responses.CreateLoanResp, error)
}
