package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Vdolganov/shortify/internal/app/models"
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

func (h *Handlers) JSONPost(w http.ResponseWriter, r *http.Request) {
	var requestBody models.ShorterRequestBody
	jsonDecoder := json.NewDecoder(r.Body)
	err := jsonDecoder.Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(requestBody.URL) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	shortLink, err := h.shorter.AddLink(requestBody.URL)
	shortLinkURL := fmt.Sprintf(`%s/%s`, h.baseURL, shortLink)
	b, err := json.Marshal(models.ShorterResponse{Result: shortLinkURL})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

func New(baseURL string, shorter Shorter) Handlers {
	return Handlers{
		baseURL: baseURL,
		shorter: shorter,
	}
}
