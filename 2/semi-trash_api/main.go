package main

import (
	"log"
	"net/http"
	"os"
	"semi-trash_api/utils"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	env := godotenv.Load()
	if env != nil {
		log.Fatal("Can't find .env files", env)
	}
	port = os.Getenv("app_port")
}

const (
	prefix = "/api/v1"
)

var (
	port                    string
	bookResourcePrefix      = prefix + "/book"
	manyBooksResourcePrefix = prefix + "/books"
)

func main() {
	log.Println("Trying to start initialise server")
	router := mux.NewRouter()
	utils.InitBookResource(router, bookResourcePrefix)
	utils.InitManyBooksResourse(router, manyBooksResourcePrefix)
	log.Println("Server is ready, lets go!")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
