package controllers

import (
	"encoding/json"
	"fmt"
	"go/api/database"
	"go/api/models"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!")
}

func GetNames(w http.ResponseWriter, r *http.Request) {
	var names []models.Name
	database.DB.Find(&names)
	json.NewEncoder(w).Encode(names)

}

func GetById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	var name models.Name

	database.DB.First(&name, id)
	json.NewEncoder(w).Encode(name)

}

func CreateName(w http.ResponseWriter, r *http.Request) {
	var name models.Name

	json.NewDecoder(r.Body).Decode(&name)

	database.DB.Create(&name)
	json.NewEncoder(w).Encode(name)
}

func DeleteName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var name models.Name
	database.DB.Delete(&name, id)
	json.NewEncoder(w).Encode(name)

}

func EditName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var name models.Name
	database.DB.First(&name, id)

	json.NewDecoder(r.Body).Decode(&name)
	database.DB.Save(&name)

	json.NewEncoder(w).Encode(name)

}
