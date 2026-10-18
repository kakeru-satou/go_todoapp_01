package handlers

import (
	"encoding/json"
	"errors"
	"go-learning/todoApp/auth"
	"go-learning/todoApp/middleware"
	"go-learning/todoApp/models"
	"go-learning/todoApp/response"
	"net/http"
)

type UserCreator interface {
	Signup(name, email, password string) (string, error)
}

type UserFinder interface {
	Signin(email, password string) (string, error)
	GetUserProfile(id int) (models.User, error)
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
		response.WriteErrMessage(w, http.StatusBadRequest, "invalid request")
		return
	}

	token, err := h.userCreator.Signup(req.Name, req.Email, req.Password)
	if err != nil {
		if errors.Is(err, models.ErrEmailDuplicate) {
			response.WriteErrMessage(w, http.StatusConflict, "すでに登録されているメールアドレスです")
			return
		}
		response.WriteErrMessage(w, http.StatusInternalServerError, "登録に失敗しました")
		return
	}

	res := models.SignupResponse{
		Message: "サインアップに成功",
		Token:   token,
	}

	err = response.WriteJSON(w, http.StatusCreated, res)
	if err != nil {
		response.WriteErrMessage(w, http.StatusInternalServerError, "failed to encode response")
	}
}

func (h *UserHandler) SigninHandler(w http.ResponseWriter, r *http.Request) {
	var req models.SigninRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteErrMessage(w, http.StatusBadRequest, "invalid request")
		return
	}

	token, err := h.userFinder.Signin(req.Email, req.Password)
	if err != nil {
		response.WriteErrMessage(w, http.StatusUnauthorized, "メールアドレス、またはパスワードが違います")
		return
	}

	res := models.SigninResponse{
		Message: "ログインに成功",
		Token:   token,
	}

	err = response.WriteJSON(w, http.StatusOK, res)
	if err != nil {
		response.WriteErrMessage(w, http.StatusInternalServerError, "failed to encode response")
	}
}

func (h *UserHandler) MeHandler(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(middleware.UserContextKey).(auth.AccessTokenClaims)

	user, err := h.userFinder.GetUserProfile(claims.UserID)
	if err != nil {
		response.WriteErrMessage(w, http.StatusNotFound, "failed to get user")
		return
	}

	err = response.WriteJSON(w, http.StatusOK, user)
	if err != nil {
		response.WriteErrMessage(w, http.StatusInternalServerError, "failed to encode response")
	}
}
