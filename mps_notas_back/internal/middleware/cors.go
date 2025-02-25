package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

// CORSConfig define as configurações para CORS
type CORSConfig struct {
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	ExposedHeaders   []string
	AllowCredentials bool
	MaxAge           int
}

// DefaultCORSConfig retorna uma configuração padrão para CORS
func DefaultCORSConfig() *CORSConfig {
	return &CORSConfig{
		AllowedOrigins:   []string{"http://localhost:3001"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           600,
	}
}

// CORS implementa o middleware 
func CORS(config *CORSConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			
			// Verificar se a origem é permitida
			originAllowed := false
			for _, allowedOrigin := range config.AllowedOrigins {
				if allowedOrigin == "*" || allowedOrigin == origin {
					originAllowed = true
					break
				}
			}
			
			if originAllowed {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				
				// Configurar outros cabeçalhos CORS
				if config.AllowCredentials {
					w.Header().Set("Access-Control-Allow-Credentials", "true")
				}
				
				if len(config.ExposedHeaders) > 0 {
					w.Header().Set("Access-Control-Expose-Headers", strings.Join(config.ExposedHeaders, ", "))
				}
				
				// Lidar com requisições pre-flight (OPTIONS)
				if r.Method == http.MethodOptions {
					w.Header().Set("Access-Control-Allow-Methods", strings.Join(config.AllowedMethods, ", "))
					w.Header().Set("Access-Control-Allow-Headers", strings.Join(config.AllowedHeaders, ", "))
					w.Header().Set("Access-Control-Max-Age", fmt.Sprint(config.MaxAge))
					w.WriteHeader(http.StatusNoContent)
					return
				}
			}
			
			// Processar a requisição
			next.ServeHTTP(w, r)
		})
	}
}