package loghttp

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

type responseData struct {
	status int
	size   int
}

type loggingResponseWriter struct {
	http.ResponseWriter
	responseData *responseData
}

func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b)
	r.responseData.size = size
	return size, err
}

func (r *loggingResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
	r.responseData.status = statusCode
}

func LogHttp(h http.Handler) http.Handler {
	logFn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger, _ := zap.NewProduction()
		sugar := logger.Sugar()

		respData := responseData{
			status: 0,
			size:   0,
		}

		start := time.Now()

		lw := loggingResponseWriter{
			w,
			&respData,
		}
		h.ServeHTTP(&lw, r)
		duration := time.Since(start)

		sugar.Infoln(
			"uri", r.RequestURI,
			"method", r.Method,
			"status", respData.status,
			"size", respData.size,
			"duration", duration,
		)

	})
	return http.HandlerFunc(logFn)
}
