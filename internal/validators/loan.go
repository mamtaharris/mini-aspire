package validators

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mamtaharris/mini-aspire/internal/constants"
	"github.com/mamtaharris/mini-aspire/internal/models/requests"
)

type loanReqValidator struct {
	validator ValidatorInterface
}

func NewLoanValidator(validator ValidatorInterface) LoanReqValidatorInterface {
	return &loanReqValidator{
		validator: validator,
	}
}

//go:generate mockgen -package mocks -source=loan.go -destination=mocks/loan_mocks.go
type LoanReqValidatorInterface interface {
	ValidateCreateLoanReq(ctx *gin.Context) (requests.CreateLoanReq, error)
	ValidateUpdateLoanReq(ctx *gin.Context) (requests.UpdateLoanReq, int, error)
	ValidateGetLoanReq(ctx *gin.Context) (int, error)
	ValidateRepayLoanReq(ctx *gin.Context) (requests.RepayLoanReq, int, int, error)
}

func (v *loanReqValidator) ValidateCreateLoanReq(ctx *gin.Context) (requests.CreateLoanReq, error) {
	var reqBody requests.CreateLoanReq
	err := v.validator.ValidateUnknownParams(&reqBody, ctx)
	if err != nil {
		return requests.CreateLoanReq{}, err
	}
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		return requests.CreateLoanReq{}, err
	}
	return reqBody, nil
}

func (v *loanReqValidator) ValidateUpdateLoanReq(ctx *gin.Context) (requests.UpdateLoanReq, int, error) {
	var reqBody requests.UpdateLoanReq
	err := v.validator.ValidateUnknownParams(&reqBody, ctx)
	if err != nil {
		return requests.UpdateLoanReq{}, 0, err
	}
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		return requests.UpdateLoanReq{}, 0, err
	}
	if reqBody.Status != constants.LoanStatus.Approved && reqBody.Status != constants.LoanStatus.Rejected {
		return requests.UpdateLoanReq{}, 0, errors.New("invalid status")
	}
	loanID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return requests.UpdateLoanReq{}, 0, err
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
	err := v.validator.ValidateUnknownParams(&reqBody, ctx)
	if err != nil {
		return requests.RepayLoanReq{}, 0, 0, err
	}
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		return requests.RepayLoanReq{}, 0, 0, err
	}
	loanID, err := strconv.Atoi(ctx.Param("loanID"))
	if err != nil {
		return requests.RepayLoanReq{}, 0, 0, err
	}
	repaymentID, err := strconv.Atoi(ctx.Param("repaymentID"))
	if err != nil {
		return requests.RepayLoanReq{}, 0, 0, err
	}
	return reqBody, loanID, repaymentID, nil
}
