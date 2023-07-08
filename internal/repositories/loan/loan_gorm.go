package loan

import (
	"context"
	"fmt"

	"github.com/mamtaharris/mini-aspire/config"
	"github.com/mamtaharris/mini-aspire/internal/models/entities"
)

func (r *loanRepo) Create(ctx context.Context, loan entities.Loans) (entities.Loans, error) {
	fmt.Println(config.DB.QueryTimeout)
	ctx, cancel := context.WithTimeout(ctx, config.DB.QueryTimeout)
	defer cancel()

	result := r.writeDB.WithContext(ctx).Create(&loan)
	if result.Error != nil {
		return loan, result.Error
	}
	return loan, nil
}
