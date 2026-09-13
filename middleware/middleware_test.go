package middleware

import (
	"go-learning/todoApp/auth"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddleware_Success(t *testing.T) {
	next := &MockMiddleware{}

	token, err := auth.GenerateAccessToken(0, "test", "test")
	if err != nil {
		t.Fatalf("エラー:%s", err)
	}

	r := httptest.NewRequest(http.MethodGet, "/test", nil)
	r.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	han := Middleware(next)
	han.ServeHTTP(w, r)

	claims := next.r.Context().Value(UserContextKey).(auth.AccessTokenClaims)

	if next.CallCount != 1 {
		t.Errorf("mockが呼ばれていない:%v", next.CallCount)
	}
	if claims.UserID != 0 {
		t.Errorf("User IDが正しくない:%v", claims.UserID)
	}
}

func TestMiddleware_InvalidToken(t *testing.T) {
	next := &MockMiddleware{}

	r := httptest.NewRequest(http.MethodGet, "/test", nil)
	r.Header.Set("Authorization", "Bearer "+"invalid token")
	w := httptest.NewRecorder()

	han := Middleware(next)
	han.ServeHTTP(w, r)

	if next.CallCount != 0 {
		t.Errorf("nextが呼ばれている:%v", next.CallCount)
	}
	if w.Code != http.StatusUnauthorized {
		t.Errorf("取得ステータスが正しくない:%v", w.Code)
	}
}
