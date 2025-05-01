package routes

import (
	"mps_notas_back/internal/facade"
)

func GenUserRoutes(f facade.Facade) []Route {
	return []Route{
		{
			URI:         "POST /api/users",
			Func:        f.CreateUser,
			RequireAuth: false,
		},
		{
			URI:         "GET /api/users",
			Func:        f.GetAllUsers,
			RequireAuth: true,
		},
		{
			URI:         "GET /api/users/{id}",
			Func:        f.GetUserByID,
			RequireAuth: true,
		},
		{
			URI:         "PUT /api/users/{id}",
			Func:        f.Update,
			RequireAuth: true,
		},
		{
			URI:         "DELETE /api/users/{id}",
			Func:        f.Delete,
			RequireAuth: true,
		},
		Route{
			URI:         "POST /api/login",
			Func:        f.Login,
			RequireAuth: false,
		},
		Route{
			URI:         "GET /api/users/report",
			Func:        f.GenerateReport,
			RequireAuth: true,
		},
	}
}
