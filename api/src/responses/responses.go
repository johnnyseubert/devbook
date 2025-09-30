package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorReponse struct {
	Message string `json:"message"`
}

func Json(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

func Error(w http.ResponseWriter, statusCode int, err error) {
	Json(w, statusCode, errorReponse{
		Message: err.Error(),
	})
}
