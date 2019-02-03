package routes

import "net/http"

func Index(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello World from https://howto.prepareexam.icu/"))
}
