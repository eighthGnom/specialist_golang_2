package handlers

import "net/http"

func initHendlers(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

type Message struct {
	Message string
}
