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

// Router for category endpoints
func CategoryRouter() chi.Router {
	// Create new router
	router := chi.NewRouter()

	// sub-routes
	router.Get("/", listCategories)
	router.Get("/search", searchCategory)
	router.Get("/{id}", displayCategory)

	return router
}

func listCategories(w http.ResponseWriter, r *http.Request) {
	// Get categories
	categories := utils.GetCategories()

	// Encode categories
	bytes, err := json.Marshal(categories)

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

// Get details of a specific category by ID
func displayCategory(w http.ResponseWriter, r *http.Request) {
	// Queried category
	id := chi.URLParam(r, "id")

	// Get categories
	categories := utils.GetCategories().CategoryList

	// Check if category exists
	for i := 0; i < len(categories); i++ {
		if strings.EqualFold(categories[i].ID, id) {
			// Encode category
			bytes, err := json.Marshal(categories[i])

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

	// Send 404 if category not found
	ResourceNotFound(w, r)
}

// Find categories with provided query
func searchCategory(w http.ResponseWriter, r *http.Request) {
	// Get params from request
	queryParams := r.URL.Query()

	// Process the param
	searchQuery := strings.ToLower(queryParams.Get("q"))

	// Return 404 on empty request
	if searchQuery == "" {
		ResourceNotFound(w, r)
		return
	}

	// Get categories
	categories := utils.GetCategories().CategoryList

	// Create a new Category list to store matching categories in
	var matchCategories models.Categories

	// Search for categories
	for i := 0; i < len(categories); i++ {
		if strings.Contains(categories[i].ID, searchQuery) {
			matchCategories.CategoryList = append(matchCategories.CategoryList, categories[i])
		} else if strings.Contains(categories[i].Name, searchQuery) {
			matchCategories.CategoryList = append(matchCategories.CategoryList, categories[i])
		} else if strings.Contains(categories[i].Description, searchQuery) {
			matchCategories.CategoryList = append(matchCategories.CategoryList, categories[i])
		}
	}

	// If none found, return 404
	if len(matchCategories.CategoryList) == 0 {
		ResourceNotFound(w, r)
		return
	}

	// Encode the new category list
	bytes, err := json.Marshal(matchCategories)

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
