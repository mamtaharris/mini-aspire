package requests

type CreateLoanReq struct {
	Amount float64 `json:"amount"`
	Term   int     `json:"term"`
}

type UpdateLoanReq struct {
	Status string `json:"status"`
}

type RepayLoanReq struct {
	Amount float64 `json:"amount"`
}
