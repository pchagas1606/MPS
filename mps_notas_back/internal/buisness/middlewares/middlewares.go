package middlewares

import (
	"encoding/json"
	"log"
	"mps_notas_back/internal/buisness/auth"
	"net/http"
)

// logger log the request infos
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n%s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Auth verifies if the user in request are athenticated
func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidateToken(r); err != nil {
			Error(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

// Error render a JSON error message
func Error(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode,
		struct {
			Erro string `json:"error"`
		}{
			Erro: erro.Error(),
		},
	)
}
