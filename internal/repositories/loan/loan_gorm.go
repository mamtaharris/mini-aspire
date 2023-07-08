package loan

import (
	"context"

	"github.com/mamtaharris/mini-aspire/config"
	"github.com/mamtaharris/mini-aspire/internal/models/entities"
)

func (r *loanRepo) Create(ctx context.Context, loan entities.Loans) (entities.Loans, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DB.QueryTimeout)
	defer cancel()

	result := r.writeDB.WithContext(ctx).Create(&loan)
	if result.Error != nil {
		return loan, result.Error
	}
	return loan, nil
}

func (r *loanRepo) GetByID(ctx context.Context, loanID int) (entities.Loans, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DB.QueryTimeout)
	defer cancel()

	loan := entities.Loans{}
	result := r.readDB.WithContext(ctx).Where("loan_id = ?", loanID).Find(&loan)
	if result.Error != nil {
		return loan, result.Error
	}
	return loan, nil
}

func (r *loanRepo) Update(ctx context.Context, loan entities.Loans) (entities.Loans, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DB.QueryTimeout)
	defer cancel()

	result := r.writeDB.WithContext(ctx).
		Model(&loan).
		Where("loan_id = ?", loan.ID).
		Updates(loan)
	if result.Error != nil {
		return loan, result.Error
	}
	return loan, nil
}
