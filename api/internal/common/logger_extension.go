package common

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
)

// DefaultLogger Creates new logger for chi with zap logger.
func DefaultLogger(logger *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()
			defer func() {
				scheme := "http"
				if r.TLS != nil {
					scheme = "https"
				}
				fullPath := fmt.Sprintf("%s://%s%s %s\" ", scheme, r.Host, r.RequestURI, r.Proto)

				reqID := middleware.GetReqID(r.Context())

				logger.Info("",
					zap.Int("status", ww.Status()),
					zap.Int("bytes-written", ww.BytesWritten()),
					zap.String("request-id", reqID),
					zap.String("method", r.Method),
					zap.String("full-path", fullPath),
					zap.String("request-uri", r.RequestURI),
					zap.String("time-duration", time.Since(t1).String()),
					zap.String("ip-address", r.RemoteAddr),
				)
			}()

			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}
