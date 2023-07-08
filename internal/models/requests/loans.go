package requests

type CreateLoanReq struct {
	Amount float64 `json:"amount"`
	Term   int     `json:"term"`
}
