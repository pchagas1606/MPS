package router

import (
	"mps_notas_back/internal/buisness/router/routes"
	"mps_notas_back/internal/buisness/service"
	"mps_notas_back/internal/facade"
	"net/http"
)

// New configura e retorna um novo router HTTP
func New(userService *service.UserService, taskService *service.TaskService) http.Handler {

	//Fachada para as funções que lidam com o handle das chamadas.
	facadeImpl := facade.FacadeImpl.GetFacade(facade.FacadeImpl{}, userService, taskService)

	// Criar mux (multiplexador de rotas)
	mux := http.NewServeMux()

	// Rota de verificação de saúde
	mux.HandleFunc("GET /api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})

	// Configura as rotas e uma corrente de responsabildades.
	routes.ConfigRoutes(mux, facadeImpl)

	return mux
}
