package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	loanS "github.com/mamtaharris/mini-aspire/internal/services/loan"
	loanV "github.com/mamtaharris/mini-aspire/internal/validators/loan"
)

type LoanHandler struct {
	loanSvc          loanS.LoanService
	loanReqValidator loanV.LoanReqValidatorInterface
}

func NewLoanHandler(loanSvc loanS.LoanService, loanReqValidator loanV.LoanReqValidatorInterface) *LoanHandler {
	return &LoanHandler{
		loanSvc:          loanSvc,
		loanReqValidator: loanReqValidator,
	}
}

func (h *LoanHandler) CreateLoanHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "success")
}
