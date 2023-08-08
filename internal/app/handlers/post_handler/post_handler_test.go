package posthandler

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
			w := httptest.NewRecorder()
			bodyReader := strings.NewReader(v.link)
			r := httptest.NewRequest(http.MethodPost, "/", bodyReader)
			PostHandler(w, r)
			res := w.Result()
			assert.Equal(t, v.want.statusCode, res.StatusCode)
			assert.Equal(t, v.want.contentType, res.Header.Get("Content-Type"))
			if v.needCheckBody {
				body, err := io.ReadAll(res.Body)
				require.NoError(t, err)
				err = res.Body.Close()
				require.NoError(t, err)
				val := strings.SplitN(string(body), "/", 4)
				shortLinkID := val[len(val)-1]
				link, _ := sh.GetFullURL(shortLinkID)
				assert.Equal(t, v.link, link)
			}
		})
	}
}
