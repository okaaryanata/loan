package domain

type (
	Loan struct {
		LoanID             int64   `json:"loanID"`
		UserID             int64   `json:"userID"`
		Principal          float64 `json:"principal"`
		InterestRate       float64 `json:"interestRate"`
		TotalWeeks         int     `json:"totalWeeks"`
		WeeklyRepayment    float64 `json:"weeklyRepayment"`
		OutstandingBalance float64 `json:"outstandingBalance"`
		MissedPayments     int     `json:"missedPayments"`
		IsDelinquent       bool    `json:"isDelinquent"`
	}

	LoanRepayment struct {
		LoanRepaymentID int64   `json:"loanRepaymentID"`
		LoanID          int64   `json:"loanID"`
		Week            int     `json:"week"`
		Amount          float64 `json:"amount"`
		Paid            bool    `json:"paid"`
	}

	LoanRequest struct {
		UserID    int64   `json:"userID"`
		Principal float64 `json:"principal"`
		Interest  float64 `json:"interest"`
		Weeks     int     `json:"weeks"`
	}

	MakePaymentRequest struct {
		UserID int64   `json:"userID"`
		LoanID int64   `json:"loanID"`
		Week   int     `json:"week"`
		Amount float64 `json:"amount"`
	}
)
