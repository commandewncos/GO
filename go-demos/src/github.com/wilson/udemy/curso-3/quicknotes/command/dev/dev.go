package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

var (
	conn *pgx.Conn
	err  error
)

type Post struct {
	id                     int
	title, content, author string
}

func main() {

	uniformResourceLocator := "postgres://postgres:secret@localhost:5432/postgres"

	conn, err = pgx.Connect(context.Background(), uniformResourceLocator)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection with POSTGRES OK!")
	defer conn.Close(context.Background())

	createTable()

	// insertPostWithReturn("Post 5", "Content of post 5", "Ana Bernardo")

	selectById(50)
	selectAllPosts()
}

func createTable() {
	query := `
		CREATE TABLE IF NOT EXISTS posts(
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			content TEXT,
			author TEXT NOT NULL
		);
	`

	_, e := conn.Exec(context.Background(), query)
	if e != nil {
		panic(e)
	}

	fmt.Println("Table posts created")
}

// func insertPostWithoutReturn(title, content, author string) {

// 	//PLACEHOLDER
// 	query := `
// 		INSERT INTO posts(title, content, author)
// 		VALUES($1, $2, $3)
// 	`
// 	_, e := conn.Exec(context.Background(), query, title, content, author)

// 	if e != nil {
// 		panic(e)
// 	} else {
// 		fmt.Println("Post created!")
// 	}
// }

// func insertPostWithReturn(title, content, author string) {
// 	var id int
// 	query := `
// 			INSERT INTO posts(title, content, author)
// 			VALUES($1, $2, $3) RETURNING id;
// 		`
// 	row := conn.QueryRow(context.Background(), query, title, content, author)
// 	e := row.Scan(&id)
// 	if e != nil {
// 		panic(e)
// 	} else {
// 		fmt.Printf("Post created id:\t%d\n", id)
// 	}
// }

func selectById(id int) {
	var title, content, author string
	query := `SELECT title, content, author FROM posts WHERE id = $1`
	row := conn.QueryRow(context.Background(), query, id)

	e := row.Scan(&title, &content, &author)

	switch {
	case e == pgx.ErrNoRows:
		fmt.Println("Id not found", id)
		return
	case e != nil:
		panic(e)
	default:
		fmt.Printf("Post:\ntitle=%s\ncontent=%s\nauthor=%s\n", title, content, author)

	}

}

func selectAllPosts() {
	r, e := conn.Query(context.Background(), `SELECT * FROM posts`)

	if e != nil {
		panic(e)
	}

	defer r.Close()
	if r.Err() != nil {
		panic(r.Err())
	}
	var posts []Post
	for r.Next() {
		var post Post
		e = r.Scan(&post.id, &post.title, &post.content, &post.author)
		if e != nil {
			panic(e)
		}
		posts = append(posts, post)

	}

	for _, post := range posts {
		fmt.Println(post)
	}

}
