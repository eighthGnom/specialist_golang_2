package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"semi-trash_api/models"
	"strconv"

	"github.com/gorilla/mux"
)

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	initHendlers(w)
	log.Println("Trying to get book by ID")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("Error with converting id to int", err)
		w.WriteHeader(400)
		msg := Message{Message: "Used ID is not support"}
		json.NewEncoder(w).Encode(msg)
		return
	}
	book, ok := models.FindBookById(id)
	if !ok {
		w.WriteHeader(404)
		msg := Message{Message: fmt.Sprintf("It is no book with ID %d", id)}
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(book)
}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	initHendlers(w)
	var book models.Book
	log.Println("Trying to get new book from client")
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Println("Error while decode request", err)
		w.WriteHeader(400)
		msg := Message{Message: "Provided json file is invalid"}
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(201)
	id := len(models.DB) + 2
	book.ID = id
	models.DB = append(models.DB, book)
	json.NewEncoder(w).Encode(book)

}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	initHendlers(w)
	log.Println("Trying to update book by ID")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("Error with converting id to int", err)
		w.WriteHeader(400)
		msg := Message{Message: "Used ID is not support"}
		json.NewEncoder(w).Encode(msg)
		return
	}
	var book models.Book
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Println("Error while decode request", err)
		w.WriteHeader(400)
		msg := Message{Message: "Provided json file is invalid"}
		json.NewEncoder(w).Encode(msg)
		return
	}
	newBook, ok := models.FindAndUpdateBookByID(id, book)
	if !ok {
		w.WriteHeader(404)
		msg := Message{Message: fmt.Sprintf("It is no book with ID %d", id)}
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(newBook)

}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	initHendlers(w)
	log.Println("Trying to delete book by ID")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("Error with converting id to int", err)
		w.WriteHeader(400)
		msg := Message{Message: "Used ID is not support"}
		json.NewEncoder(w).Encode(msg)
		return
	}
	book, ok := models.FindAndDeleteBookByID(id)
	if !ok {
		w.WriteHeader(404)
		msg := Message{Message: fmt.Sprintf("It is no book with ID %d", id)}
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(book)

}
