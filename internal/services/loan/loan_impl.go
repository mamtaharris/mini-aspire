package loan

import (
	"context"

	"github.com/mamtaharris/mini-aspire/internal/constants"
	"github.com/mamtaharris/mini-aspire/internal/models/entities"
	"github.com/mamtaharris/mini-aspire/internal/models/requests"
	"github.com/mamtaharris/mini-aspire/internal/models/responses"
)

func (l *loanService) CreateLoan(ctx context.Context, req requests.CreateLoanReq) (responses.CreateLoanResp, error) {
	loan := entities.Loans{
		Amount:  req.Amount,
		Term:    req.Term,
		Status:  constants.LoanStatus.Pending,
		UsersID: 123, //TODO: set from jwt
	}
	loan, err := l.loanRepo.Create(ctx, loan)
	if err != nil {
		return responses.CreateLoanResp{}, err
	}
	repaymentAmount := loan.Amount / float64(loan.Term)
	for i := 0; i < loan.Term; i++ {
		repayment := entities.Repayments{
			LoanID:  loan.ID,
			Amount:  repaymentAmount,
			Status:  constants.RepaymentStatus.Pending,
			UsersID: 123, //TODO: set from jwt
		}
		repayment, err := l.repaymentRepo.Create(ctx, repayment)
		if err != nil {
			return responses.CreateLoanResp{}, err
		}
	}

	repayments, err := l.repaymentRepo.GetAllRepaymentsForLoanID(ctx, loan.ID)
	if err != nil {
		return responses.CreateLoanResp{}, err
	}
	var repaymentsResp []responses.Repayments
	for _, repayment := range repayments {
		repaymentResp := responses.Repayments{
			ID:     repayment.ID,
			Amount: repayment.Amount,
			Status: repayment.Status,
		}
		repaymentsResp = append(repaymentsResp, repaymentResp)
	}

	resp := responses.CreateLoanResp{
		ID:         loan.ID,
		Amount:     loan.Amount,
		Term:       loan.Term,
		Repayments: repaymentsResp,
		Status:     loan.Status,
	}
	return resp, nil
}
