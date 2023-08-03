package api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Vdolganov/shortify/internal/app/shorter"
)

type Server struct {
	Shorter shorter.Shorter
}

func (s *Server) getHandler(w http.ResponseWriter, r *http.Request) {
	result, exist := s.Shorter.GetFullURL(r.URL.Path[1:])
	if exist {
		w.Header().Add("Content-Type", "text/plain")
		w.Header().Add("Location", result)
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func (s *Server) postHandler(w http.ResponseWriter, r *http.Request) {
	responseData, err := io.ReadAll(r.Body)
	if err != nil || len(responseData) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	responseString := string(responseData)
	shortLink := s.Shorter.AddLink(responseString)
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf(`http://%s/%s`, r.Host, shortLink)))
}

func (s *Server) mainHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		s.getHandler(w, r)
	}
	if r.Method == http.MethodPost {
		s.postHandler(w, r)
	}
}

func (s *Server) RunApp() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.mainHandler)
	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}

func GetNewServer() *Server {
	return &Server{
		Shorter: *shorter.GetShorter(),
	}
}
