package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/search", search)
	http.ListenAndServe(":8080", nil)
}

func search(w http.ResponseWriter, r *http.Request) {
	hl := r.FormValue("hl")
	source := r.FormValue("source")
	ei := r.FormValue("ei")
	q := r.FormValue("q")

	data := "/search?hl=" + hl + "&source=" + source + "&ei=" + ei + "&q=" + q

	w.Write([]byte(data))
}
