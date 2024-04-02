package app

import (
	"github.com/gorilla/mux"
	"github.com/wandersonpaes/runners-api/internal/pkg/database"
	"github.com/wandersonpaes/runners-api/internal/pkg/security"
)

func Start() *mux.Router {
	database.SetUp()
	security.SetUp()

	r := mux.NewRouter()
	return SetUp(r)
}
