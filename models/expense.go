package models

type Expense struct {
	TransactionID string `json:"transactionID"`
	Amount        int    `json:"amount"`
	Date          string `json:"date"`
	CategoryID    string `json:"category"`
	UserID        string `json:"user"`
}

type Expenses struct {
	ExpenseList []Expense `json:"expenses"`
}

type ExpenseSummary struct {
	TotalTransactions int     `json:"totalTransactions"`
	TotalAmount       int     `json:"totalAmount"`
	AverageExpense    float64 `json:"averageExpense"`
}
