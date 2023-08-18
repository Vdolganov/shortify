package app

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Vdolganov/shortify/internal/app/shorter"
	"github.com/Vdolganov/shortify/internal/app/storage/links"
	"github.com/Vdolganov/shortify/internal/config"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRouter(t *testing.T) {
	st := links.New()
	sh := shorter.New(st)

	s := New(config.Config{BaseURL: "localhost:8080", ServerAddress: ":8080"}, sh)
	ts := httptest.NewServer(s.LinksRouter())
	defer ts.Close()

	client := resty.New()
	client.SetBaseURL(ts.URL)
	client.SetRedirectPolicy(resty.FlexibleRedirectPolicy(0))

	var testTable = []struct {
		name      string
		url       string
		status    int
		payload   string
		getStatus int
	}{
		{name: "1 success test - ya.ru", url: "/", status: http.StatusCreated, payload: "ya.ru", getStatus: http.StatusTemporaryRedirect},
		{name: "2 success test - google.com", url: "/", status: http.StatusCreated, payload: "google.com", getStatus: http.StatusTemporaryRedirect},
	}

	for _, v := range testTable {
		t.Run("test name", func(t *testing.T) {
			resp, err := client.R().
				SetBody([]byte(v.payload)).
				Post(v.url)
			require.NoError(t, err)
			body := resp.Body()
			resp.RawBody().Close()
			assert.Equal(t, v.status, resp.StatusCode())
			val := strings.SplitN(string(body), "/", 4)
			valStr := val[len(val)-1]
			pResp, _ := client.R().
				Get(valStr)
			locationHeader := pResp.Header().Get("Location")
			assert.Equal(t, v.payload, locationHeader)
			assert.Equal(t, v.getStatus, pResp.StatusCode())
		})
	}
}
