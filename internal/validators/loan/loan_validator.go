package loan

import (
	"github.com/gin-gonic/gin"
	"github.com/mamtaharris/mini-aspire/internal/models/requests"
)

func (v *loanReqValidator) ValidateCreateLoanReq(ctx *gin.Context) (requests.CreateLoanReq, error) {
	return requests.CreateLoanReq{}, nil
}
