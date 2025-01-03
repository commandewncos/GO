package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Usuario
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// UsuarioHandler analisa o request e delegar para a função adequada
func UsuarioHandler(w http.ResponseWriter, r *http.Request) {

	//Ele vai eleminar tudo que estiver antes da ultima barra
	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		usuariosPorID(w, id)

	case r.Method == "GET":
		usuarioTodos(w)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `Ops...`)
	}

}

func usuariosPorID(w http.ResponseWriter, id int) {
	db, err := sql.Open("mysql", "root:root@tcp(172.17.0.2)/cursogolang")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var usuario User

	//Query row retorna uma unica linha
	db.QueryRow("SELECT id, nome FROM usuarios WHERE id = ?", id).Scan(&usuario.ID, &usuario.Name)

	json, _ := json.Marshal(usuario)

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func usuarioTodos(w http.ResponseWriter) {
	db, err := sql.Open("mysql", "root:root@tcp(172.17.0.2)/cursogolang")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("SELECT id, nome FROM usuarios")
	defer rows.Close()

	var usuarios []User
	for rows.Next() {
		var usuario User

		rows.Scan(&usuario.ID, &usuario.Name)
		usuarios = append(usuarios, usuario)
	}

	json, _ := json.Marshal(usuarios)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)

}
