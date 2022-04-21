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
	t := template.Must(template.ParseFiles("templates/index.html")) // registrando um template para ser servido
	if err := t.ExecuteTemplate(w, "index.html", nil); err != nil { // verificando se houve erro na execução do template
		http.Error(w, err.Error(), http.StatusInternalServerError) // se houver erro, manda um status code 500
	}
}
