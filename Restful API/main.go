package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	x1 := &messageHandler{"Onur"}
	x2 := &messageHandler{"Teber"}

	mux.Handle("/bir", x1)
	mux.Handle("/iki", x2)

	log.Println("Listening...")

	http.ListenAndServe(":8080", mux)

	/*
		var i ironman
		var w wolverine

		mux := http.NewServeMux()

		mux.Handle("/ironman", i)
		mux.Handle("/wolverine", w)
	*/

	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Merhaba"))
		})
	*/

}

type messageHandler struct {
	message string
}

func (x *messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, x.message)
}

/*
	type ironman int

	func (x ironman) ServeHTTP(res http.ResponseWriter, r *http.Request) {
		io.WriteString(res, "Mr. Iron!")
	}

	type wolverine int

	func (x wolverine) ServeHTTP(res http.ResponseWriter, r *http.Request) {
		io.WriteString(res, "Mr. Wolverine!")
	}
*/

/*
	func indexHandler(w http.ResponseWriter, r *http.Request) {
		//w.Write([]byte("Index"))
		x := r.URL.Path[1:]
		data := "Merhaba " + x
		w.Write([]byte(data))
	}

	func aboutHandler(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("About"))
	}
*/
