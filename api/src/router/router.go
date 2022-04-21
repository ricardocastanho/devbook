package router

import "github.com/gorilla/mux"

func BuildRoutes() *mux.Router {
    return mux.NewRouter()
}
