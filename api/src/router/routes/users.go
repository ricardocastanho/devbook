package routes

import (
	"net/http"

	"api/src/controller"
)

var usersRoutes = []route{
    {
        URI: "/users",
        Method: http.MethodGet,
        HandlerFunc: controller.GetUsers,
        NeedAuth: false,
    },
    {
        URI: "/users/{id}",
        Method: http.MethodGet,
        HandlerFunc: controller.FindUser,
        NeedAuth: false,
    },
    {
        URI: "/users",
        Method: http.MethodPost,
        HandlerFunc: controller.CreateUser,
        NeedAuth: false,
    },
    {
        URI: "/users/{id}",
        Method: http.MethodPut,
        HandlerFunc: controller.UpdateUser,
        NeedAuth: false,
    },
    {
        URI: "/users/{id}",
        Method: http.MethodDelete,
        HandlerFunc: controller.DeleteUser,
        NeedAuth: false,
    },
}
