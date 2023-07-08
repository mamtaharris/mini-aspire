package loan

import (
	"github.com/gin-gonic/gin"
	"github.com/mamtaharris/mini-aspire/internal/models/requests"
)

type loanReqValidator struct{}

func NewValidator() LoanReqValidatorInterface {
	return &loanReqValidator{}
}

//go:generate mockgen -package mocks -source=loan_interface.go -destination=mocks/loan_interface_mocks.go
type LoanReqValidatorInterface interface {
	ValidateCreateLoanReq(ctx *gin.Context) (requests.CreateLoanReq, error)
	ValidateUpdateLoanReq(ctx *gin.Context) (requests.UpdateLoanReq, int, error)
	ValidateGetLoanReq(ctx *gin.Context) (int, error)
	ValidateRepayLoanReq(ctx *gin.Context) (requests.RepayLoanReq, int, int, error)
}
