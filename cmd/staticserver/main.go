package main

import (
	"log"
	"net/http"
)

type Middleware struct {
	handler http.Handler
}

func (m *Middleware) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Cache-Control", "public, max-age=30")
	m.handler.ServeHTTP(w, req)
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", &Middleware{handler: fs})

	log.Print("Listening on :80...")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
