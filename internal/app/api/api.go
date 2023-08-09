package api

import (
	"net/http"

	gethandler "github.com/Vdolganov/shortify/internal/app/handlers/get_handler"
	posthandler "github.com/Vdolganov/shortify/internal/app/handlers/post_handler"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	Address   string
	ShortBase string
}

func (s *Server) LinksRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/{linkId}", gethandler.GetHandler)
	r.Post("/", posthandler.PostHandler(s.ShortBase))
	return r
}

func (s *Server) RunApp() {
	err := http.ListenAndServe(s.Address, s.LinksRouter())
	if err != nil {
		panic(err)
	}
}

func GetNewServer(address, shortBase string) *Server {
	return &Server{
		Address:   address,
		ShortBase: shortBase,
	}
}
