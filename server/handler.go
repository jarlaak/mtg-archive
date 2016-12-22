package server

import (
	"net/http"
	"time"
)

type requestHandler func(http.ResponseWriter, *http.Request)

type AdvancedResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewAdvancedResponseWriter(w http.ResponseWriter) *AdvancedResponseWriter {
	return &AdvancedResponseWriter{w, http.StatusOK}
}

func (mrw *AdvancedResponseWriter) WriteHeader(code int) {
	mrw.statusCode = code
	mrw.ResponseWriter.WriteHeader(code)
}

func LogRequest(handler requestHandler) requestHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		mrw := NewAdvancedResponseWriter(w)
		logger.Info(r.Method, r.URL)
		handler(mrw, r)
		elapsed := time.Since(start)
		logger.Info("request STATUS", mrw.statusCode, ", TIMING: ", elapsed)
	}
}
