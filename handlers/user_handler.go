package handlers

import (
	"encoding/json"
	"go-learning/todoApp/models"
	"net/http"
)

type UserCreator interface {
	Signup(name, email, password string) (models.User, error)
}

type UserHandler struct {
	userCreator UserCreator
}

func NewUserHandler(userCreator UserCreator) *UserHandler {
	return &UserHandler{
		userCreator: userCreator,
	}
}

func (h *UserHandler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	var req models.SignupRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	user, err := h.userCreator.Signup(req.Name, req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
