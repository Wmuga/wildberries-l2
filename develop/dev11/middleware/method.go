package middleware

import "net/http"

func SetMethod(method string, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(404)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

func SetMethodFunc(method string, f func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return SetMethod(method, http.HandlerFunc(f))
}
