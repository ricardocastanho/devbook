package routes

import (
	"api/src/router/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

type route struct {
	URI         string
	Method      string
	HandlerFunc func(w http.ResponseWriter, r *http.Request)
	NeedAuth    bool
}

func BuildRoutes(r *mux.Router) *mux.Router {
	var routes []route

	routes = append(routes, authRoutes...)
	routes = append(routes, usersRoutes...)

	for _, route := range routes {
		if route.NeedAuth {
			r.HandleFunc(
				route.URI,
				middleware.Logger(middleware.Auth(route.HandlerFunc)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(
				route.URI,
				middleware.Logger(route.HandlerFunc),
			).Methods(route.Method)
		}
	}

	return r
}
