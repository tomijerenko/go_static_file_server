package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type endpoint struct {
	address   string
	port      string
	directory string
	socket    string
}

type responseRecorder struct {
	http.ResponseWriter
	status int
}

func Start(address string, port string, directory string) {
	e := &endpoint{
		address:   address,
		port:      port,
		directory: directory,
		socket:    fmt.Sprintf("%s:%s", address, port),
	}

	http.Handle("/", logger(http.FileServer(http.Dir(e.directory))))
	http.Handle("/health", logger(http.HandlerFunc(health)))
	fmt.Println(fmt.Sprintf("Serving from ./%s on %s", e.directory, e.socket))
	log.Fatal(http.ListenAndServe(e.socket, nil))
}

// Endpoints
func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

// Middleware
func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		recorder := &responseRecorder{ResponseWriter: w}
		next.ServeHTTP(recorder, r)
		duration := time.Since(startTime)
		log.Printf("[%s] %s %s %s %d %s",
			r.RemoteAddr,
			r.Method,
			r.RequestURI,
			r.Proto,
			recorder.status,
			duration,
		)
	})
}

func (r *responseRecorder) WriteHeader(statusCode int) {
	r.status = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func (r *responseRecorder) Write(data []byte) (int, error) {
	if r.status == 0 {
		r.status = http.StatusOK
	}
	return r.ResponseWriter.Write(data)
}
