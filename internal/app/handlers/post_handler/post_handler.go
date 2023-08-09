package posthandler

import (
	"fmt"
	"io"
	"net/http"
	"strings"

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
		if !strings.HasPrefix(baseAddress, "http") {
			baseAddress = fmt.Sprintf(`http://%s`, baseAddress)
		}
		w.Write([]byte(fmt.Sprintf(`%s/%s`, baseAddress, shortLink)))
	}
}
