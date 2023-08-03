package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Vdolganov/shortify/internal/app/shorter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetHandler(t *testing.T) {
	type Want struct {
		statusCode  int
		header      string
		contentType string
	}
	tests := []struct {
		name  string
		links map[string]string
		want  Want
		path  string
	}{
		{
			name: "Success test 1",
			links: map[string]string{
				"asdasd":  "yandex.ru",
				"jgkjkgk": "google.com",
			},
			want: Want{
				statusCode:  http.StatusTemporaryRedirect,
				header:      "yandex.ru",
				contentType: "text/plain",
			},
			path: "/asdasd",
		},
		{
			name: "Server error 1",
			links: map[string]string{
				"jgkjkgk": "google.com",
			},
			want: Want{
				statusCode:  http.StatusBadRequest,
				header:      "",
				contentType: "",
			},
			path: "/asda",
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			sh := shorter.Shorter{
				Links: v.links,
			}
			server := Server{Shorter: sh}
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, v.path, nil)
			server.getHandler(w, r)
			res := w.Result()
			assert.Equal(t, v.want.statusCode, res.StatusCode)
			assert.Equal(t, v.want.header, res.Header.Get("Location"))
			assert.Equal(t, v.want.contentType, res.Header.Get("Content-Type"))
		})
	}

}

func TestPostHandler(t *testing.T) {
	type Want struct {
		statusCode  int
		contentType string
	}
	tests := []struct {
		name          string
		link          string
		want          Want
		needCheckBody bool
	}{
		{
			name: "Success test 1",
			link: "https://yandex.ru",
			want: Want{
				statusCode:  http.StatusCreated,
				contentType: "text/plain",
			},
			needCheckBody: true,
		},
		{
			name: "Success test 1",
			link: "",
			want: Want{
				statusCode:  http.StatusBadRequest,
				contentType: "",
			},
			needCheckBody: false,
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			sh := shorter.GetShorter()
			server := Server{Shorter: *sh}
			w := httptest.NewRecorder()
			bodyReader := strings.NewReader(v.link)
			r := httptest.NewRequest(http.MethodPost, "/", bodyReader)
			server.postHandler(w, r)
			res := w.Result()
			assert.Equal(t, v.want.statusCode, res.StatusCode)
			assert.Equal(t, v.want.contentType, res.Header.Get("Content-Type"))
			if v.needCheckBody {
				body, err := io.ReadAll(res.Body)
				require.NoError(t, err)
				val := strings.SplitN(string(body), "/", 4)
				shortLinkId := val[len(val)-1]
				link := sh.Links[shortLinkId]
				assert.Equal(t, v.link, link)
			}
		})
	}

}
