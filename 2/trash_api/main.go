package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var (
	port = "8080"
	DB   []Pizza
)

func init() {
	p1 := Pizza{
		ID:    1,
		Name:  "Margarita",
		Size:  25,
		Price: 100.23,
	}
	p2 := Pizza{
		ID:    2,
		Name:  "BBQ",
		Size:  25,
		Price: 100.23,
	}
	p3 := Pizza{
		ID:    3,
		Name:  "Сheese",
		Size:  25,
		Price: 100.23,
	}
	DB = []Pizza{p1, p2, p3}
}

type Pizza struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Size  int     `json:"size"`
	Price float64 `json:"price"`
}

type ErrorMassage struct {
	Message string `json:"message"`
}

func isPizza(id int) (Pizza, bool) {
	for _, pizza := range DB {
		if pizza.ID == id {
			return pizza, true
		}
	}
	return Pizza{}, false
}

func GetAllPizzas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bytes, err := json.MarshalIndent(DB, "", "    ")
	if err != nil {
		w.WriteHeader(404)
		log.Println("Some problems with marshaling", err)
		return
	}
	log.Println("Get infos about all pizzas in database")
	w.WriteHeader(200)
	if _, err := w.Write(bytes); err != nil {
		w.WriteHeader(404)
		log.Println("Some problems with writing response", err)
	}
}

func GetPizzaByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ids := mux.Vars(r)
	id, err := strconv.Atoi(ids["id"])
	if err != nil {
		w.WriteHeader(400)
		log.Println("Problem with converting ID to string", err)
		msg := ErrorMassage{"Used ID not supported"}
		json.NewEncoder(w).Encode(msg)
		return
	}

	log.Println("Try to get pizza by ID")
	if pizza, ok := isPizza(id); ok {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(pizza)
		return
	}
	w.WriteHeader(404)
	msg := ErrorMassage{Message: fmt.Sprintf("It is no pizza with ID: %d", id)}
	json.NewEncoder(w).Encode(msg)

}

func main() {
	log.Println("Trying to start http sever")
	router := mux.NewRouter()
	router.HandleFunc("/pizzas", GetAllPizzas).Methods("GET")
	router.HandleFunc("/pizza/{id}", GetPizzaByID).Methods("GET")
	log.Println("Router configured successfully")
	log.Fatal(http.ListenAndServe("localhost:"+port, router))
}
