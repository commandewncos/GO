package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(172.17.0.2)/cursogolang")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Update
	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")
	stmt.Exec("Mario", 1)
	stmt.Exec("Jererê", 2)

	//Deletando o id 6 que estava duplicado
	stmt2, _ := db.Prepare("delete from usuarios where id = ?")
	stmt2.Exec(6)

}
