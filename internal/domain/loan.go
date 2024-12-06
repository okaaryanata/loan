package domain

type (
	Loan struct {
		LoanID             int64   `json:"loanID"`
		UserID             int64   `json:"userID"`
		Principal          int     `json:"principal"`
		InterestRate       float64 `json:"interestRate"`
		TotalWeeks         int     `json:"totalWeeks"`
		WeeklyRepayment    int     `json:"weeklyRepayment"`
		OutstandingBalance int     `json:"outstandingBalance"`
		MissedPayments     int     `json:"missedPayments"`
		IsDelinquent       bool    `json:"isDelinquent"`
	}

	LoanRepayment struct {
		LoanRepaymentID int64   `json:"loanRepaymentID"`
		LoanID          int64   `json:"loanID"`
		Week            int     `json:"week"`
		Amount          float64 `json:"amount"`
	}
)
