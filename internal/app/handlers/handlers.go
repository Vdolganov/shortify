package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Shorter interface {
	GetFullURL(link string) (string, bool)
	AddLink(fullLink string) (string, error)
}

type Handlers struct {
	baseURL string
	shorter Shorter
}

func (h *Handlers) Get(w http.ResponseWriter, r *http.Request) {

	linkID := chi.URLParam(r, "linkId")
	result, exist := h.shorter.GetFullURL(linkID)
	if !exist {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-Type", "text/plain")
	w.Header().Add("Location", result)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (h *Handlers) Post(w http.ResponseWriter, r *http.Request) {
	responseData, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil || len(responseData) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	responseString := string(responseData)
	shortLink, err := h.shorter.AddLink(responseString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf(`%s/%s`, h.baseURL, shortLink)))
}

func New(baseURL string, shorter Shorter) Handlers {
	return Handlers{
		baseURL: baseURL,
		shorter: shorter,
	}
}
