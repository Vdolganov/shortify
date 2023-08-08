package gethandler

import (
	"net/http"

	"github.com/Vdolganov/shortify/internal/app/shorter"
	"github.com/go-chi/chi/v5"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var shorterInstance = shorter.GetShorter()

	linkID := chi.URLParam(r, "linkId")
	result, exist := shorterInstance.GetFullURL(linkID)
	if exist {
		w.Header().Add("Content-Type", "text/plain")
		w.Header().Add("Location", result)
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}
