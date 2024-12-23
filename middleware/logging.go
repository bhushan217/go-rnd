package middleware

import (
	"net/http"
	"log"
	"time"
)

type wrappedWriter struct {
	http.ResponseWriter
	status int
}

func (w *wrappedWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[INIT] %s %s", r.Method, r.URL.Path)
		wrapped := &wrappedWriter{ResponseWriter: w}
		next.ServeHTTP(wrapped, r)
		log.Printf("[DONE] %s %s %d %s", r.Method, r.URL.Path, wrapped.status, time.Since(start))
	})
}