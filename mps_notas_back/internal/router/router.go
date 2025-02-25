package router

import (
	"net/http"

	"mps_notas_back/internal/handler"
	"mps_notas_back/internal/service"
)

// New configura e retorna um novo router HTTP
func New(userService *service.UserService) http.Handler {
	// Criar manipuladores
	userHandler := handler.NewUserHandler(userService)
	
	// Criar mux (multiplexador de rotas)
	mux := http.NewServeMux()

	// Rota de verificação de saúde
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})
	
	// Rotas de usuários
	mux.HandleFunc("GET /api/users", userHandler.GetAllUsers)
	mux.HandleFunc("GET /api/users/{id}", userHandler.GetUserByID)
	mux.HandleFunc("POST /api/users", userHandler.CreateUser)
	
	return mux
}