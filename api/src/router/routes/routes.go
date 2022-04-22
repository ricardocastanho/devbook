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
    routes := usersRoutes

    for _, route := range routes {
        r.HandleFunc(route.URI, route.HandlerFunc).Methods(route.Method)
    }

    return r
}
