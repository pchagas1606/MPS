package routes

import (
	"mps_notas_back/internal/facade"
)

func GenStatusRoutes(f facade.Facade) []Route {
	return []Route{
		{
			URI:         "GET /api/status",
			Func:        f.GetStatus,
			RequireAuth: false,
		},
	}
}
