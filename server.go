package main

import (
	"fmt"
	"net/http"

	"./routes"
	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/", routes.Index).Methods("GET")

	port := ":8083"
	fmt.Println("start on port: " + port)
	http.ListenAndServe(port, mux)
}
