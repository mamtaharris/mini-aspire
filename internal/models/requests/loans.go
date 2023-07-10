package requests

type CreateLoanReq struct {
	Amount float64 `json:"amount" binding:"required"`
	Term   int     `json:"term" binding:"required"`
}

type UpdateLoanReq struct {
	Status string `json:"status" binding:"required"`
}

type RepayLoanReq struct {
	Amount float64 `json:"amount" binding:"required"`
}
