package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type route struct {
    URI string
    Method string
    HandlerFunc func(w http.ResponseWriter, r *http.Request)
    NeedAuth bool
}

func BuildRoutes(r *mux.Router) *mux.Router {
    rotas := usersRoutes

    for _, rota := range rotas {
        r.HandleFunc(rota.URI, rota.HandlerFunc).Methods(rota.Method)
    }

    return r
}
