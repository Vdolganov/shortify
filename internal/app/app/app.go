package app

import (
	"net/http"

	"github.com/Vdolganov/shortify/internal/app/handlers"
	loghttp "github.com/Vdolganov/shortify/internal/app/middlewares"
	"github.com/Vdolganov/shortify/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Shorter interface {
	GetFullURL(link string) (string, bool)
	AddLink(fullLink string) (string, error)
}

type App struct {
	ServerAddress string
	BaseURL       string
	Shorter       Shorter
}

func (s *App) LinksRouter() chi.Router {
	handlers := handlers.New(s.BaseURL, s.Shorter)
	r := chi.NewRouter()
	r.Use(loghttp.LogHTTP)
	r.Use(middleware.Recoverer)

	r.Get("/{linkId}", handlers.Get)
	r.Post("/", handlers.Post)
	return r
}

func (s *App) RunApp() {
	err := http.ListenAndServe(s.ServerAddress, s.LinksRouter())
	if err != nil {
		panic(err)
	}
}

func New(cnfg config.Config, shorter Shorter) *App {
	return &App{
		ServerAddress: cnfg.ServerAddress,
		BaseURL:       cnfg.BaseURL,
		Shorter:       shorter,
	}
}
