package utils

import (
	"semi-trash_api/handlers"

	"github.com/gorilla/mux"
)

func InitBookResource(router *mux.Router, prefix string) {
	router.HandleFunc(prefix+"/{id}", handlers.GetBookByID).Methods("GET")
	router.HandleFunc(prefix, handlers.CreateBook).Methods("POST")
	router.HandleFunc(prefix+"/{id}", handlers.UpdateBook).Methods("PUT")
	router.HandleFunc(prefix+"/{id}", handlers.DeleteBook).Methods("DELETE")
}

func InitManyBooksResourse(router *mux.Router, prefix string) {
	router.HandleFunc(prefix, handlers.GetAllBooks).Methods("GET")
}
