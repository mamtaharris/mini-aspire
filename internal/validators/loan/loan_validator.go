package loan

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mamtaharris/mini-aspire/internal/constants"
	"github.com/mamtaharris/mini-aspire/internal/models/requests"
	"github.com/mamtaharris/mini-aspire/internal/validators"
)

func (v *loanReqValidator) ValidateCreateLoanReq(ctx *gin.Context) (requests.CreateLoanReq, error) {
	var reqBody requests.CreateLoanReq
	err := validators.ValidateUnknownParams(&reqBody, ctx)
	if err != nil {
		return reqBody, err
	}
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		return reqBody, err
	}
	return reqBody, nil
}

func (v *loanReqValidator) ValidateUpdateLoanReq(ctx *gin.Context) (requests.UpdateLoanReq, int, error) {
	var reqBody requests.UpdateLoanReq
	err := validators.ValidateUnknownParams(&reqBody, ctx)
	if err != nil {
		return reqBody, 0, err
	}
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		return reqBody, 0, err
	}
	if reqBody.Status != constants.LoanStatus.Approved && reqBody.Status != constants.LoanStatus.Rejected {
		return reqBody, 0, errors.New("invalid status")
	}
	loanID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return reqBody, 0, err
	}
	return reqBody, loanID, nil
}

func (v *loanReqValidator) ValidateGetLoanReq(ctx *gin.Context) (int, error) {
	loanID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return 0, err
	}
	return loanID, nil
}

func (v *loanReqValidator) ValidateRepayLoanReq(ctx *gin.Context) (requests.RepayLoanReq, int, int, error) {
	var reqBody requests.RepayLoanReq
	err := validators.ValidateUnknownParams(&reqBody, ctx)
	if err != nil {
		return reqBody, 0, 0, err
	}
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		return reqBody, 0, 0, err
	}
	loanID, err := strconv.Atoi(ctx.Param("loanID"))
	if err != nil {
		return reqBody, 0, 0, err
	}
	repaymentID, err := strconv.Atoi(ctx.Param("repaymentID"))
	if err != nil {
		return reqBody, 0, 0, err
	}
	return reqBody, loanID, repaymentID, nil
}
