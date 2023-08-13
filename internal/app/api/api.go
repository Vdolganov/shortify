package api

import (
	"net/http"

	gethandler "github.com/Vdolganov/shortify/internal/app/handlers/get_handler"
	posthandler "github.com/Vdolganov/shortify/internal/app/handlers/post_handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	ServerAddress string
	BaseURL       string
}

func (s *Server) LinksRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)

	r.Get("/{linkId}", gethandler.GetHandler)
	r.Post("/", posthandler.PostHandler(s.BaseURL))
	return r
}

func (s *Server) RunApp() {
	err := http.ListenAndServe(s.ServerAddress, s.LinksRouter())
	if err != nil {
		panic(err)
	}
}

func NewServer(ServerAddress, BaseURL string) *Server {
	return &Server{
		ServerAddress,
		BaseURL,
	}
}
