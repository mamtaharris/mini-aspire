package repayment

import (
	"context"

	"github.com/mamtaharris/mini-aspire/config"
	"github.com/mamtaharris/mini-aspire/internal/models/entities"
)

func (r *repaymentRepo) Create(ctx context.Context, loan entities.Repayments) (entities.Repayments, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DB.QueryTimeout)
	defer cancel()

	result := r.writeDB.WithContext(ctx).Create(&loan)
	if result.Error != nil {
		return loan, result.Error
	}
	return loan, nil
}

func (r *repaymentRepo) GetAllRepaymentsForLoanID(ctx context.Context, loanID int) ([]entities.Repayments, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DB.QueryTimeout)
	defer cancel()

	repayments := []entities.Repayments{}
	result := r.readDB.WithContext(ctx).Where("loan_id = ?", loanID).Find(&repayments)
	if result.Error != nil {
		return repayments, result.Error
	}
	return repayments, nil
}
