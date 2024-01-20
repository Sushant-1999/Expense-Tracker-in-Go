package routes

import (
	"net/http"

	"expense-tracker-api/internal/utils"
)

// Serves the API documentation as HTML (default) or txt file
func ServeAPIDoc(w http.ResponseWriter, r *http.Request) {
	// check the requested type and send appropriate file
	if r.Header.Get("Accept") == "text/plain" {
		w.Header().Set("Content-Type", "text/plain")
		w.Write(utils.GetDocTXT())
	} else {
		w.Header().Set("Content-Type", "text/html")
		w.Write(utils.GetDocHTML())
	}
}

// Custom ResourceNotFound - 404 Response
func ResourceNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)                                // set 404 status code
	w.Write([]byte("The resource you're looking for doesn't exist.")) // send message
}
