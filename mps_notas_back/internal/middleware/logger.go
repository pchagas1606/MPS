package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logger implementa um middleware de logging para requisições HTTP
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// Criar um ResponseWriter personalizado para capturar o status code
		rw := newResponseWriter(w)
		
		// Processar a requisição
		next.ServeHTTP(rw, r)
		
		// Registrar informações da requisição
		duration := time.Since(start)
		log.Printf(
			"%s %s %d %s",
			r.Method,
			r.URL.Path,
			rw.statusCode,
			duration,
		)
	})
}

// responseWriter é um wrapper para http.ResponseWriter que captura o status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// newResponseWriter cria um novo responseWriter
func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

// WriteHeader captura o status code e o passa para o ResponseWriter original
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}