package router

import (
	"fmt"
	"go-learning/todoApp/db"
	"go-learning/todoApp/handlers"
	"go-learning/todoApp/models"
	"go-learning/todoApp/repositories"
	"go-learning/todoApp/services"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func UserRouterTestSetup() *http.ServeMux {
	db.Init()

	rep := repositories.UserRepository{}
	ser := services.NewUserService(rep, rep)
	han := handlers.NewUserHandler(ser, ser)

	return SetupUserRoutes(han)
}

func SendUserRequest(mux *http.ServeMux, method, path, body string) *httptest.ResponseRecorder {
	reader := strings.NewReader(body)

	req := httptest.NewRequest(
		method,
		path,
		reader,
	)

	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	return w
}

func TestTableUserSignup(t *testing.T) {
	tests := []struct {
		name              string
		email             string
		method            string
		path              string
		status            int
		isCreateSameEmail bool
	}{
		{"RouterSignup_Success", "test@test.com", http.MethodPost, "/api/auth/signup", http.StatusCreated, false},
		{"RouterSignup_SameEmail", "test@test.com", http.MethodPost, "/api/auth/signup", http.StatusBadRequest, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mux := UserRouterTestSetup()
			body := fmt.Sprintf(`{"name":"test","email":"%s","password":"test"}`, tt.email)

			db.DB.Where("email = ?", tt.email).Delete(&models.User{})

			var w *httptest.ResponseRecorder

			if tt.isCreateSameEmail {
				w = SendUserRequest(mux, tt.method, tt.path, body)
				if w.Code != http.StatusCreated {
					t.Fatalf(
						"%s:想定ステータス: %d , 取得ステータス: %d",
						tt.name,
						http.StatusCreated,
						w.Code,
					)
				}
			}

			w = SendUserRequest(mux, tt.method, tt.path, body)
			if w.Code != tt.status {
				t.Errorf(
					"%s:想定ステータス: %d , 取得ステータス: %d",
					tt.name,
					tt.status,
					w.Code,
				)
			}
		})
	}
}

func TestTableUserSignin(t *testing.T) {
	tests := []struct {
		name           string
		createEmail    string
		createPassword string
		signinEmail    string
		signinPassword string
		method         string
		path           string
		status         int
	}{
		{"RouterSignin_Success", "test@test.com", "test", "test@test.com", "test", http.MethodPost, "/api/auth/signin", http.StatusOK},
		{"RouterSignin_InvalidEmail", "test@test.com", "test", "invalid@test.com", "test", http.MethodPost, "/api/auth/signin", http.StatusUnauthorized},
		{"RouterSignin_InvalidPassword", "test@test.com", "test", "test@test.com", "invalid", http.MethodPost, "/api/auth/signin", http.StatusUnauthorized},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mux := UserRouterTestSetup()
			body := fmt.Sprintf(`{"name":"test","email":"%s","password":"%s"}`, tt.createEmail, tt.createPassword)

			db.DB.Where("email = ?", tt.createEmail).Delete(&models.User{})

			var w *httptest.ResponseRecorder

			w = SendUserRequest(mux, http.MethodPost, "/api/auth/signup", body)
			if w.Code != http.StatusCreated {
				t.Fatalf(
					"%s:想定ステータス: %d , 取得ステータス: %d",
					tt.name,
					http.StatusCreated,
					w.Code,
				)
			}

			body = fmt.Sprintf(`{"email":"%s","password":"%s"}`, tt.signinEmail, tt.signinPassword)

			w = SendUserRequest(mux, tt.method, tt.path, body)
			if w.Code != tt.status {
				t.Errorf(
					"%s:想定ステータス: %d , 取得ステータス: %d",
					tt.name,
					tt.status,
					w.Code,
				)
			}
		})
	}
}
