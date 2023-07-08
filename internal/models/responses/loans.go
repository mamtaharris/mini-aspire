package responses

type LoanResp struct {
	ID         int          `json:"loan_id"`
	Amount     float64      `json:"amount"`
	Term       int          `json:"term"`
	Repayments []Repayments `json:"repayments"`
	Status     string       `json:"status"`
}
