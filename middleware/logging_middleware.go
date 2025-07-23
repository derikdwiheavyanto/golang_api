package middleware

import (
	"fmt"
	"net/http"
)

type LoggingMiddleware struct {
	handler http.Handler
}

func NewLoggingMidleware(handler http.Handler) *LoggingMiddleware {
	return &LoggingMiddleware{handler}
}

func (m *LoggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Method: ", r.Method, "\nPath: ", r.URL.Path, "\nHeaders: ", r.Header, "\nBody: ", r.Body, "\n")
	m.handler.ServeHTTP(w, r)
}
