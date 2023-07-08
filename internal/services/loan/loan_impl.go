package loan

import (
	"context"

	"github.com/mamtaharris/mini-aspire/internal/models/requests"
	"github.com/mamtaharris/mini-aspire/internal/models/responses"
)

func (l *loanService) CreateLoan(ctx context.Context, req requests.CreateLoanReq) (responses.CreateLoanResp, error) {
	return responses.CreateLoanResp{}, nil
}
