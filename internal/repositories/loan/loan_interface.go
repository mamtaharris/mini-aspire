package loan

import (
	"context"

	"github.com/mamtaharris/mini-aspire/internal/models/entities"
	"gorm.io/gorm"
)

type loanRepo struct {
	writeDB *gorm.DB
	readDB  *gorm.DB
}

func NewRepo(writeDB, readDB *gorm.DB) LoanRepo {
	return &loanRepo{
		writeDB: writeDB,
		readDB:  readDB,
	}
}

//go:generate mockgen -package mocks -source=loan_interface.go -destination=mocks/loan_interface_mocks.go
type LoanRepo interface {
	Create(ctx context.Context, loan entities.Loans) (entities.Loans, error)
}
