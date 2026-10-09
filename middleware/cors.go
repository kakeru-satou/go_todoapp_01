package middleware

import "net/http"

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "content-type,authorization")
		w.Header().Set("Access-Control-Allow-Methods", "POST,DELETE,PATCH,GET")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}