package webutils

import (
	"encoding/json"
	"log"
	"net/http"
)

// HandleError handles any http errors
func HandleError(w *http.ResponseWriter, code int, responseText string, logMessage string, err error) {
	errorMessage := ""
	writer := *w

	if err != nil {
		errorMessage = err.Error()
	}

	log.Println(logMessage, errorMessage)
	writer.WriteHeader(code)
	writer.Write([]byte(responseText))
}

// HandleSuccess handles valid http requests
func HandleSuccess(w *http.ResponseWriter, result interface{}) {
	writer := *w

	marshalled, err := json.Marshal(result)

	if err != nil {
		HandleError(w, 500, "Internal Server Error", "Error marshalling response JSON", err)
		return
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(200)
	writer.Write(marshalled)
}
