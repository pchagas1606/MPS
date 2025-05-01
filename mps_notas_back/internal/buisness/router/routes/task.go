package routes

import (
	"mps_notas_back/internal/facade"
)

func GenTaskRoutes(f facade.Facade) []Route {
	return []Route{
		{
			URI:         "GET /api/tasks/{id}",
			Func:        f.GetTaskByID,
			RequireAuth: true,
		},
		{
			URI:         "POST /api/tasks",
			Func:        f.CreateTask,
			RequireAuth: true,
		},
		{
			URI:         "PUT /api/tasks/{id}",
			Func:        f.UpdateTask,
			RequireAuth: true,
		},
		{
			URI:         "DELETE /api/tasks/{id}",
			Func:        f.DeleteTask,
			RequireAuth: true,
		},
		{
			URI:         "GET /api/tasks",
			Func:        f.GetAllTasks,
			RequireAuth: true,
		},
	}
}
