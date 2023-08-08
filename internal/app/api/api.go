package api

import (
	"net/http"

	gethandler "github.com/Vdolganov/shortify/internal/app/handlers/get_handler"
	posthandler "github.com/Vdolganov/shortify/internal/app/handlers/post_handler"
	"github.com/go-chi/chi/v5"
)

type Server struct{}

func (s *Server) LinksRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/{linkId}", gethandler.GetHandler)
	r.Post("/", posthandler.PostHandler)
	return r
}

func (s *Server) RunApp() {
	err := http.ListenAndServe(`:8080`, s.LinksRouter())
	if err != nil {
		panic(err)
	}
}

func GetNewServer() *Server {
	return &Server{}
}
