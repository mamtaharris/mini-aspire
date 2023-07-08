package loan

import (
	"context"

	"github.com/mamtaharris/mini-aspire/internal/models/entities"
)

func (r *loanRepo) Create(ctx context.Context, loans entities.Loans) (entities.Loans, error) {
	return entities.Loans{}, nil
}
