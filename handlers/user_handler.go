package handlers

import (
	"encoding/json"
	"go-learning/todoApp/models"
	"net/http"
)

type UserCreator interface {
	Signup(name, email, password string) (models.User, error)
}

type UserFinder interface {
	Signin(email, password string) (string, error)
}

type UserHandler struct {
	userCreator UserCreator
	userFinder  UserFinder
}

func NewUserHandler(userCreator UserCreator, userFinder UserFinder) *UserHandler {
	return &UserHandler{
		userCreator: userCreator,
		userFinder:  userFinder,
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

func (h *UserHandler) SigninHandler(w http.ResponseWriter, r *http.Request) {

	var req models.SigninRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	token, err := h.userFinder.Signin(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	res := models.SigninResponse{
		Message: "ログインに成功",
		Token:   token,
	}

	err = writeJSON(w, http.StatusOK, res)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
