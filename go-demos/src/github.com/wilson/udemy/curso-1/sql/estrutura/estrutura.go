package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {

	res, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}

	return res
}

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(172.17.0.2:3306)/")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	exec(db, "CREATE DATABASE IF NOT EXISTS cursogolang")
	exec(db, "use cursogolang")
	exec(db, "drop table if exists usuarios")
	exec(db, `CREATE TABLE usuarios(
		id integer auto_increment,
		nome varchar(80),
		PRIMARY KEY (id)
		)`)
}
