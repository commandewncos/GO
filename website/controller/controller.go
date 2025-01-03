package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/Wilson-cmd/GO/connection"
	"github.com/Wilson-cmd/GO/model"
)

var (
	t = template.Must(template.ParseGlob("templates/*.html"))
)

func GetAllNames(w http.ResponseWriter, r *http.Request) {
	connection.DB.Find(&model.Names)
	t.ExecuteTemplate(w, "main", model.Names)

}

func InsertNames(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "create", nil)
}

func GetInfoByForm(w http.ResponseWriter, r *http.Request) {
	var names model.Name
	names.Name = r.FormValue("fname")
	value, err := strconv.Atoi(r.FormValue("fage"))
	if err != nil {
		fmt.Println("Erro ao converter para inteiro")
	}
	names.Age = value

	connection.DB.Create(&names)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
	t.ExecuteTemplate(w, "create", nil)

}

func DeleteById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	connection.DB.Delete(&model.Names, id)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)

}

func EditeById(w http.ResponseWriter, r *http.Request) {
	var name model.Name

	id := r.URL.Query().Get("id")
	connection.DB.First(&name, id)

	t.ExecuteTemplate(w, "edit", name)

}

func Update(w http.ResponseWriter, r *http.Request) {
	var name model.Name
	name.Name = r.FormValue("fname")
	value, err := strconv.Atoi(r.FormValue("fage"))
	checkError(err)

	name.Age = value

	id, e := strconv.Atoi(r.URL.Query().Get("id"))
	checkError(e)

	name.Id = id

	connection.DB.Save(&name)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

// Check error
func checkError(e error) {
	if e != nil {
		fmt.Println("Erro ao converter para inteiro")

	}
}
