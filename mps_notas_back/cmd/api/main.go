package main

import (
	"fmt"
	"log"
	"mps_notas_back/internal/buisness/config"
	"mps_notas_back/internal/buisness/middleware"
	"mps_notas_back/internal/buisness/router"
	"mps_notas_back/internal/buisness/service"
	"mps_notas_back/internal/cmd"
	factory "mps_notas_back/internal/infra/Factory"
	"mps_notas_back/internal/infra/database"

	"net/http"
)

func main() {
	// Carregar configurações
	cfg := config.New()

	db, err := database.GenConn("internal/infra/database/database.db")
	if err != nil {
		log.Fatal(err)
	}
	// Inicializar repositórios
	userRepo := factory.UserDAOSql{}.GenFactory(db)
	taskRepo := factory.TaskDAOSql{}.GenFactory(db)
	// Inicializar serviços
	userService := service.NewUserService(userRepo)
	taskService := service.NewTaskService(taskRepo)
	statusService := service.NewStatusService()
	// Inicializar comando de status
	statusCommand := cmd.NewGetStatusCommand(statusService)

	// Configurar router
	r := router.New(userService, taskService, statusCommand)

	// Aplicar middlewares, desativado CORS por enquanto, pois não há mais necessidade
	//corsConfig := middleware.DefaultCORSConfig()
	//handler := middleware.CORS(corsConfig)(r)
	handler := middleware.Logger(r)

	// Iniciar servidor HTTP
	serverAddr := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("Servidor iniciado na porta %d", cfg.Port)
	log.Fatal(http.ListenAndServe(serverAddr, handler))
}
