package loan

import (
	"context"

	"github.com/mamtaharris/mini-aspire/internal/constants"
	"github.com/mamtaharris/mini-aspire/internal/models/entities"
	"github.com/mamtaharris/mini-aspire/internal/models/requests"
	"github.com/mamtaharris/mini-aspire/internal/models/responses"
)

func (l *loanService) CreateLoan(ctx context.Context, req requests.CreateLoanReq) (responses.LoanResp, error) {
	loan, err := l.createLoanEntity(ctx, req)
	if err != nil {
		return responses.LoanResp{}, err
	}
	err = l.createRepaymentEntity(ctx, loan)
	if err != nil {
		return responses.LoanResp{}, err
	}
	resp, err := l.getResponseObject(ctx, loan)
	if err != nil {
		return responses.LoanResp{}, err
	}
	return resp, nil
}

func (l *loanService) UpdateLoan(ctx context.Context, req requests.UpdateLoanReq, loanID int) (responses.LoanResp, error) {
	loan, err := l.loanRepo.GetByID(ctx, loanID)
	if err != nil {
		return responses.LoanResp{}, nil
	}
	loan.Status = req.Status
	loan, err = l.loanRepo.Update(ctx, loan)
	if err != nil {
		return responses.LoanResp{}, nil
	}
	resp, err := l.getResponseObject(ctx, loan)
	if err != nil {
		return responses.LoanResp{}, err
	}
	return resp, nil
}

func (l *loanService) GetLoan(ctx context.Context, loanID int) (responses.LoanResp, error) {
	loan, err := l.loanRepo.GetByID(ctx, loanID)
	if err != nil {
		return responses.LoanResp{}, nil
	}
	resp, err := l.getResponseObject(ctx, loan)
	if err != nil {
		return responses.LoanResp{}, err
	}
	return resp, nil
}

func (l *loanService) createLoanEntity(ctx context.Context, req requests.CreateLoanReq) (entities.Loans, error) {
	loan := entities.Loans{
		Amount:  req.Amount,
		Term:    req.Term,
		Status:  constants.LoanStatus.Pending,
		UsersID: 123, //TODO: set from jwt
	}
	loan, err := l.loanRepo.Create(ctx, loan)
	if err != nil {
		return loan, err
	}
	return loan, nil
}

func (l *loanService) createRepaymentEntity(ctx context.Context, loan entities.Loans) error {
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
			return err
		}
	}
	return nil
}

func (l *loanService) getResponseObject(ctx context.Context, loan entities.Loans) (responses.LoanResp, error) {
	repayments, err := l.repaymentRepo.GetAllRepaymentsForLoanID(ctx, loan.ID)
	if err != nil {
		return responses.LoanResp{}, err
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

	resp := responses.LoanResp{
		ID:         loan.ID,
		Amount:     loan.Amount,
		Term:       loan.Term,
		Repayments: repaymentsResp,
		Status:     loan.Status,
	}
	return resp, nil
}
