package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/", index).Methods("GET")

	port := ":8083"
	fmt.Println("start on port: " + port)
	http.ListenAndServe(port, mux)
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello World from https://howto.prepareexam.icu/"))
}
