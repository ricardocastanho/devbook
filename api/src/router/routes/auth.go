package routes

import (
	"api/src/controller"
	"net/http"
)

var authRoutes = []route{
	{
		URI:         "/login",
		Method:      http.MethodPost,
		HandlerFunc: controller.Login,
		NeedAuth:    false,
	},
}
