package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"expense-tracker-api/internal/utils"
	"expense-tracker-api/models"

	"github.com/go-chi/chi/v5"
)

// Router for expense endpoints
func ExpenseRouter() chi.Router {
	// Create new router
	router := chi.NewRouter()

	// sub-routes
	router.Get("/", listExpenses)
	router.Post("/", createNewExpense)
	router.Get("/filter", filterExpenses)
	router.Get("/summary", summarizeExpenses)
	router.Get("/{id}", displayExpense)
	router.Put("/{id}", editExpense)
	router.Delete("/{id}", deleteExpense)

	return router
}

// Get a list of all expenses
func listExpenses(w http.ResponseWriter, r *http.Request) {
	// Get expenses
	expenses := utils.GetExpenses()

	// Encode expenses
	bytes, err := json.Marshal(expenses)

	// JSON Encoding error handling
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to process your request, please try again!"))
		fmt.Println(err)
	}

	// Send data as json
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

// Create a new expense record
func createNewExpense(w http.ResponseWriter, r *http.Request) {
	// Get params from request
	queryParams := r.URL.Query()

	// Process the params
	transactionID := queryParams.Get("transactionID")
	amount, _ := strconv.Atoi(queryParams.Get("amount"))
	date := queryParams.Get("date")
	categoryID := queryParams.Get("category")
	userID := queryParams.Get("userID")

	// Get existing users
	userList := utils.GetUsers().UserList

	// Check if user exists
	found := false
	for i := 0; i < len(userList); i++ {
		if userList[i].Username == userID {
			found = true
			break
		}
	}

	// Return error if no user exists
	if !found {
		w.Write([]byte("No such user exists, please try again!"))
		return
	}

	// Get existing expenses
	expenses := utils.GetExpenses()
	expenseList := expenses.ExpenseList

	// Check if transaction id already exists
	for i := 0; i < len(expenseList); i++ {
		if expenseList[i].TransactionID == transactionID {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("A transaction with this ID already exists, please try again!"))
			return
		}
	}

	// Create a new Expense struct
	newExpense := models.Expense{TransactionID: transactionID, Amount: amount, Date: date, CategoryID: categoryID, UserID: userID}

	// Update the expense list
	expenses.ExpenseList = append(expenses.ExpenseList, newExpense)

	// Save the expense list as json
	if utils.SaveExpenses(expenses) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Successfully added new transaction!"))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to add new transaction, please try again!"))
	}
}

// Filter expenses by category and date
func filterExpenses(w http.ResponseWriter, r *http.Request) {
	// Get params from request
	queryParams := r.URL.Query()

	// Process the param
	date := queryParams.Get("amount")
	category := queryParams.Get("amount")

	// Get Expenses
	expenses := utils.GetExpenses().ExpenseList

	// Create a new Expense list to store matching expenses in
	var matchExpenses models.Expenses

	// Search for expenses
	for i := 0; i < len(expenses); i++ {
		expense := expenses[i]

		if (date == "" || date == expense.Date) && (category == "" || category == expense.CategoryID) {
			matchExpenses.ExpenseList = append(matchExpenses.ExpenseList, expense)
		}
	}

	// If none found, return 404
	if len(matchExpenses.ExpenseList) == 0 {
		ResourceNotFound(w, r)
		return
	}

	// Encode the new expense list
	bytes, err := json.Marshal(matchExpenses)

	// JSON Encoding error handling
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to process your request, please try again!"))
		fmt.Println(err)
	}

	// Send data as json
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

// Get a summary of expenses
func summarizeExpenses(w http.ResponseWriter, r *http.Request) {
	// Get Expenses
	expenses := utils.GetExpenses().ExpenseList

	// Calculate metrics
	totalTransactions := len(expenses)

	// Total Expense by all users
	totalAmount := 0
	for i := 0; i < totalTransactions; i++ {
		expense := expenses[i]
		totalAmount += expense.Amount
	}

	// Average transaction amount
	averageExpense := float64(totalAmount) / float64(totalTransactions)

	// Create summary from metrics
	summary := models.ExpenseSummary{TotalTransactions: totalTransactions, TotalAmount: totalAmount, AverageExpense: averageExpense}

	// Encode the summary
	bytes, err := json.Marshal(summary)

	// JSON Encoding error handling
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to process your request, please try again!"))
		fmt.Println(err)
	}

	// Send data as json
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

// Get details of a specific expense by ID
func displayExpense(w http.ResponseWriter, r *http.Request) {
	// Queried expense
	id := chi.URLParam(r, "id")

	// Get expenses
	expenses := utils.GetExpenses().ExpenseList

	// Check if expense exists
	for i := 0; i < len(expenses); i++ {
		if expenses[i].TransactionID == id {
			// Encode expense
			bytes, err := json.Marshal(expenses[i])

			// JSON Encoding error handling
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Unable to process your request, please try again!"))
				fmt.Println(err)
			}

			// Send data as json
			w.Header().Set("Content-Type", "application/json")
			w.Write(bytes)

			return
		}
	}

	// Send 404 if expense not found
	ResourceNotFound(w, r)
}

// Update details of a specific expense by ID
func editExpense(w http.ResponseWriter, r *http.Request) {
	// Get params from request
	queryParams := r.URL.Query()

	// Process the params
	transactionID := queryParams.Get("transactionID")
	amount, _ := strconv.Atoi(queryParams.Get("amount"))
	date := queryParams.Get("date")
	categoryID := queryParams.Get("category")
	userID := queryParams.Get("userID")

	// Get expenses
	expenses := utils.GetExpenses()

	for i := 0; i < len(expenses.ExpenseList); i++ {
		if expenses.ExpenseList[i].TransactionID == transactionID {
			// Replace amount
			if !(amount == 0 || amount == expenses.ExpenseList[i].Amount) {
				expenses.ExpenseList[i].Amount = amount
			}

			// Replace date
			if !(date == "" || date == expenses.ExpenseList[i].Date) {
				expenses.ExpenseList[i].Date = date
			}

			// Replace category
			if !(categoryID == "" || categoryID == expenses.ExpenseList[i].CategoryID) {
				expenses.ExpenseList[i].CategoryID = categoryID
			}

			// Replace user
			if !(userID == "" || userID == expenses.ExpenseList[i].UserID) {
				expenses.ExpenseList[i].UserID = userID
			}

			// Save the expense list as json
			if utils.SaveExpenses(expenses) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Successfully edited expense!"))
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Unable to edit expense, please try again!"))
			}

			return
		}
	}
}

// Delete a specific expense by ID
func deleteExpense(w http.ResponseWriter, r *http.Request) {
	// Queried expense
	id := chi.URLParam(r, "id")

	// Get expenses
	expenses := utils.GetExpenses()

	// New expense list
	var newExpenses models.Expenses

	// Check if expense exists
	found := false
	for i := 0; i < len(expenses.ExpenseList); i++ {
		if expenses.ExpenseList[i].TransactionID != id {
			newExpenses.ExpenseList = append(newExpenses.ExpenseList, expenses.ExpenseList[i])
		} else {
			found = true
		}
	}

	// Send 404 if expense not found
	if !found {
		ResourceNotFound(w, r)
		return
	}

	// Save the expense list as json
	if utils.SaveExpenses(expenses) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Successfully added new transaction!"))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to add new transaction, please try again!"))
	}
}
