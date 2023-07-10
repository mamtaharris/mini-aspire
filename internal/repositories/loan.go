package repositories

import (
	"context"

	"github.com/mamtaharris/mini-aspire/config"
	"github.com/mamtaharris/mini-aspire/internal/models/entities"
	"gorm.io/gorm"
)

type loanRepo struct {
	writeDB *gorm.DB
	readDB  *gorm.DB
}

func NewLoanRepo(writeDB, readDB *gorm.DB) LoanRepo {
	return &loanRepo{
		writeDB: writeDB,
		readDB:  readDB,
	}
}

//go:generate mockgen -package mocks -source=loan_interface.go -destination=mocks/loan_interface_mocks.go
type LoanRepo interface {
	Create(ctx context.Context, loan entities.Loans) (entities.Loans, error)
	GetByID(ctx context.Context, loanID int) (entities.Loans, error)
	Update(ctx context.Context, loan entities.Loans) (entities.Loans, error)
}

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
