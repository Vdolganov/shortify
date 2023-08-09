package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRouter(t *testing.T) {

	s := GetNewServer(":8080", "localhost:8080")
	ts := httptest.NewServer(s.LinksRouter())
	defer ts.Close()

	client := resty.New()
	client.SetBaseURL(ts.URL)
	client.SetRedirectPolicy(resty.FlexibleRedirectPolicy(0))

	var testTable = []struct {
		url       string
		status    int
		payload   string
		getStatus int
	}{
		{url: "/", status: http.StatusCreated, payload: "ya.ru", getStatus: http.StatusTemporaryRedirect},
		{url: "/", status: http.StatusCreated, payload: "google.com", getStatus: http.StatusTemporaryRedirect},
	}

	for _, v := range testTable {
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
	}
}
