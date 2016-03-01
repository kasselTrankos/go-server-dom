package main

import (
	"log"
	"fmt"
	"net/http"
  "html/template"
	"./com/bd"
)


type Context struct {
    Title  string
		Script template.HTML
    Static string
}
const STATIC_URL string = "/public/"


func main() {
	bd.init()
	http.HandleFunc("/", Home)
	http.HandleFunc("/public/", Home)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func Home(w http.ResponseWriter, req *http.Request) {
	context := Context{Title: "About", Script: template.HTML(`<script>alert('hola')</script>`)}

	 render(w, "index", context)
}

func render(w http.ResponseWriter, tmpl string, context Context) {
    context.Static = STATIC_URL
    tmpl_list := []string{fmt.Sprintf("public/%s.html", tmpl)}
    t, err := template.ParseFiles(tmpl_list...)
    if err != nil {
        log.Print("template parsing error: ", err)
    }
    err = t.Execute(w, context)
    if err != nil {
        log.Print("template executing error: ", err)
    }
}
