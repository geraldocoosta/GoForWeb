package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id    int
	Title string
	Body  string
}

var db, error = sql.Open("mysql", "root:root@/go-course?charset=utf8")

func main() {

	db.Ping() // verificando se houve erro na conexão com o banco de dados

	insertValue()

	http.HandleFunc("/", Handle)                   // registrando uma função para ser executada ao ter uma requisição no path /
	fmt.Println(http.ListenAndServe(":8080", nil)) // subindo e servindo uma aplicação
}

func insertValue() {
	stmt, error := db.Prepare("INSERT INTO posts(title, body) VALUES(?, ?)") // preparando a query para ser executada
	if error != nil {
		panic(error)
	}

	_, error = stmt.Exec("My second post", "This is my second post") // executando a query
	if error != nil {
		panic(error)
	}
}

// func Handle é adicionada como função para ser executada ao ter uma requisição no path /
func Handle(w http.ResponseWriter, r *http.Request) {
	post := Post{
		Id:    1,
		Title: "My first post",
		Body:  "This is my first post",
	}

	if r.FormValue("title") != "" {
		post.Title = r.FormValue("title")
	}

	t := template.Must(template.ParseFiles("templates/index.html"))  // registrando um template para ser servido
	if err := t.ExecuteTemplate(w, "index.html", post); err != nil { // verificando se houve erro na execução do template, passando o post pro template
		http.Error(w, err.Error(), http.StatusInternalServerError) // se houver erro, manda um status code 500
	}
}
