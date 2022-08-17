package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	addr := ":" + port

	fmt.Printf("Starting web server, listening on %s\n", addr)

	server := &Server{}

	err := http.ListenAndServe(addr, server)
	if err != nil {
		panic(err)
	}
}

type Server struct {
}

type Note struct {
	Data     []byte
	Destruct bool
}

func (s *Server) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" || r.Method == "HEAD" {
		noteID := strings.TrimPrefix(r.URL.Path, "/")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(
			fmt.Sprintf(
				"You requested the note with the ID '%s'.",
				noteID)))
		return
	}

	if r.Method == "POST" && r.URL.Path == "/" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("You posted to /."))
		return
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not Found"))

}

func (s *Server) notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not Found"))
}

func (s *Server) handlePOST(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("You posted to /."))
}

func (s *Server) handleGET(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/" {
		return
	}

	noteID := strings.TrimPrefix(path, "/")
	ctx := r.Context()
	note := &Note{}

	w.WriteHeader(http.StatusOK)
	w.Write(note.Data)
}
