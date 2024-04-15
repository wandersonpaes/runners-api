package app

import (
	"github.com/gorilla/mux"
	"github.com/wandersonpaes/runners-api/internal/pkg/routes"
)

func SetUp(r *mux.Router) *mux.Router {
	runnersRoutes := routes.Users
	runnersRoutes = append(runnersRoutes, routes.Login)
	runnersRoutes = append(runnersRoutes, routes.PostRoutes...)

	for _, route := range runnersRoutes {
		if route.Authentication {
			r.HandleFunc(route.URI, logger(authenticate(route.Function))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, logger(route.Function)).Methods(route.Method)
		}
	}

	return r
}
