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
	defer db.Close() // fechando a conexão com o banco de dados
	db.Ping()        // verificando se houve erro na conexão com o banco de dados

	// insertValue()
	items := getValues()

	http.HandleFunc("/", returnHandle(items))      // registrando uma função para ser executada ao ter uma requisição no path /
	fmt.Println(http.ListenAndServe(":8080", nil)) // subindo e servindo uma aplicação
}

func getValues() []Post {
	items := []Post{}
	rows, error := db.Query("SELECT id, title, body FROM posts") // executando a query
	if error != nil {
		panic(error)
	}

	for rows.Next() { // percorrendo os resultados da query
		post := Post{}

		error = rows.Scan(&post.Id, &post.Title, &post.Body) // pegando os resultados da query e salvando em variáveis
		if error != nil {
			panic(error)
		}
		items = append(items, post) // salvando os resultados da query em um slice de posts
	}
	return items
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

func returnHandle(posts []Post) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("templates/index.html"))   // registrando um template para ser servido
		if err := t.ExecuteTemplate(w, "index.html", posts); err != nil { // verificando se houve erro na execução do template, passando o post pro template
			http.Error(w, err.Error(), http.StatusInternalServerError) // se houver erro, manda um status code 500
		}
	}
}
