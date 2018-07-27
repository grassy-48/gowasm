package main

import (
	"flag"
	"log"
	"net/http"
)

type server struct {
	Handler http.Handler
}

func (s *server) handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/wasm")
	s.Handler.ServeHTTP(w, r)
}

func main() {
	addr := flag.String("addr", ":5555", "server address:port")
	flag.Parse()
	srv := http.FileServer(http.Dir("."))
	var s server
	s.Handler = srv
	log.Printf("listening on %q...", *addr)
	log.Fatal(http.ListenAndServe(*addr, s.handler()))
}
