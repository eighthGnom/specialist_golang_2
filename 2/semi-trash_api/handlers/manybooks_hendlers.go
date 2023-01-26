package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"semi-trash_api/models"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	initHendlers(w)
	log.Println("Trying to send to client all books")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(models.DB)
}
