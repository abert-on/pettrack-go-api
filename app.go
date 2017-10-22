package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"gopkg.in/mgo.v2/bson"

	C "github.com/abert-on/pettrack-go-api/config"
	D "github.com/abert-on/pettrack-go-api/dao"
	M "github.com/abert-on/pettrack-go-api/models"
	"github.com/gorilla/mux"
)

var config = C.Config{}
var dao = D.PetsDAO{}

/*
AllPetsEndpoint fetches all Pet objects from DB
*/
func AllPetsEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pets, err := dao.FindAllByUserID(params["userId"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, pets)
}

/*
FindPetEndpoint fetches a Pet object with matching ID from DB
*/
func FindPetEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pet, err := dao.FindByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Pet ID")
		return
	}
	respondWithJSON(w, http.StatusOK, pet)
}

/*
CreatePetEndpoint creates a Pet entry in DB
*/
func CreatePetEndpoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))
	var pet M.Pet
	if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	pet.ID = bson.NewObjectId()
	if err := dao.Insert(pet); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	respondWithJSON(w, http.StatusCreated, pet)
}

/*
UpdatePetEndpoint updates an existng Pet entry in DB
*/
func UpdatePetEndpoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var pet M.Pet
	if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(pet); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

/*
DeletePetEndpoint deletes a pet entry from DB
*/
func DeletePetEndpoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var pet M.Pet
	if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(pet); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the config file 'config.toml' and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

// Define HTTP request routes
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
