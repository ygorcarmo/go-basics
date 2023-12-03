package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func handleHello(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello from Go"))
}

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("templates/index.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal err"))
	}
	html.Execute(w, "Test")
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello", handleHello)
	server.HandleFunc("/template", handleTemplate)

	fs := http.FileServer(http.Dir("./public"))

	server.Handle("/", fs)

	err := http.ListenAndServe(":3000", server)
	if err != nil {
		fmt.Println("Error while openning the server")
	}
}
