package repositories

import (
	"context"

	"github.com/mamtaharris/mini-aspire/config"
	"github.com/mamtaharris/mini-aspire/internal/models/entities"
	"gorm.io/gorm"
)

type repaymentRepo struct {
	writeDB *gorm.DB
	readDB  *gorm.DB
}

func NewRepaymentRepo(writeDB, readDB *gorm.DB) RepaymentRepo {
	return &repaymentRepo{
		writeDB: writeDB,
		readDB:  readDB,
	}
}

//go:generate mockgen -package mocks -source=repayment.go -destination=mocks/repayment_mocks.go
type RepaymentRepo interface {
	Create(ctx context.Context, repayment entities.Repayments) (entities.Repayments, error)
	GetAllRepaymentsForLoanID(ctx context.Context, loanID int) ([]entities.Repayments, error)
	GetByID(ctx context.Context, repaymentID int) (entities.Repayments, error)
	Update(ctx context.Context, repayment entities.Repayments) (entities.Repayments, error)
}

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

func (r *repaymentRepo) GetByID(ctx context.Context, repaymentID int) (entities.Repayments, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DB.QueryTimeout)
	defer cancel()

	repayments := entities.Repayments{}
	result := r.readDB.WithContext(ctx).Where("repayment_id = ?", repaymentID).Find(&repayments)
	if result.Error != nil {
		return repayments, result.Error
	}
	return repayments, nil
}

func (r *repaymentRepo) Update(ctx context.Context, repayments entities.Repayments) (entities.Repayments, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DB.QueryTimeout)
	defer cancel()

	result := r.writeDB.WithContext(ctx).
		Model(&repayments).
		Where("repayment_id = ?", repayments.ID).
		Updates(repayments)
	if result.Error != nil {
		return repayments, result.Error
	}
	return repayments, nil
}
