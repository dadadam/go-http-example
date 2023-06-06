package middleware

import (
	"net"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

// Logger returns a request logging middleware
func Logger(category string, logger logrus.FieldLogger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			defer func() {
				remoteIP, _, err := net.SplitHostPort(r.RemoteAddr)
				if err != nil {
					remoteIP = r.RemoteAddr
				}
				scheme := "http"
				if r.TLS != nil {
					scheme = "https"
				}
				fields := logrus.Fields{
					"status_code": ww.Status(),
					"bytes":       ww.BytesWritten(),
					"category":    category,
					"remote_ip":   remoteIP,
					"proto":       r.Proto,
					"method":      r.Method,
				}

				reqID := middleware.GetReqID(r.Context())
				if len(reqID) > 0 {
					fields["request_id"] = reqID
				}

				logger.WithFields(fields).Infof("%s://%s%s", scheme, r.Host, r.RequestURI)
			}()

			h.ServeHTTP(ww, r)
		}

		return http.HandlerFunc(fn)
	}
}
