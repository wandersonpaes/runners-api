package router

import (
	"github.com/gorilla/mux"
	"github.com/wandersonpaes/runners-api/internal/app/routes"
)

func CreateMux() *mux.Router {
	r := mux.NewRouter()
	return routes.SetUp(r)
}
