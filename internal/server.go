package server

import (
	"fmt"
	"log"
	"net/http"
)

type endpoint struct {
	address   string
	port      string
	directory string
	socket    string
}

func Start(address string, port string, directory string) {
	e := &endpoint{
		address:   address,
		port:      port,
		directory: directory,
		socket:    fmt.Sprintf("%s:%s", address, port),
	}

	http.HandleFunc("/health", e.health)
	http.Handle("/", http.FileServer(http.Dir(e.directory)))
	fmt.Println(fmt.Sprintf("Serving from ./%s on %s", e.directory, e.socket))
	log.Fatal(http.ListenAndServe(e.socket, nil))
}

func (e *endpoint) health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}
