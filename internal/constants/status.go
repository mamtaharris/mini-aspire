package constants

var LoanStatus = struct {
	Pending  string
	Approved string
	Rejected string
	Paid     string
}{
	Pending:  "PENDING",
	Approved: "APPROVED",
	Rejected: "REJECTED",
	Paid:     "PAID",
}

var RepaymentStatus = struct {
	Pending string
	Paid    string
}{
	Pending: "PENDING",
	Paid:    "PAID",
}
