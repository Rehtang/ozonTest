package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	router  *mux.Router
	handler *Handler
}

func NewServer(handler *Handler) *Server {
	router := mux.NewRouter()
	server := &Server{
		router:  router,
		handler: handler,
	}

	router.HandleFunc("/shorten", server.handler.CreateShortURL).Methods(http.MethodPost)
	router.HandleFunc("/shorten/{shortURL}", server.handler.GetOriginalURL).Methods(http.MethodGet)

	return server
}

func (s *Server) Start(port string) {
	log.Printf("Server started on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, s.router))
}
