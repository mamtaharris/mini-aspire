package loan

import (
	"context"

	"github.com/mamtaharris/mini-aspire/internal/models/requests"
	"github.com/mamtaharris/mini-aspire/internal/models/responses"
	repo "github.com/mamtaharris/mini-aspire/internal/repositories/loan"
)

type loanService struct {
	loanRepo repo.LoanRepo
}

func NewService(loanRepo repo.LoanRepo) LoanService {
	return &loanService{
		loanRepo: loanRepo,
	}
}

//go:generate mockgen -package mocks -source=loan_interface.go -destination=mocks/loan_interface_mocks.go
type LoanService interface {
	CreateLoan(ctx context.Context, req requests.CreateLoanReq) (responses.CreateLoanResp, error)
}
