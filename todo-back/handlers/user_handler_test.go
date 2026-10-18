package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"go-learning/todoApp/auth"
	"go-learning/todoApp/db"
	"go-learning/todoApp/middleware"
	"go-learning/todoApp/models"
	"go-learning/todoApp/repositories"
	"go-learning/todoApp/services"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func UserHandlerTestSetup() (*UserHandler, repositories.UserRepository) {
	db.Init()
	rep := repositories.UserRepository{}
	ser := services.NewUserService(rep, rep)
	han := NewUserHandler(ser, ser)

	return han, rep
}

func SendUserRequest(method, path, body string, handlerMethod func(http.ResponseWriter, *http.Request)) *httptest.ResponseRecorder {
	reader := strings.NewReader(body)

	r := httptest.NewRequest(
		method,
		path,
		reader,
	)
	w := httptest.NewRecorder()
	handlerMethod(w, r)
	return w
}

func TestTableUserHandlerSignup(t *testing.T) {
	tests := []struct {
		name              string
		email             string
		status            int
		isCreateSameEmail bool
	}{
		{"HandlerSignup_Success", "test@test.com", http.StatusCreated, false},
		{"HandlerSignup_SameEmail", "test@test.com", http.StatusConflict, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			han, _ := UserHandlerTestSetup()
			method := http.MethodPost
			body := fmt.Sprintf(`{"name":"test","email":"%s","password":"test"}`, tt.email)
			db.DB.Where("email = ?", tt.email).Delete(&models.User{})

			if tt.isCreateSameEmail {
				w := SendUserRequest(method, "/api/auth/signup", body, han.SignupHandler)
				if w.Code != http.StatusCreated {
					t.Fatalf("ステータス:%v", w.Code)
				}
			}

			w := SendUserRequest(method, "/api/auth/signup", body, han.SignupHandler)
			if w.Code != tt.status {
				t.Errorf(
					"想定ステータス: %d , 取得したステータス: %d",
					tt.status,
					w.Code,
				)
			}

			if tt.status == http.StatusCreated {
				var responseSignup models.SignupResponse
				err := json.Unmarshal(w.Body.Bytes(), &responseSignup)
				if err != nil {
					t.Fatalf("response body: %v", err)
				}
				if responseSignup.Token == "" {
					t.Errorf("トークンが空")
				}
			}
		})
	}
}

func TestTableUserHandlerSignin(t *testing.T) {
	tests := []struct {
		name           string
		createEmail    string
		createPassword string
		signinEmail    string
		signinPassword string
		status         int
	}{
		{"HandlerSignin_Success", "test@test.com", "test", "test@test.com", "test", http.StatusOK},
		{"HandlerSignin_InvalidEmail", "test@test.com", "test", "invalid@test.com", "test", http.StatusUnauthorized},
		{"HandlerSignin_InvalidPassword", "test@test.com", "test", "test@test.com", "invalid", http.StatusUnauthorized},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			han, _ := UserHandlerTestSetup()
			method := http.MethodPost
			body := fmt.Sprintf(`{"name":"test","email":"%s","password":"%s"}`, tt.createEmail, tt.createPassword)
			db.DB.Where("email = ?", tt.createEmail).Delete(&models.User{})

			w := SendUserRequest(method, "/api/auth/signup", body, han.SignupHandler)
			if w.Code != http.StatusCreated {
				t.Fatalf(
					"想定ステータス: %d , 取得したステータス: %d",
					http.StatusCreated,
					w.Code,
				)
			}

			body = fmt.Sprintf(`{"email":"%s","password":"%s"}`, tt.signinEmail, tt.signinPassword)
			w = SendUserRequest(method, "/api/auth/signin", body, han.SigninHandler)
			if w.Code != tt.status {
				t.Errorf(
					"想定ステータス: %d , 取得したステータス: %d",
					tt.status,
					w.Code,
				)
			}

			if tt.status == http.StatusOK {
				var responseSignin models.SigninResponse
				err := json.Unmarshal(w.Body.Bytes(), &responseSignin)
				if err != nil {
					t.Fatalf("response body: %v", err)
				}
				if responseSignin.Token == "" {
					t.Errorf("トークンが空")
				}
			}
		})
	}
}

func TestTableUserHandlerMe(t *testing.T) {
	tests := []struct {
		name         string
		status       int
		isDeleteUser bool
	}{
		{"HandlerMe_Success", http.StatusOK, false},
		{"HandlerMe_InvalidUserID", http.StatusNotFound, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			han, rep := UserHandlerTestSetup()
			email := "test@test.com"
			method := http.MethodPost
			body := fmt.Sprintf(`{"name":"test","email":"%s","password":"test"}`, email)
			db.DB.Where("email = ?", email).Delete(&models.User{})

			w := SendUserRequest(method, "/api/auth/signup", body, han.SignupHandler)
			if w.Code != http.StatusCreated {
				t.Fatalf(
					"想定ステータス: %d , 取得したステータス: %d",
					http.StatusCreated,
					w.Code,
				)
			}

			user, err := rep.GetByEmail(email)
			if err != nil {
				t.Fatalf("エラー:%v", err)
			}
			if tt.isDeleteUser {
				db.DB.Where("email = ?", email).Delete(&models.User{})
			}

			req := httptest.NewRequest(
				http.MethodGet,
				"/api/users/me",
				nil,
			)
			w = httptest.NewRecorder()
			ctx := context.WithValue(req.Context(), middleware.UserContextKey, auth.AccessTokenClaims{UserID: user.ID})
			req = req.WithContext(ctx)

			han.MeHandler(w, req)
			if w.Code != tt.status {
				t.Errorf(
					"想定ステータス: %d , 取得したステータス: %d",
					tt.status,
					w.Code,
				)
			}

			if tt.status == http.StatusOK {
				var responseUser models.User
				err := json.Unmarshal(w.Body.Bytes(), &responseUser)
				if err != nil {
					t.Fatalf("response body: %v", err)
				}
				if responseUser.Email != email {
					t.Errorf("responseEmail:%s, email:%s", responseUser.Email, email)
				}
			}
		})
	}
}
