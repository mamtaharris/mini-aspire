package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mamtaharris/mini-aspire/internal/models/responses"
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

func (h *LoanHandler) CreateLoanHandler(ctx *gin.Context) {
	req, err := h.loanReqValidator.ValidateCreateLoanReq(ctx)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, responses.ErrorResp{Error: err.Error()})
		return
	}
	response, err := h.loanSvc.CreateLoan(ctx, req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResp{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (h *LoanHandler) UpdateLoanHandler(ctx *gin.Context) {
	req, loanID, err := h.loanReqValidator.ValidateUpdateLoanReq(ctx)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, responses.ErrorResp{Error: err.Error()})
		return
	}
	response, err := h.loanSvc.UpdateLoan(ctx, req, loanID)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResp{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (h *LoanHandler) GetLoanHandler(ctx *gin.Context) {
	loanID, err := h.loanReqValidator.ValidateGetLoanReq(ctx)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, responses.ErrorResp{Error: err.Error()})
		return
	}
	response, err := h.loanSvc.GetLoan(ctx, loanID)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResp{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}
