package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	. "github.com/abert-on/pettrack-go-api/models"
	"github.com/gorilla/mux"
)

/*
AllPetsEndpoint fetches all Pet objects from DB
*/
func AllPetsEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not implemented yet")
}

/*
FindPetEndpoint fetches a Pet object with matching ID from DB
*/
func FindPetEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not implemented yet")
}

/*
CreatePetEndpoint creates a Pet entry in DB
*/
func CreatePetEndpoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var pet Pet
	if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	pet.ID = bson.NewObjectId()
	if err := dao.Insert(pet); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	respondWithJson(w, http.StatusCreated, pet)
}

/*
UpdatePetEndpoint updates an exisitng Pet entry in DB
*/
func UpdatePetEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not implemented yet")
}

/*
DeletePetEndpoint deletes a pet entry from DB
*/
func DeletePetEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not implemented yet")
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w htpp.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/pets", AllPetsEndpoint).Methods("GET")
	r.HandleFunc("/pets", CreatePetEndpoint).Methods("POST")
	r.HandleFunc("/pets", UpdatePetEndpoint).Methods("PUT")
	r.HandleFunc("/pets", DeletePetEndpoint).Methods("DELETE")
	r.HandleFunc("/pets/{id}", FindPetEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
