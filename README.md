# Expense Tracker API

A Go API for managing personal expenses.

## Overview

This API provides a set of endpoints for creating, reading, updating, and deleting expenses, as well as generating reports to track spending habits. It's written in Go and designed to be lightweight, efficient, and easy to integrate into personal finance applications.

## Features

- **CRUD operations for expenses:** Create, read, update, and delete expenses with ease.
- **Expense categorization:** Organize expenses by category for better insights.
- **Date-based filtering:** Filter expenses by date range for specific analysis.

## Getting Started

Follow these steps to set up and run the API on your local machine:

1. **Install Go:** Ensure you have Go installed on your system. See [here](https://go.dev/doc/install) for more details.

2. **Clone the repository:**

```bash
git clone https://expense-tracker-api.git
```

3. **Regenerate Sample Expenses** (optional):

```bash
python ./data/generate_expenses.py
```

4. **Install Dependencies:**

```bash
go get
```

5. **Run the Application:**

```bash
go run ./cmd/expense-tracker-api/main.go
```

6. **Explore the API:**

```bash
Open your browser and visit http://localhost:8080/ to view the API documentation.
```

## Endpoints

### `/expense`

- `GET /`: Get a list of all expenses.
- `POST /`: Create a new expense.
- `GET /filter`: Filter expenses by category and date.
- `GET /summary`: Get a summary of expenses.
- `GET /{id}`: Get details of a specific expense.
- `PUT /{id}`: Update details of a specific expense.
- `DELETE /{id}`: Delete a specific expense.

### `/category`

- `GET /`: Get a list of all expense categories.
- `GET /search`: Find categories with provided query.
- `GET /{id}`: Get details of a specific category.

### `/user`

- `GET /`: Get a list of all users.
- `POST /`: Create a new user.
- `GET /{username}`: Get details of a specific user.
- `PUT /{username}`: Update details of a specific user.
- `GET /{username}/expenses`: Get expenses of a specific user.

## Documentation

Detailed endpoint documentation is available at [docs/](docs/)
