package main

import (
	"io"
  "regexp"
  "io/ioutil"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}
func ReadHTML() (string){
  contents, _ := ioutil.ReadFile("public/index.html")

  //println(string(contents))
  return string(contents)
}
var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {

	server := http.Server{
		Addr:    ":8000",
		Handler: &myHandler{},
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/"] = hello

	server.ListenAndServe()
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}
  re := regexp.MustCompile("\\{\\w*\\}")
  replace := re.ReplaceAllLiteralString(ReadHTML() , r.URL.String())

	io.WriteString(w, replace)
}
