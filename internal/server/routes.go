package server

import (
	"encoding/json"
	"log"
	"net/http"

	"agendahora/cmd/web"

	"github.com/a-h/templ"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /hello-world", s.HelloWorldHandler)
	mux.HandleFunc("GET /health", s.healthHandler)

	fileServer := http.FileServer(http.FS(web.Files))

	mux.Handle("GET /assets/", fileServer)
	mux.Handle("GET /web", templ.Handler(web.HelloForm()))
	mux.HandleFunc("POST /hello", web.HelloWebHandler)

	return mux
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
