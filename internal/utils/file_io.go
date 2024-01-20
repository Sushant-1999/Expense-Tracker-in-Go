package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"expense-tracker-api/models"
)

// Get the documentation text file
func GetDocTXT() []byte {
	// Read file
	bytes, err := os.ReadFile("./data/docs.txt")

	if err != nil {
		fmt.Println("Error loading file: ", err)
		return nil
	}

	return bytes
}

// Get the documentation html file
func GetDocHTML() []byte {
	// Read file
	bytes, err := os.ReadFile("./docs/index.html")

	if err != nil {
		fmt.Println("Error loading file: ", err)
		return nil
	}

	return bytes
}

// Read and parse the categories json
func GetCategories() models.Categories {
	// Read file
	bytes, err := os.ReadFile("./docs/categories.json")

	if err != nil {
		fmt.Println("Error loading file: ", err)
	}

	// Parse json
	var cat models.Categories
	err = json.Unmarshal(bytes, &cat)

	if err != nil {
		fmt.Println("Error parsing file: ", err)
	}

	return cat
}

// Read and parse the users json
func GetUsers() models.Users {
	// Read file
	bytes, err := os.ReadFile("./data/users.json")

	if err != nil {
		fmt.Println("Error loading file: ", err)
	}

	// Parse json
	var users models.Users
	err = json.Unmarshal(bytes, &users)

	if err != nil {
		fmt.Println("Error parsing file: ", err)
	}

	return users
}

// Save the users as json
func SaveUsers(users models.Users) bool {
	// Encode the data
	bytes, err := json.MarshalIndent(users, "", "\t")

	if err != nil {
		fmt.Println("Error saving file: ", err)
		return false
	}

	// Save json
	err = os.WriteFile("./data/users.json", bytes, 0666)

	if err != nil {
		fmt.Println("Error saving file: ", err)
		return false
	}

	return true
}

// Read and parse the expenses json
func GetExpenses() models.Expenses {
	// Read file
	bytes, err := os.ReadFile("./data/expenses.json")

	if err != nil {
		fmt.Println("Error loading file: ", err)
	}

	// Parse json
	var expenses models.Expenses
	err = json.Unmarshal(bytes, &expenses)

	if err != nil {
		fmt.Println("Error parsing file: ", err)
	}

	return expenses
}

// Save the expenses as json
func SaveExpenses(expenses models.Expenses) bool {
	// Encode the data
	bytes, err := json.MarshalIndent(expenses, "", "\t")

	if err != nil {
		fmt.Println("Error saving file: ", err)
		return false
	}

	// Save json
	err = os.WriteFile("./data/expenses.json", bytes, 0666)

	if err != nil {
		fmt.Println("Error saving file: ", err)
		return false
	}

	return true
}
