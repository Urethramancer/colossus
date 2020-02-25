package main

import (
	"net/http"
)

func addJSONHeaders(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func addHTMLHeaders(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func addCORS(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func addSecureHeaders(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Strict-Transport-Security", "max-age=1000; includeSubDomains; preload")
		w.Header().Set("Content-Security-Policy", "upgrade-insecure-requests")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func (ws *WebServer) addLogger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		go func() {
			ws.Logger.TMsg("client %s %s %s", r.RemoteAddr, r.Method, r.RequestURI)
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
