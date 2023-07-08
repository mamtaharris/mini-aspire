package responses

type Repayments struct {
	ID     int     `json:"repayment_id"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
}
