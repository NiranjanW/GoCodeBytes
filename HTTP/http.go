package main

import "net/http"

/**
To build your service:
	http.Handler
	http.HandlerFunc
http.Server
To handle requests:
	http.ResponseWriter
	http.Request
To build a client:
	http.Client
	http.Transport

**/
func main() {
	http.HandleFunc("/", publicHandler)
	http.HandleFunc("/admin", authenticate(adminHandler))
	http.ListenAndServe(":8080", nil)
}

func publicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("public"))
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("admin"))
}

func authenticate(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if ok && user == "john" && pass == "secret" {
			h.ServeHTTP(w, r)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
	}

}
