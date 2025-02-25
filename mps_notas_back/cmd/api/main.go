package main

import (
	"fmt"
	"log"
	"mps_notas_back/internal/config"
	"mps_notas_back/internal/middleware"
	"mps_notas_back/internal/repository"
	"mps_notas_back/internal/router"
	"mps_notas_back/internal/service"
	"net/http"
)

func main() {
	// Carregar configurações
	cfg := config.New()

	// Inicializar repositórios
	userRepo := repository.NewUserRepository()

	// Inicializar serviços
	userService := service.NewUserService(userRepo)

	// Configurar router
	r := router.New(userService)

	// Aplicar middlewares, desativado CORS por enquanto, pois não há mais necessidade
	//corsConfig := middleware.DefaultCORSConfig()
	//handler := middleware.CORS(corsConfig)(r)
	handler := middleware.Logger(r)

	// Iniciar servidor HTTP
	serverAddr := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("Servidor iniciado na porta %d", cfg.Port)
	log.Fatal(http.ListenAndServe(serverAddr, handler))
}