package main

import (
	db "building-restapi-orm-with-gorm/databases"
	model "building-restapi-orm-with-gorm/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var dbContext model.DbContext

func getUsers(w http.ResponseWriter, r *http.Request) {
}

func createUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var body [string]interface
	

}

func updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	router.HandleFunc("/user", createUser).Methods("POST")
	router.HandleFunc("user", updateUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	db.Initialize(&dbContext)

	handleRequests()
}
