package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Post struct {
	Id    int
	Title string
	Body  string
}

func main() {
	http.HandleFunc("/", Handle)                   // registrando uma função para ser executada ao ter uma requisição no path /
	fmt.Println(http.ListenAndServe(":8080", nil)) // subindo e servindo uma aplicação
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
