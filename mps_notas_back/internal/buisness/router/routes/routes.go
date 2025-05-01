package routes

import (
	"mps_notas_back/internal/buisness/middlewares"
	"mps_notas_back/internal/facade"
	"net/http"


)

// Route are a struct that represents all routes of API
type Route struct {
	URI         string
	Func        func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

// ConfigRoutes Put all routes in the router Creating a chain of responsability
func ConfigRoutes(r *http.ServeMux, f facade.Facade) http.Handler {
	routes := GenUserRoutes(f)
	routes = append(routes, GenTaskRoutes(f)...)

	for _, route := range routes {
		if route.RequireAuth {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Auth(route.Func)))
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Func))
		}
	}
	return r
}
