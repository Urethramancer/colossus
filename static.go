package main

import "net/http"

func static(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Moo"))
}
