package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	databse, err := sql.Open("mysql", "root:root@tcp(172.17.0.2:3306)/cursogolang")
	if err != nil {
		log.Fatal(err)
	}

	defer databse.Close()

	tx, _ := databse.Begin()

	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?,?)")

	stmt.Exec(15, "Ana")
	stmt.Exec(16, "Sonia")
	_, err = stmt.Exec(1, "Wilson") //chave duplicada

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
