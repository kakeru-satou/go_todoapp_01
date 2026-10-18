package middleware

import "net/http"

type MockMiddleware struct {
	CallCount int
	w         http.ResponseWriter
	r         *http.Request
}

func (m *MockMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.CallCount++
	m.w = w
	m.r = r
}
