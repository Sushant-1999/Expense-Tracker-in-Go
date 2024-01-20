package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	// _ "expense-tracker-api/cmd/expense-tracker-api/docs"
	_ "expense-tracker-api/internal/routes/docs"

	"expense-tracker-api/internal/routes"

	httpSwagger "github.com/swaggo/http-swagger"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	router := chi.NewRouter()

	// Middleware
	router.Use(middleware.Logger)
	router.Use(middleware.CleanPath)
	router.Use(middleware.Heartbeat("/ping"))
	router.Use(middleware.Recoverer)
	router.Use(middleware.RedirectSlashes)

	// Mount Routers
	router.Get("/", routes.ServeAPIDoc)
	router.Mount("/category", routes.CategoryRouter())
	router.Mount("/user", routes.UserRouter())
	router.Mount("/expense", routes.ExpenseRouter())
	// Swagger UI
	// router.Get("/swagger/*", httpSwagger.Handler(
	// 	httpSwagger.URL("docs/swagger.json"), // The URL to the generated Swagger JSON file
	// ))
	router.Mount("/swagger", httpSwagger.WrapHandler)

	// Handle 404
	router.NotFound(routes.ResourceNotFound)

	// Create Server at specified port
	port := 8080
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	// Start the server
	fmt.Printf("Server is starting on port %d...\n", port)
	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Failed to listen to server!", err)
	}
}
