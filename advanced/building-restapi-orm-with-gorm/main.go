package main

import (
	db "building-restapi-orm-with-gorm/databases"
	model "building-restapi-orm-with-gorm/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var dbContext db.DbContext

func getUsers(w http.ResponseWriter, r *http.Request) {
	var users []model.User

	dbContext.Db.Find(&users)

	respondWithJSON(w, http.StatusOK, users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	dbContext.Db.Create(&user)

	respondWithJSON(w, http.StatusCreated, user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	var userEntity model.User
	dbContext.Db.Where("id= ?", user.Id).Find(&userEntity)
	userEntity.Name = user.Name
	userEntity.Email = user.Email

	dbContext.Db.Save(&userEntity)

	respondWithJSON(w, http.StatusCreated, user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	var user model.User
	dbContext.Db.Where("id= ?", id).Find(&user)
	dbContext.Db.Delete(&user)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	router.HandleFunc("/user", createUser).Methods("POST")
	router.HandleFunc("user", updateUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8081", router))
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func main() {
	dbContext.Initialize()

	handleRequests()
}
