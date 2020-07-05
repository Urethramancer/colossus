package mid

import "net/http"

// AddJSONHeaders for JSON responses.
func AddJSONHeaders(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// AddHTMLHeaders for web pages.
func AddHTMLHeaders(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// Preflight returns options for REST calls.
func Preflight(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST,GET,OPTIONS")
	w.Header().Set("Access-Control-Max-Age", "86400")
	http.Error(w, "", 204)
}

// AddCORS to allow REST access from other domains.
func AddCORS(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// AddSecureHeaders for SSL/TLS requests.
func AddSecureHeaders(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Strict-Transport-Security", "max-age=1000; includeSubDomains; preload")
		w.Header().Set("Content-Security-Policy", "upgrade-insecure-requests")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
