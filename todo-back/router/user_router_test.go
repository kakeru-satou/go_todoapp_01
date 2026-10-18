package router

import (
	"encoding/json"
	"fmt"
	"go-learning/todoApp/auth"
	"go-learning/todoApp/db"
	"go-learning/todoApp/handlers"
	"go-learning/todoApp/middleware"
	"go-learning/todoApp/models"
	"go-learning/todoApp/repositories"
	"go-learning/todoApp/services"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func UserRouterTestSetup(isWrap bool) (http.Handler, repositories.UserRepository) {
	db.Init()

	rep := repositories.UserRepository{}
	ser := services.NewUserService(rep, rep)
	han := handlers.NewUserHandler(ser, ser)

	mux := SetupUserRoutes(han)
	if isWrap {
		protectedMux := middleware.Middleware(mux)
		return protectedMux, rep
	}
	return mux, rep
}

func SendUserRequest(mux http.Handler, method, path, body string, token ...string) *httptest.ResponseRecorder {
	reader := strings.NewReader(body)

	req := httptest.NewRequest(
		method,
		path,
		reader,
	)

	w := httptest.NewRecorder()

	if len(token) > 0 {
		req.Header.Set("Authorization", "Bearer "+token[0])
	}

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
		isSuccess         bool
	}{
		{"RouterSignup_Success", "test@test.com", http.MethodPost, "/api/auth/signup", http.StatusCreated, false, true},
		{"RouterSignup_SameEmail", "test@test.com", http.MethodPost, "/api/auth/signup", http.StatusConflict, true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mux, _ := UserRouterTestSetup(false)
			meMux, _ := UserRouterTestSetup(true)
			body := fmt.Sprintf(`{"name":"test","email":"%s","password":"test"}`, tt.email)
			response := models.SignupResponse{}
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
			if tt.isSuccess {
				err := json.Unmarshal(w.Body.Bytes(), &response)
				if err != nil {
					t.Fatalf("エラー: %v", err)
				}
				w = SendUserRequest(meMux, http.MethodGet, "/api/users/me", "", response.Token)
				if w.Code != http.StatusOK {
					t.Errorf(
						"%s:想定ステータス: %d , 取得ステータス: %d",
						tt.name,
						http.StatusOK,
						w.Code,
					)
				}

				var responseUser models.User
				err = json.Unmarshal(w.Body.Bytes(), &responseUser)
				if err != nil {
					t.Fatalf("response body: %v", err)
				}
				if responseUser.Email != tt.email {
					t.Errorf("responseEmail:%s, email:%s", responseUser.Email, tt.email)
				}
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
		isSuccess      bool
	}{
		{"RouterSignin_Success", "test@test.com", "test", "test@test.com", "test", http.MethodPost, "/api/auth/signin", http.StatusOK, true},
		{"RouterSignin_InvalidEmail", "test@test.com", "test", "invalid@test.com", "test", http.MethodPost, "/api/auth/signin", http.StatusUnauthorized, false},
		{"RouterSignin_InvalidPassword", "test@test.com", "test", "test@test.com", "invalid", http.MethodPost, "/api/auth/signin", http.StatusUnauthorized, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mux, _ := UserRouterTestSetup(false)
			meMux, _ := UserRouterTestSetup(true)
			body := fmt.Sprintf(`{"name":"test","email":"%s","password":"%s"}`, tt.createEmail, tt.createPassword)
			response := models.SigninResponse{}
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
			if tt.isSuccess {
				err := json.Unmarshal(w.Body.Bytes(), &response)
				if err != nil {
					t.Fatalf("エラー: %v", err)
				}
				w = SendUserRequest(meMux, http.MethodGet, "/api/users/me", "", response.Token)
				if w.Code != http.StatusOK {
					t.Errorf(
						"%s:想定ステータス: %d , 取得ステータス: %d",
						tt.name,
						http.StatusOK,
						w.Code,
					)
				}

				var responseUser models.User
				err = json.Unmarshal(w.Body.Bytes(), &responseUser)
				if err != nil {
					t.Fatalf("response body: %v", err)
				}
				if responseUser.Email != tt.signinEmail {
					t.Errorf("responseEmail:%s, email:%s", responseUser.Email, tt.signinEmail)
				}
			}
		})
	}
}

func TestTableUserMe(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		isToken bool
		status  int
	}{
		{"RouterMe_Success", "/api/users/me", true, http.StatusOK},
		{"RouterMe_InvalidToken", "/api/users/me", false, http.StatusUnauthorized},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testMux, rep := UserRouterTestSetup(false)
			mux, _ := UserRouterTestSetup(true)
			email := "test@test.com"
			body := fmt.Sprintf(`{"name":"test","email":"%s","password":"teset"}`, email)
			token := ""

			db.DB.Where("email = ?", email).Delete(&models.User{})

			var w *httptest.ResponseRecorder

			w = SendUserRequest(testMux, http.MethodPost, "/api/auth/signup", body)
			if w.Code != http.StatusCreated {
				t.Fatalf(
					"%s:想定ステータス: %d , 取得ステータス: %d",
					tt.name,
					http.StatusCreated,
					w.Code,
				)
			}

			user, err := rep.GetByEmail(email)
			if err != nil {
				t.Fatalf("エラー:%v", err)
			}

			body = ""
			if tt.isToken {
				token, err = auth.GenerateAccessToken(user.ID, "test", "test")
				if err != nil {
					t.Fatalf("エラー:%v", err)
				}
			}

			w = SendUserRequest(mux, http.MethodGet, tt.path, body, token)
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
