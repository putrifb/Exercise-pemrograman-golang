package main

import (
	"fmt"
	"net/http"
)

func AdminHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to Admin page"))
	}
}

func RequestMethodGetMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method is not allowed"))
			// http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	}) // TODO: replace this
}

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("role") != "ADMIN" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Role not authorized"))
			// http.Error(w, "Role not authorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}) // TODO: replace this
}

func main() {
	// TODO: answer here
	adminHandler := http.HandlerFunc(AdminHandler())

	http.Handle("/admin", RequestMethodGetMiddleware(AdminMiddleware(adminHandler)))

	fmt.Println("Starting server at port 8080")

	http.ListenAndServe("localhost:8080", nil)
}
