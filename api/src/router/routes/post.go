package routes

import (
	"api/src/controller"
	"net/http"
)

var postsRoutes = []route{
	{
		URI:         "/posts",
		Method:      http.MethodPost,
		HandlerFunc: controller.CreatePosts,
		NeedAuth:    true,
	},
}
