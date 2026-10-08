package response

import (
	"encoding/json"
	"go-learning/todoApp/models"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteErrMessage(w http.ResponseWriter, status int, errMessage string) {
	message := models.ErrorMessageResponse{Message: errMessage}

	WriteJSON(w, status, message)
}