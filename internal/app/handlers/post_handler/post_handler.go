package posthandler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Vdolganov/shortify/internal/app/shorter"
)

func PostHandler(baseAddress string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shorterInstance := shorter.GetShorter()
		responseData, err := io.ReadAll(r.Body)
		if err != nil || len(responseData) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		responseString := string(responseData)
		shortLink := shorterInstance.AddLink(responseString)
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf(`http://%s/%s`, baseAddress, shortLink)))
	}
}
