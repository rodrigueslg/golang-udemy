package main

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func Sobre(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Sobre"))
}

func main() {
	http.HandleFunc("/home", Home)
	http.HandleFunc("/sobre", Sobre)
	http.ListenAndServe(":5000", nil)
}
