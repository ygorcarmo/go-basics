package main

import (
	"example/museum/api"
	"example/museum/data"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func handleHello(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello from Go"))
}

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("templates/index.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	err = templ.Execute(w, data.GetAll())
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello", handleHello)
	server.HandleFunc("/template", handleTemplate)
	server.HandleFunc("/api/exhibitions", api.Get)
	server.HandleFunc("/api/exhibitions/new", api.Post)

	fs := http.FileServer(http.Dir("./public"))

	server.Handle("/", fs)

	err := http.ListenAndServe(":3000", server)
	if err != nil {
		fmt.Println("Error while openning the server")
	}
}
