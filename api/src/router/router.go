package router

import (
	"github.com/gorilla/mux"

	"api/src/router/routes"
)

func BuildRouter() *mux.Router {
    r := mux.NewRouter()

    return routes.BuildRoutes(r)
}
