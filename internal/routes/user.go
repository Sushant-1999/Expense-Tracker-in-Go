package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"expense-tracker-api/internal/utils"
	"expense-tracker-api/models"

	"github.com/go-chi/chi/v5"
)

// Router for user endpoints
func UserRouter() chi.Router {
	// Create new router
	router := chi.NewRouter()

	// sub-routes
	router.Get("/", listUsers)
	router.Post("/", createNewUser)
	router.Get("/{username}", displayUser)
	router.Put("/{username}", editUser)
	router.Get("/{username}/expenses", listExpenses)

	return router
}

// Get a list of users
// GetUsers godoc
// @Summary      Shows list of User for expenses
// @Description  get list of users
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @router  	 /user/   [get]
func listUsers(w http.ResponseWriter, r *http.Request) {
	// Get existing users
	users := utils.GetUsers()

	// Encode users
	bytes, err := json.Marshal(users)

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

// Get details of a specific user by username
func displayUser(w http.ResponseWriter, r *http.Request) {
	// Queried username
	username := strings.ToLower(chi.URLParam(r, "username"))

	// Get existing users
	users := utils.GetUsers().UserList

	// Check if user exists
	for i := 0; i < len(users); i++ {
		if strings.EqualFold(users[i].Username, username) {
			// Encode user
			bytes, err := json.Marshal(users[i])

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

	// Send 404 if user not found
	ResourceNotFound(w, r)
}

// Create a new user record
func createNewUser(w http.ResponseWriter, r *http.Request) {
	// Get params from request
	queryParams := r.URL.Query()

	// Process the params
	username := strings.ToLower(queryParams.Get("username"))
	fname := queryParams.Get("firstName")
	lname := queryParams.Get("lastName")
	email := queryParams.Get("email")

	// Get existing users
	users := utils.GetUsers()
	userList := users.UserList

	// Check if user already exists
	for i := 0; i < len(userList); i++ {
		if strings.EqualFold(username, userList[i].Username) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("A user with this username already exists, please try again!"))
			return
		} else if strings.EqualFold(email, userList[i].Email) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("A user with this email already exists, please try again!"))
			return
		}
	}

	// Create a new User struct
	newUser := models.User{Username: username, FirstName: fname, LastName: lname, Email: email}

	// Update the user list
	users.UserList = append(users.UserList, newUser)

	// Save the user list as json
	if utils.SaveUsers(users) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Successfully added new user!"))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to add new user, please try again!"))
	}
}

// Update details of a specific user by username
func editUser(w http.ResponseWriter, r *http.Request) {
	// Username to be edited
	username := strings.ToLower(chi.URLParam(r, "username"))

	// Get params from request
	queryParams := r.URL.Query()

	// Process the params
	fname := queryParams.Get("firstName")
	lname := queryParams.Get("lastName")
	email := queryParams.Get("email")

	// Get existing users
	users := utils.GetUsers()

	// Check if user already exists
	for i := 0; i < len(users.UserList); i++ {
		if strings.EqualFold(username, users.UserList[i].Username) {
			// Replace First Name
			if !(fname == "" || fname == users.UserList[i].FirstName) {
				users.UserList[i].FirstName = fname
			}

			// Replace Last Name
			if !(lname == "" || lname == users.UserList[i].LastName) {
				users.UserList[i].LastName = lname
			}

			// Replace Email
			if !(email == "" || email == users.UserList[i].Email) {
				users.UserList[i].Email = email
			}

			// Save the user list as json
			if utils.SaveUsers(users) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Successfully edited user!"))
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Unable to edit user, please try again!"))
			}

			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Unable to locate user, please try again!"))
}
