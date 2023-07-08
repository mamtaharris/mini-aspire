package responses

type CreateLoanResp struct {
	ID     string  `json:"loan_id"`
	Amount float64 `json:"amount"`
	Term   int     `json:"term"`
	Status string  `json:"status"`
}
