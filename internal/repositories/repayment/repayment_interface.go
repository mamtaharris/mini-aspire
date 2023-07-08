package repayment

import (
	"context"

	"github.com/mamtaharris/mini-aspire/internal/models/entities"
	"gorm.io/gorm"
)

type repaymentRepo struct {
	writeDB *gorm.DB
	readDB  *gorm.DB
}

func NewRepo(writeDB, readDB *gorm.DB) RepaymentRepo {
	return &repaymentRepo{
		writeDB: writeDB,
		readDB:  readDB,
	}
}

//go:generate mockgen -package mocks -source=repayment_interface.go -destination=mocks/repayment_interface_mocks.go
type RepaymentRepo interface {
	Create(ctx context.Context, repayment entities.Repayments) (entities.Repayments, error)
	GetAllRepaymentsForLoanID(ctx context.Context, loanID int) ([]entities.Repayments, error)
}
