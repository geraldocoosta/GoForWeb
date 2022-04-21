package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
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

	r := mux.NewRouter() // criando um novo router

	// func() {
	// 	defer recoverPanic()
	// 	insertValue()
	// }()

	r.HandleFunc("/", homeHandler)          // registrando uma função para ser executada ao ter uma requisição no path /
	r.HandleFunc("/{id}/view", byIdHandler) // registrando uma função para ser executada ao ter uma requisição no path /

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) // registrando um prefixo para o path /static/ e passando o diretório de arquivos para o handler

	fmt.Println(http.ListenAndServe(":8080", r)) // subindo e servindo uma aplicação
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	items := func() []Post {
		defer recoverPanic()
		return listPosts()
	}()

	t := template.Must(template.ParseFiles("templates/body.html", "templates/index.html")) // registrando um template para ser servido
	if err := t.ExecuteTemplate(w, "index.html", items); err != nil {                      // verificando se houve erro na execução do template, passando o post pro template
		http.Error(w, err.Error(), http.StatusInternalServerError) // se houver erro, manda um status code 500
	}

}

func byIdHandler(w http.ResponseWriter, r *http.Request) {
	posts := []Post{}
	id := mux.Vars(r)["id"]                                                                // pegando o id do post
	post := getPostById(id)                                                                // pegando o post pelo id
	posts = append(posts, *post)                                                           // salvando o post no slice de posts
	t := template.Must(template.ParseFiles("templates/body.html", "templates/index.html")) // registrando um template para ser servido
	if err := t.ExecuteTemplate(w, "index.html", posts); err != nil {                      // verificando se houve erro na execução do template, passando o post pro template
		http.Error(w, err.Error(), http.StatusInternalServerError) // se houver erro, manda um status code 500
	}
}

func getPostById(id string) *Post {
	row := db.QueryRow("SELECT id, title, body FROM posts WHERE id = ?", id) // executando a query
	post := Post{}
	row.Scan(&post.Id, &post.Title, &post.Body)
	return &post
}

func listPosts() []Post {
	items := []Post{}
	rows, error := db.Query("SELECT id, title, body FROM posts") // executando a query
	if error != nil {
		log.Panicln(error)
	}

	for rows.Next() { // percorrendo os resultados da query
		post := Post{}

		error = rows.Scan(&post.Id, &post.Title, &post.Body) // pegando os resultados da query e salvando em variáveis
		if error != nil {
			log.Panicln(error)
		}
		items = append(items, post) // salvando os resultados da query em um slice de posts
	}
	return items
}

func insertValue() {
	stmt, error := db.Prepare("INSERT INTO posts(title, body) VALUES(?, ?)") // preparando a query para ser executada
	if error != nil {
		log.Panicln(error)
	}

	_, error = stmt.Exec("My second post", "This is my second post") // executando a query
	if error != nil {
		log.Panicln(error)
	}
}

func recoverPanic() {
	if err := recover(); err != nil {
		log.Println("panic occurred: ", err)
	}
}
